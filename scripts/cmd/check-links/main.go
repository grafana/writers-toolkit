package main

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"html"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"regexp"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"
)

const defaultRelativePrefix = "/docs/"

var (
	linkAttributePattern = regexp.MustCompile(`(?is)\b(?:href|src|poster)\s*=\s*(?:"([^"]+)"|'([^']+)'|([^\s"'=<>` + "`" + `\\]+))`)
	srcsetPattern        = regexp.MustCompile(`(?is)\bsrcset\s*=\s*(?:"([^"]+)"|'([^']+)'|([^\s"'=<>` + "`" + `\\]+))`)
	ignoredHTMLPattern   = regexp.MustCompile(`(?is)<(?:pre|code|script|style)\b[^>]*>.*?</(?:pre|code|script|style)>`)
	windowPathPattern    = regexp.MustCompile(`window\.Path\s*=\s*("([^"\\]|\\.)*"|'([^'\\]|\\.)*')`)
)

type options struct {
	outputPath     string
	logRequests    bool
	nginxPort      string
	relativePaths  []string
	startupTimeout time.Duration
	requestTimeout time.Duration
	maxConcurrency int
	excludePattern string
	excludeRegex   *regexp.Regexp
	nginxReadyURL  string
}

type sourceConfig struct {
	RelativePrefix string `json:"relative_prefix"`
}

type linkReport struct {
	ID    string `json:"id,omitempty"`
	URL   string `json:"url"`
	Raw   string `json:"raw,omitempty"`
	Error string `json:"error"`
}

type pageReport struct {
	URL        string       `json:"url"`
	SourcePath string       `json:"source_path,omitempty"`
	Links      []linkReport `json:"links"`
}

type pageLink struct {
	ID  string
	URL string
	Raw string
}

type linkCheckResult struct {
	URL   string
	Error string
}

func main() {
	opts := parseFlags()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	if err := run(ctx, opts); err != nil {
		fmt.Fprintf(os.Stderr, "check-links failed: %v\n", err)
		os.Exit(1)
	}
}

// parseFlags builds runtime options from flags and environment defaults.
func parseFlags() options {
	var opts options
	flag.StringVar(&opts.outputPath, "output", "links.json", "Output JSON path")
	flag.BoolVar(&opts.logRequests, "log-requests", true, "Stream nginx access logs while checks run")
	flag.StringVar(&opts.nginxPort, "nginx-port", "3002", "Local nginx listen port")
	flag.DurationVar(&opts.startupTimeout, "startup-timeout", 45*time.Second, "Nginx startup timeout")
	flag.DurationVar(&opts.requestTimeout, "request-timeout", 15*time.Second, "HTTP request timeout")
	flag.IntVar(&opts.maxConcurrency, "max-concurrency", 16, "Maximum concurrent link checks")
	flag.StringVar(&opts.excludePattern, "exclude", "", "Skip checking URLs matched by this regex")
	flag.Parse()

	if opts.maxConcurrency < 1 {
		opts.maxConcurrency = 1
	}

	opts.relativePaths = resolveRelativePrefixes()
	if opts.excludePattern == "" {
		opts.excludePattern = fmt.Sprintf(
			"(?:http://(?:127\\.0\\.0\\.1|localhost):%s|https://(?:[^/]+\\.)?grafana\\.com)/(?:static|media|api|connect|launch|web)(?:/|$)",
			opts.nginxPort,
		)
	}
	if opts.excludePattern != "" {
		regex, err := regexp.Compile(opts.excludePattern)
		if err != nil {
			fmt.Fprintf(os.Stderr, "invalid exclude regex %q: %v\n", opts.excludePattern, err)
			os.Exit(1)
		}
		opts.excludeRegex = regex
	}
	opts.nginxReadyURL = baseURLForPrefix(opts.relativePaths[0], opts.nginxPort)

	return opts
}

// normalizeRelativePrefix normalizes path prefixes to "/prefix/" form.
func normalizeRelativePrefix(prefix string) string {
	prefix = strings.TrimSpace(prefix)
	if prefix == "" {
		return defaultRelativePrefix
	}
	if !strings.HasPrefix(prefix, "/") {
		prefix = "/" + prefix
	}
	if !strings.HasSuffix(prefix, "/") {
		prefix += "/"
	}
	return prefix
}

// resolveRelativePrefixes resolves documentation URL prefixes from env config.
func resolveRelativePrefixes() []string {
	if prefix := strings.TrimSpace(os.Getenv("RELATIVE_PREFIX")); prefix != "" {
		return []string{normalizeRelativePrefix(prefix)}
	}

	sources := strings.TrimSpace(os.Getenv("SOURCES"))
	if sources == "" {
		return []string{defaultRelativePrefix}
	}

	var configs []sourceConfig
	if err := json.Unmarshal([]byte(sources), &configs); err != nil {
		return []string{defaultRelativePrefix}
	}

	prefixes := make([]string, 0, len(configs))
	seen := map[string]struct{}{}
	for _, config := range configs {
		if strings.TrimSpace(config.RelativePrefix) == "" {
			continue
		}
		normalized := normalizeRelativePrefix(config.RelativePrefix)
		if _, ok := seen[normalized]; ok {
			continue
		}
		seen[normalized] = struct{}{}
		prefixes = append(prefixes, normalized)
	}
	if len(prefixes) == 0 {
		return []string{defaultRelativePrefix}
	}

	return prefixes
}

// baseURLForPrefix returns the local URL to use for a given relative prefix.
func baseURLForPrefix(prefix, port string) string {
	return fmt.Sprintf("http://127.0.0.1:%s%s", port, normalizeRelativePrefix(prefix))
}

// run orchestrates nginx startup, crawl, check, and report writing.
func run(ctx context.Context, opts options) error {
	if err := prepareNginxConfig(opts.nginxPort, opts.relativePaths); err != nil {
		return err
	}

	if err := testNginxConfig(); err != nil {
		return err
	}

	nginxCmd, nginxState, err := startNginx(ctx)
	if err != nil {
		return err
	}
	stopLogStreaming := startNginxLogStreaming(ctx, opts.logRequests, os.Stdout, os.Stderr)
	defer stopLogStreaming()
	defer func() {
		_ = stopProcessGroup(nginxCmd, nginxState.done, 10*time.Second)
	}()

	if err := waitForNginx(ctx, opts.nginxReadyURL, nginxState, opts.startupTimeout); err != nil {
		return err
	}

	if err := runChecks(ctx, opts, nginxState); err != nil {
		return err
	}

	select {
	case <-nginxState.done:
		if err := nginxState.err(); err != nil {
			return fmt.Errorf("nginx exited before completion: %w", err)
		}
		return errors.New("nginx exited before completion")
	default:
		return nil
	}
}

// prepareNginxConfig renders and writes local nginx config artifacts and logs.
func prepareNginxConfig(nginxPort string, relativePrefixes []string) error {
	if err := os.MkdirAll("dist", 0o755); err != nil {
		return fmt.Errorf("create dist dir: %w", err)
	}
	if err := os.MkdirAll("run", 0o755); err != nil {
		return fmt.Errorf("create run dir: %w", err)
	}

	templateBytes, err := os.ReadFile("scripts/cmd/check-links/nginx.local.conf")
	if err != nil {
		return fmt.Errorf("read scripts/cmd/check-links/nginx.local.conf: %w", err)
	}

	rendered, err := renderCheckLinksNginxConfig(string(templateBytes), nginxPort, relativePrefixes, os.Getenv("SHA"))
	if err != nil {
		return err
	}
	if err := os.WriteFile("nginx.conf", []byte(rendered), 0o644); err != nil {
		return fmt.Errorf("write nginx.conf: %w", err)
	}

	pidFile, err := os.OpenFile("run/nginx.pid", os.O_RDONLY|os.O_CREATE, 0o644)
	if err != nil {
		return fmt.Errorf("touch run/nginx.pid: %w", err)
	}
	_ = pidFile.Close()
	if err := truncateFile("access.log"); err != nil {
		return fmt.Errorf("truncate access.log: %w", err)
	}
	if err := truncateFile("error.log"); err != nil {
		return fmt.Errorf("truncate error.log: %w", err)
	}

	return nil
}

// buildServerConfig builds location blocks that map relative prefixes to their dist dir.
// For example, /docs/writers-toolkit/ gets aliased to dist/docs/writers-toolkit/.
func buildServerConfig(relativePrefixes []string, distRoot, sha string) string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("add_header 'Build' '%s';\n\n", sha))

	hasDocsPrefix := false
	for _, relativePrefix := range relativePrefixes {
		normalizedPrefix := normalizeRelativePrefix(relativePrefix)
		prefixWithoutTrailingSlash := strings.TrimSuffix(normalizedPrefix, "/")
		if strings.HasPrefix(prefixWithoutTrailingSlash, "/docs/") {
			hasDocsPrefix = true
		}

		builder.WriteString(fmt.Sprintf(`location = %s {
  return 301 %s;
}

location ^~ %s {
  alias %s%s;
}

		`, prefixWithoutTrailingSlash, normalizedPrefix, normalizedPrefix, distRoot, normalizedPrefix))
	}

	if hasDocsPrefix {
		builder.WriteString(`location ^~ /docs/ {
  proxy_pass https://grafana.com/docs/;
}
`)
	}

	return builder.String()
}

// renderCheckLinksNginxConfig injects runtime values into the local check-links template.
func renderCheckLinksNginxConfig(config, nginxPort string, relativePrefixes []string, sha string) (string, error) {
	replacements := []struct {
		old string
		new string
	}{
		{"listen 80;", fmt.Sprintf("listen %s;", nginxPort)},
		{"include /etc/nginx/build.conf;", buildServerConfig(relativePrefixes, "dist", sha)},
		{"include /etc/nginx/locations.conf;", "include deploy-preview/locations.conf;"},
	}

	rendered := config
	for _, replacement := range replacements {
		if !strings.Contains(rendered, replacement.old) {
			return "", fmt.Errorf("expected nginx config to contain %q", replacement.old)
		}
		rendered = strings.ReplaceAll(rendered, replacement.old, replacement.new)
	}

	return rendered, nil
}

// truncateFile truncates or creates a file.
func truncateFile(path string) error {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		return err
	}
	return file.Close()
}

// testNginxConfig validates the generated nginx configuration.
func testNginxConfig() error {
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("get working directory: %w", err)
	}

	cmd := exec.Command("nginx", "-p", cwd, "-c", "nginx.conf", "-t")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("validate nginx config: %w", err)
	}
	return nil
}

// startNginx starts nginx in foreground mode and tracks its process state.
func startNginx(ctx context.Context) (*exec.Cmd, *processState, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, nil, fmt.Errorf("get working directory: %w", err)
	}

	cmd := exec.CommandContext(ctx, "nginx", "-p", cwd, "-c", "nginx.conf", "-g", "daemon off;")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if runtime.GOOS != "windows" {
		cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	}

	if err := cmd.Start(); err != nil {
		return nil, nil, fmt.Errorf("start nginx: %w", err)
	}

	state := &processState{done: make(chan struct{})}
	go func() {
		state.setErr(cmd.Wait())
	}()

	return cmd, state, nil
}

// waitForNginx waits until nginx responds or exits/fails.
func waitForNginx(ctx context.Context, readyURL string, nginxState *processState, timeout time.Duration) error {
	client := &http.Client{Timeout: 2 * time.Second}
	deadline := time.NewTimer(timeout)
	defer deadline.Stop()
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-deadline.C:
			return fmt.Errorf("nginx did not become ready within %s", timeout)
		case <-nginxState.done:
			if err := nginxState.err(); err != nil {
				return fmt.Errorf("nginx exited before ready: %w", err)
			}
			return errors.New("nginx exited before ready")
		case <-ticker.C:
			response, err := client.Get(readyURL)
			if err != nil {
				continue
			}
			_ = response.Body.Close()
			if response.StatusCode > 0 {
				return nil
			}
		}
	}
}

// runChecks crawls source pages, checks targets, and writes a links report.
func runChecks(ctx context.Context, opts options, nginxState *processState) error {
	sourceURLs, err := collectSourcePageURLs(opts.relativePaths, opts.nginxPort)
	if err != nil {
		return err
	}

	localClient := newHTTPClient(opts.requestTimeout)
	pageTargets := make(map[string][]pageLink, len(sourceURLs))
	sourcePaths := make(map[string]string, len(sourceURLs))
	checkResults := make(map[string]linkCheckResult, len(sourceURLs))

	for _, sourceURL := range sourceURLs {
		fmt.Fprintf(os.Stdout, "collecting links from: %s\n", sourceURL)
		body, err := fetchPageBody(ctx, localClient, sourceURL)
		if err != nil {
			return err
		}

		checkResults[sourceURL] = linkCheckResult{URL: sourceURL}
		sourcePath, links := extractPageData(sourceURL, string(body), opts.nginxPort)
		sourcePaths[sourceURL] = sourcePath
		pageTargets[sourceURL] = filterPageLinks(links, opts.excludeRegex)
	}

	targetToPages := invertPageTargets(pageTargets)
	targetURLs := make([]string, 0, len(targetToPages))
	for targetURL := range targetToPages {
		if _, alreadyChecked := checkResults[targetURL]; alreadyChecked {
			continue
		}
		targetURLs = append(targetURLs, targetURL)
	}
	sort.Strings(targetURLs)

	if err := checkTargets(ctx, localClient, opts, nginxState, targetURLs, checkResults); err != nil {
		return err
	}

	reports := buildReports(sourceURLs, sourcePaths, pageTargets, checkResults)
	return writeReport(opts.outputPath, reports)
}

// extractWindowPath extracts window.Path from rendered page HTML.
func extractWindowPath(body string) string {
	match := windowPathPattern.FindStringSubmatch(body)
	if len(match) < 2 {
		return ""
	}

	quoted := strings.TrimSpace(match[1])
	if quoted == "" {
		return ""
	}
	unquoted, err := unquoteJavaScriptString(quoted)
	if err != nil {
		if len(quoted) >= 2 {
			return quoted[1 : len(quoted)-1]
		}
		return ""
	}
	return strings.TrimSpace(unquoted)
}

// unquoteJavaScriptString unquotes single- or double-quoted JavaScript string literals.
func unquoteJavaScriptString(value string) (string, error) {
	if len(value) < 2 {
		return "", fmt.Errorf("invalid quoted value: %q", value)
	}
	if strings.HasPrefix(value, "'") && strings.HasSuffix(value, "'") {
		inner := value[1 : len(value)-1]
		inner = strings.ReplaceAll(inner, `\`, `\\`)
		inner = strings.ReplaceAll(inner, `"`, `\"`)
		return strconv.Unquote(`"` + inner + `"`)
	}
	return strconv.Unquote(value)
}

// collectSourcePageURLs collects rendered page URLs from dist html output.
func collectSourcePageURLs(relativePrefixes []string, port string) ([]string, error) {
	urls := make([]string, 0)
	seen := map[string]struct{}{}

	for _, relativePrefix := range relativePrefixes {
		prefixRoot := filepath.Join("dist", filepath.FromSlash(strings.TrimPrefix(normalizeRelativePrefix(relativePrefix), "/")))
		err := filepath.Walk(prefixRoot, func(path string, info os.FileInfo, walkErr error) error {
			if walkErr != nil {
				return walkErr
			}
			if info.IsDir() || !shouldCrawlHTMLFile(info.Name()) {
				return nil
			}

			urlPath, ok := htmlFilePathToURLPath(path)
			if !ok {
				return nil
			}

			sourceURL := fmt.Sprintf("http://127.0.0.1:%s%s", port, urlPath)
			if _, exists := seen[sourceURL]; exists {
				return nil
			}
			seen[sourceURL] = struct{}{}
			urls = append(urls, sourceURL)
			return nil
		})
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				continue
			}
			return nil, fmt.Errorf("walk %s: %w", prefixRoot, err)
		}
	}

	sort.Strings(urls)
	return urls, nil
}

func shouldCrawlHTMLFile(name string) bool {
	return strings.HasSuffix(name, ".html") && name != "unstyled.html"
}

// htmlFilePathToURLPath maps a dist html path to a URL path.
func htmlFilePathToURLPath(path string) (string, bool) {
	relativePath, err := filepath.Rel("dist", path)
	if err != nil {
		return "", false
	}

	relativePath = filepath.ToSlash(relativePath)
	switch {
	case relativePath == "index.html":
		return "/", true
	case strings.HasSuffix(relativePath, "/index.html"):
		return "/" + strings.TrimSuffix(relativePath, "index.html"), true
	case strings.HasSuffix(relativePath, ".html"):
		return "/" + relativePath, true
	default:
		return "", false
	}
}

// newHTTPClient builds the HTTP client used for page and target requests.
func newHTTPClient(timeout time.Duration) *http.Client {
	transport := &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		DialContext:           (&net.Dialer{Timeout: timeout, KeepAlive: 30 * time.Second}).DialContext,
		TLSHandshakeTimeout:   timeout,
		ResponseHeaderTimeout: timeout,
		ExpectContinueTimeout: 1 * time.Second,
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   20,
		IdleConnTimeout:       90 * time.Second,
	}

	return &http.Client{
		Timeout:   timeout,
		Transport: transport,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if len(via) >= 10 {
				return errors.New("stopped after 10 redirects")
			}
			return nil
		},
	}
}

// fetchPageBody fetches and validates a page body for link extraction.
func fetchPageBody(ctx context.Context, client *http.Client, pageURL string) ([]byte, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, pageURL, nil)
	if err != nil {
		return nil, fmt.Errorf("build request for %s: %w", pageURL, err)
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("fetch %s: %w", pageURL, err)
	}
	defer func() { _ = response.Body.Close() }()

	if response.StatusCode < 200 || response.StatusCode >= 400 {
		return nil, fmt.Errorf("fetch %s: unexpected status %d", pageURL, response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", pageURL, err)
	}

	return body, nil
}

// extractPageData extracts source path and normalized unique link targets from HTML.
func extractPageData(pageURL, body, nginxPort string) (string, []pageLink) {
	sourcePath := extractWindowPath(body)
	body = ignoredHTMLPattern.ReplaceAllString(body, "")

	links := make([]pageLink, 0)
	seen := map[string]struct{}{}
	linkIndex := 0

	appendLink := func(raw string) {
		normalized, ok := normalizeLinkURL(pageURL, html.UnescapeString(raw), nginxPort)
		if !ok {
			return
		}
		if _, exists := seen[normalized]; exists {
			return
		}
		linkIndex++
		seen[normalized] = struct{}{}
		links = append(links, pageLink{
			ID:  fmt.Sprintf("%s#%d", pageURL, linkIndex),
			URL: normalized,
			Raw: strings.TrimSpace(raw),
		})
	}

	for _, match := range linkAttributePattern.FindAllStringSubmatch(body, -1) {
		if value, ok := firstNonEmptyCapture(match); ok {
			appendLink(value)
		}
	}

	for _, match := range srcsetPattern.FindAllStringSubmatch(body, -1) {
		value, ok := firstNonEmptyCapture(match)
		if !ok {
			continue
		}
		for _, candidate := range strings.Split(value, ",") {
			fields := strings.Fields(strings.TrimSpace(candidate))
			if len(fields) > 0 {
				appendLink(fields[0])
			}
		}
	}

	return sourcePath, links
}

// firstNonEmptyCapture returns the first present capture group in a regex match.
func firstNonEmptyCapture(match []string) (string, bool) {
	for _, value := range match[1:] {
		if value != "" {
			return value, true
		}
	}
	return "", false
}

// normalizeLinkURL normalizes link URLs and filters unsupported/external targets.
func normalizeLinkURL(pageURL, raw, nginxPort string) (string, bool) {
	raw = strings.TrimSpace(raw)
	if raw == "" || raw == "#" {
		return "", false
	}

	lower := strings.ToLower(raw)
	for _, prefix := range []string{"mailto:", "tel:", "javascript:", "data:"} {
		if strings.HasPrefix(lower, prefix) {
			return "", false
		}
	}

	base, err := url.Parse(pageURL)
	if err != nil {
		return "", false
	}
	reference, err := url.Parse(raw)
	if err != nil {
		return "", false
	}

	resolved := base.ResolveReference(reference)
	if resolved.Scheme != "http" && resolved.Scheme != "https" {
		return "", false
	}
	resolved.Fragment = ""

	host := strings.ToLower(resolved.Hostname())
	if isLocalPreviewHost(host) {
		// Always force local preview links to the checker nginx instance.
		resolved.Scheme = "http"
		resolved.Host = net.JoinHostPort("127.0.0.1", nginxPort)
	}
	if !isLocalPreviewHost(host) && !isGrafanaHost(host) {
		return "", false
	}

	return resolved.String(), true
}

// isLocalPreviewHost reports whether a host belongs to the local preview server.
func isLocalPreviewHost(host string) bool {
	return host == "127.0.0.1" || host == "localhost"
}

// isGrafanaHost reports whether a host is grafana.com or a subdomain.
func isGrafanaHost(host string) bool {
	return host == "grafana.com" || strings.HasSuffix(host, ".grafana.com")
}

// filterPageLinks applies an optional regex exclusion to candidate page links.
func filterPageLinks(links []pageLink, excludeRegex *regexp.Regexp) []pageLink {
	if excludeRegex == nil {
		return links
	}

	filtered := make([]pageLink, 0, len(links))
	for _, link := range links {
		if excludeRegex.MatchString(link.URL) {
			continue
		}
		filtered = append(filtered, link)
	}
	return filtered
}

// invertPageTargets creates a reverse index from target URL to source pages.
func invertPageTargets(pageTargets map[string][]pageLink) map[string][]string {
	targetToPages := map[string][]string{}
	for pageURL, links := range pageTargets {
		for _, link := range links {
			targetToPages[link.URL] = append(targetToPages[link.URL], pageURL)
		}
	}
	return targetToPages
}

// checkTargets concurrently validates each target URL and stores outcomes.
func checkTargets(ctx context.Context, client *http.Client, opts options, nginxState *processState, targetURLs []string, results map[string]linkCheckResult) error {
	if len(targetURLs) == 0 {
		return nil
	}

	type checkOutcome struct {
		result linkCheckResult
		err    error
	}

	jobs := make(chan string)
	outcomes := make(chan checkOutcome, len(targetURLs))
	var workers sync.WaitGroup

	for i := 0; i < opts.maxConcurrency; i++ {
		workers.Add(1)
		go func() {
			defer workers.Done()
			for targetURL := range jobs {
				result, err := checkTarget(ctx, client, targetURL)
				outcomes <- checkOutcome{result: result, err: err}
			}
		}()
	}

	go func() {
		for _, targetURL := range targetURLs {
			select {
			case <-ctx.Done():
				close(jobs)
				workers.Wait()
				close(outcomes)
				return
			case <-nginxState.done:
				close(jobs)
				workers.Wait()
				close(outcomes)
				return
			case jobs <- targetURL:
			}
		}
		close(jobs)
		workers.Wait()
		close(outcomes)
	}()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-nginxState.done:
			if err := nginxState.err(); err != nil {
				return fmt.Errorf("nginx exited while link checks were running: %w", err)
			}
			return errors.New("nginx exited while link checks were running")
		case outcome, ok := <-outcomes:
			if !ok {
				return nil
			}
			if outcome.err != nil {
				return outcome.err
			}
			results[outcome.result.URL] = outcome.result
		}
	}
}

// checkTarget performs a single HTTP check for one target URL.
func checkTarget(ctx context.Context, client *http.Client, targetURL string) (linkCheckResult, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, targetURL, nil)
	if err != nil {
		return linkCheckResult{}, fmt.Errorf("build request for %s: %w", targetURL, err)
	}

	response, err := client.Do(request)
	if err != nil {
		return linkCheckResult{
			URL:   targetURL,
			Error: classifyRequestError(err),
		}, nil
	}
	defer func() { _ = response.Body.Close() }()

	_, _ = io.CopyN(io.Discard, response.Body, 1)

	if response.StatusCode >= 200 && response.StatusCode < 400 {
		return linkCheckResult{URL: targetURL}, nil
	}

	return linkCheckResult{
		URL:   targetURL,
		Error: strconv.Itoa(response.StatusCode),
	}, nil
}

// classifyRequestError converts request errors into stable report strings.
func classifyRequestError(err error) string {
	if err == nil {
		return ""
	}

	var urlErr *url.Error
	if errors.As(err, &urlErr) {
		if isTimeoutError(urlErr.Err) {
			return "timeout"
		}
		return urlErr.Err.Error()
	}
	if isTimeoutError(err) {
		return "timeout"
	}
	return err.Error()
}

// isTimeoutError reports whether an error is timeout-related.
func isTimeoutError(err error) bool {
	if err == nil {
		return false
	}
	var netErr net.Error
	if errors.As(err, &netErr) && netErr.Timeout() {
		return true
	}
	return errors.Is(err, context.DeadlineExceeded)
}

// buildReports assembles page-centric broken-link report objects.
func buildReports(sourceURLs []string, sourcePaths map[string]string, pageTargets map[string][]pageLink, results map[string]linkCheckResult) []pageReport {
	reports := make([]pageReport, 0, len(sourceURLs))
	for _, sourceURL := range sourceURLs {
		targets := pageTargets[sourceURL]
		brokenLinks := make([]linkReport, 0)
		for _, link := range targets {
			result, ok := results[link.URL]
			if !ok || result.Error == "" {
				continue
			}
			brokenLinks = append(brokenLinks, linkReport{
				ID:    link.ID,
				URL:   link.URL,
				Raw:   link.Raw,
				Error: result.Error,
			})
		}
		if len(brokenLinks) == 0 {
			continue
		}
		reports = append(reports, pageReport{
			URL:        sourceURL,
			SourcePath: sourcePaths[sourceURL],
			Links:      brokenLinks,
		})
	}
	return reports
}

// writeReport marshals report data to JSON and writes it to disk.
func writeReport(path string, reports []pageReport) error {
	content, err := json.MarshalIndent(reports, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal report: %w", err)
	}
	content = append(content, '\n')

	if err := os.WriteFile(path, content, 0o644); err != nil {
		return fmt.Errorf("write %s: %w", path, err)
	}
	return nil
}

// startNginxLogStreaming starts asynchronous nginx error/access log tailing.
func startNginxLogStreaming(parent context.Context, logRequests bool, stdout, stderr io.Writer) func() {
	stopLoggers := []func(){startLogStreaming(parent, "error.log", "[nginx-error]", stderr)}
	if logRequests {
		stopLoggers = append(stopLoggers, startLogStreaming(parent, "access.log", "[nginx-access]", stdout))
	}

	return func() {
		for i := len(stopLoggers) - 1; i >= 0; i-- {
			stopLoggers[i]()
		}
	}
}

// startLogStreaming tails a log file and writes prefixed lines to output.
func startLogStreaming(parent context.Context, path, prefix string, output io.Writer) func() {
	ctx, cancel := context.WithCancel(parent)
	done := make(chan struct{})

	go func() {
		defer close(done)

		file, err := waitForFile(ctx, path, 5*time.Second)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to stream %s: %v\n", path, err)
			return
		}
		defer func() { _ = file.Close() }()

		reader := bufio.NewReader(file)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if errors.Is(err, io.EOF) {
					select {
					case <-ctx.Done():
						return
					case <-time.After(100 * time.Millisecond):
						continue
					}
				}
				fmt.Fprintf(os.Stderr, "failed to read %s: %v\n", path, err)
				return
			}

			_, _ = fmt.Fprintf(output, "%s %s", prefix, line)
		}
	}()

	return func() {
		cancel()
		<-done
	}
}

// waitForFile waits for a log file to exist and opens it.
func waitForFile(ctx context.Context, path string, timeout time.Duration) (*os.File, error) {
	deadline := time.NewTimer(timeout)
	defer deadline.Stop()
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for {
		file, err := os.Open(path)
		if err == nil {
			return file, nil
		}
		if !errors.Is(err, os.ErrNotExist) {
			return nil, err
		}

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-deadline.C:
			return nil, fmt.Errorf("timed out waiting for %s", path)
		case <-ticker.C:
		}
	}
}

// stopProcessGroup attempts graceful then forced process-group shutdown.
func stopProcessGroup(cmd *exec.Cmd, done <-chan struct{}, timeout time.Duration) error {
	if cmd == nil || cmd.Process == nil {
		return nil
	}

	select {
	case <-done:
		return nil
	default:
	}

	if runtime.GOOS == "windows" {
		_ = cmd.Process.Kill()
	} else {
		_ = syscall.Kill(-cmd.Process.Pid, syscall.SIGTERM)
	}

	timer := time.NewTimer(timeout)
	defer timer.Stop()

	select {
	case <-done:
		return nil
	case <-timer.C:
		if runtime.GOOS == "windows" {
			return cmd.Process.Kill()
		}
		return syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
	}
}

type processState struct {
	done chan struct{}
	mu   sync.Mutex
	errv error
}

// setErr stores process exit error and marks the process as done.
func (p *processState) setErr(err error) {
	p.mu.Lock()
	p.errv = err
	p.mu.Unlock()
	close(p.done)
}

// err returns the stored process exit error.
func (p *processState) err() error {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.errv
}
