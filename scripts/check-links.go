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
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"regexp"
	"runtime"
	"sort"
	"strings"
	"sync"
	"syscall"
	"time"
)

const defaultRelativePrefix = "/docs/"

type options struct {
	baseURL        string
	include        string
	exclude        string
	outputPath     string
	logRequests    bool
	nginxPort      string
	relativePrefix string
	relativePaths  []string
	startupTimeout time.Duration
}

type sourceConfig struct {
	RelativePrefix string `json:"relative_prefix"`
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

func parseFlags() options {
	var opts options
	flag.StringVar(&opts.baseURL, "base-url", "", "Base URL to crawl")
	flag.StringVar(&opts.include, "include", "", "Muffet include pattern")
	flag.StringVar(&opts.exclude, "exclude", "", "Muffet exclude pattern")
	flag.StringVar(&opts.outputPath, "output", "links.json", "Output JSON path")
	flag.BoolVar(&opts.logRequests, "log-requests", true, "Stream nginx access logs while muffet runs")
	flag.StringVar(&opts.nginxPort, "nginx-port", "3002", "Local nginx listen port")
	flag.DurationVar(&opts.startupTimeout, "startup-timeout", 45*time.Second, "Nginx startup timeout")
	flag.Parse()

	relativePrefixes := resolveRelativePrefixes()
	opts.relativePaths = relativePrefixes
	opts.relativePrefix = relativePrefixes[0]
	if opts.baseURL == "" {
		opts.baseURL = baseURLForPrefix(opts.relativePrefix, opts.nginxPort)
	}
	if opts.exclude == "" {
		opts.exclude = defaultExcludePattern(opts.nginxPort)
	}

	return opts
}

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
		if strings.TrimSpace(config.RelativePrefix) != "" {
			normalized := normalizeRelativePrefix(config.RelativePrefix)
			if _, ok := seen[normalized]; ok {
				continue
			}
			seen[normalized] = struct{}{}
			prefixes = append(prefixes, normalized)
		}
	}
	if len(prefixes) > 0 {
		return prefixes
	}

	return []string{defaultRelativePrefix}
}

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

func baseURLForPrefix(prefix, port string) string {
	return fmt.Sprintf("http://127.0.0.1:%s%s", port, normalizeRelativePrefix(prefix))
}

func defaultExcludePattern(port string) string {
	return fmt.Sprintf("http://(?:127\\.0\\.0\\.1|localhost):%s/(?:static|media|api|connect)", port)
}

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

	if err := waitForNginx(ctx, opts.baseURL, nginxState, opts.startupTimeout); err != nil {
		return err
	}

	if err := runMuffet(ctx, opts, nginxState); err != nil {
		return err
	}

	select {
	case <-nginxState.done:
		err := nginxState.err()
		if err != nil {
			return fmt.Errorf("nginx exited before completion: %w", err)
		}
		return errors.New("nginx exited before completion")
	default:
		return nil
	}
}

func prepareNginxConfig(nginxPort string, relativePrefixes []string) error {
	if err := os.MkdirAll("dist", 0o755); err != nil {
		return fmt.Errorf("create dist dir: %w", err)
	}
	if err := os.MkdirAll("run", 0o755); err != nil {
		return fmt.Errorf("create run dir: %w", err)
	}

	templateBytes, err := os.ReadFile("deploy-preview/nginx.conf")
	if err != nil {
		return fmt.Errorf("read deploy-preview/nginx.conf: %w", err)
	}

	rendered, err := renderLocalNginxConfig(string(templateBytes), nginxPort, relativePrefixes, os.Getenv("SHA"))
	if err != nil {
		return err
	}
	if err := os.WriteFile("nginx.conf", []byte(rendered), 0o644); err != nil {
		return fmt.Errorf("write nginx.conf: %w", err)
	}

	if err := touch("run/nginx.pid"); err != nil {
		return fmt.Errorf("touch run/nginx.pid: %w", err)
	}
	if err := truncateFile("access.log"); err != nil {
		return fmt.Errorf("truncate access.log: %w", err)
	}
	if err := truncateFile("error.log"); err != nil {
		return fmt.Errorf("truncate error.log: %w", err)
	}

	return nil
}

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

func renderLocalNginxConfig(config, nginxPort string, relativePrefixes []string, sha string) (string, error) {
	replacements := []struct {
		old string
		new string
	}{
		{"pid /run/nginx.pid;", "pid run/nginx.pid;"},
		{"access_log /var/log/nginx/access.log;", "access_log access.log;"},
		{"error_log /var/log/nginx/error.log;", "error_log error.log;"},
		{"listen 80;", fmt.Sprintf("listen %s;", nginxPort)},
		{"include /etc/nginx/build.conf;", buildServerConfig(relativePrefixes, "dist", sha)},
		{"root /usr/share/nginx/dist/;", "root dist/;"},
		{"alias /usr/share/nginx/dist/;", "alias dist/;"},
		{"alias /usr/share/nginx/assets/$1;", "alias dist/static/$1;"},
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

func touch(path string) error {
	file, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0o644)
	if err != nil {
		return err
	}
	return file.Close()
}

func truncateFile(path string) error {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		return err
	}
	return file.Close()
}

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

	state := newProcessState()
	go func() {
		state.setErr(cmd.Wait())
	}()

	return cmd, state, nil
}

func waitForNginx(ctx context.Context, url string, nginxState *processState, timeout time.Duration) error {
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
			err := nginxState.err()
			if err != nil {
				return fmt.Errorf("nginx exited before ready: %w", err)
			}
			return errors.New("nginx exited before ready")
		case <-ticker.C:
			resp, err := client.Get(url)
			if err != nil {
				continue
			}
			_ = resp.Body.Close()
			if resp.StatusCode > 0 {
				return nil
			}
		}
	}
}

func runMuffet(ctx context.Context, opts options, nginxState *processState) error {
	sourceURLs, err := collectSourcePageURLs(opts.relativePaths, opts.nginxPort)
	if err != nil {
		return err
	}
	if len(sourceURLs) == 0 {
		return writeMergedReport(opts.outputPath, nil)
	}

	client := &http.Client{Timeout: 20 * time.Second}
	pageTargets := make(map[string][]string, len(sourceURLs))
	uniqueTargets := make([]string, 0)
	sourceURLSet := make(map[string]struct{}, len(sourceURLs))
	for _, sourceURL := range sourceURLs {
		sourceURLSet[sourceURL] = struct{}{}
	}
	seenTargets := map[string]struct{}{}
	for _, sourceURL := range sourceURLs {
		fmt.Fprintf(os.Stdout, "collecting links from: %s\n", sourceURL)
		targets, err := collectTargetsFromPage(ctx, client, sourceURL)
		if err != nil {
			return err
		}
		pageTargets[sourceURL] = targets
		for _, target := range targets {
			if _, ok := sourceURLSet[target]; ok {
				continue
			}
			if _, ok := seenTargets[target]; ok {
				continue
			}
			seenTargets[target] = struct{}{}
			uniqueTargets = append(uniqueTargets, target)
		}
	}

	if len(uniqueTargets) == 0 {
		return writeMergedReport(opts.outputPath, nil)
	}

	manifestURL, cleanupManifest, err := writeTargetManifest(uniqueTargets, opts.nginxPort)
	if err != nil {
		return err
	}
	defer cleanupManifest()

	tempReportPath := filepath.Join(os.TempDir(), "muffet-targets.json")
	fmt.Fprintf(os.Stdout, "running muffet unique-target check (%d urls)\n", len(uniqueTargets))
	if err := runMuffetOnce(ctx, opts, nginxState, manifestURL, tempReportPath); err != nil {
		return err
	}

	brokenTargets, err := readBrokenTargets(tempReportPath)
	if err != nil {
		return err
	}

	reports := buildReports(sourceURLs, pageTargets, brokenTargets)
	return writeMergedReport(opts.outputPath, reports)
}

func runMuffetOnce(ctx context.Context, opts options, nginxState *processState, startURL, outputPath string) error {
	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("create %s: %w", outputPath, err)
	}
	defer func() {
		_ = outFile.Close()
	}()

	args := []string{}
	if opts.include != "" {
		args = append(args, "--include="+opts.include)
	}
	if opts.exclude != "" {
		args = append(args, "--exclude="+opts.exclude)
	}
	args = append(args, "--one-page-only", "-f", "-b", "9999", "--format=json")
	args = append(args, startURL)

	cmd := exec.CommandContext(ctx, "muffet", args...)
	cmd.Stdout = outFile
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("start muffet: %w", err)
	}

	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()

	for {
		select {
		case err := <-done:
			if err != nil {
				if isMuffetReportExit(err) {
					if jsonErr := validateJSONReport(outputPath); jsonErr == nil {
						return nil
					}
				}
				return fmt.Errorf("run muffet: %w", err)
			}
			return nil
		case <-nginxState.done:
			_ = stopCommand(cmd, done, 5*time.Second)
			err := nginxState.err()
			if err != nil {
				return fmt.Errorf("nginx exited while muffet was running: %w", err)
			}
			return errors.New("nginx exited while muffet was running")
		case <-ctx.Done():
			_ = stopCommand(cmd, done, 5*time.Second)
			return ctx.Err()
		}
	}
}

type linkReport struct {
	URL   string `json:"url"`
	Error string `json:"error"`
}

type pageReport struct {
	URL   string       `json:"url"`
	Links []linkReport `json:"links"`
}

func readReport(path string) ([]pageReport, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", path, err)
	}

	var report []pageReport
	if err := json.Unmarshal(content, &report); err != nil {
		return nil, fmt.Errorf("parse %s: %w", path, err)
	}

	return report, nil
}

func readBrokenTargets(path string) (map[string]linkReport, error) {
	report, err := readReport(path)
	if err != nil {
		return nil, err
	}

	brokenTargets := map[string]linkReport{}
	for _, page := range report {
		for _, link := range page.Links {
			if _, ok := brokenTargets[link.URL]; ok {
				continue
			}
			brokenTargets[link.URL] = link
		}
	}

	return brokenTargets, nil
}

func buildReports(sourceURLs []string, pageTargets map[string][]string, brokenTargets map[string]linkReport) []pageReport {
	reports := make([]pageReport, 0, len(sourceURLs))
	for _, sourceURL := range sourceURLs {
		targets := pageTargets[sourceURL]
		brokenLinks := make([]linkReport, 0)
		for _, target := range targets {
			link, ok := brokenTargets[target]
			if !ok {
				continue
			}
			brokenLinks = append(brokenLinks, link)
		}
		if len(brokenLinks) == 0 {
			continue
		}
		reports = append(reports, pageReport{
			URL:   sourceURL,
			Links: brokenLinks,
		})
	}
	return reports
}

func writeMergedReport(path string, reports []pageReport) error {
	content, err := json.MarshalIndent(reports, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal merged report: %w", err)
	}
	content = append(content, '\n')

	if err := os.WriteFile(path, content, 0o644); err != nil {
		return fmt.Errorf("write %s: %w", path, err)
	}

	return nil
}

func collectSourcePageURLs(relativePrefixes []string, port string) ([]string, error) {
	urls := make([]string, 0)
	seen := map[string]struct{}{}
	for _, relativePrefix := range relativePrefixes {
		prefixRoot := filepath.Join("dist", filepath.FromSlash(strings.TrimPrefix(normalizeRelativePrefix(relativePrefix), "/")))
		err := filepath.Walk(prefixRoot, func(path string, info os.FileInfo, walkErr error) error {
			if walkErr != nil {
				return walkErr
			}
			if info.IsDir() || !strings.HasSuffix(info.Name(), ".html") {
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

func collectTargetsFromPage(ctx context.Context, client *http.Client, pageURL string) ([]string, error) {
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

	return extractLinks(pageURL, string(body)), nil
}

var (
	linkAttributePattern = regexp.MustCompile(`(?is)\b(?:href|src|poster)=["']([^"'<>]+)["']`)
	srcsetPattern        = regexp.MustCompile(`(?is)\bsrcset=["']([^"'<>]+)["']`)
)

func extractLinks(pageURL, body string) []string {
	links := make([]string, 0)
	seen := map[string]struct{}{}

	for _, match := range linkAttributePattern.FindAllStringSubmatch(body, -1) {
		if len(match) < 2 {
			continue
		}
		normalized, ok := normalizeLinkURL(pageURL, html.UnescapeString(match[1]))
		if !ok {
			continue
		}
		if _, exists := seen[normalized]; exists {
			continue
		}
		seen[normalized] = struct{}{}
		links = append(links, normalized)
	}

	for _, match := range srcsetPattern.FindAllStringSubmatch(body, -1) {
		if len(match) < 2 {
			continue
		}
		for _, candidate := range strings.Split(match[1], ",") {
			fields := strings.Fields(strings.TrimSpace(candidate))
			if len(fields) == 0 {
				continue
			}
			normalized, ok := normalizeLinkURL(pageURL, html.UnescapeString(fields[0]))
			if !ok {
				continue
			}
			if _, exists := seen[normalized]; exists {
				continue
			}
			seen[normalized] = struct{}{}
			links = append(links, normalized)
		}
	}

	return links
}

func normalizeLinkURL(pageURL, raw string) (string, bool) {
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

	return resolved.String(), true
}

func writeTargetManifest(targets []string, port string) (string, func(), error) {
	manifestDir := filepath.Join("dist", "__linkcheck__")
	if err := os.MkdirAll(manifestDir, 0o755); err != nil {
		return "", nil, fmt.Errorf("create manifest dir: %w", err)
	}

	var builder strings.Builder
	builder.WriteString("<!doctype html><html><body>\n")
	for _, target := range targets {
		escaped := html.EscapeString(target)
		builder.WriteString(`<a href="`)
		builder.WriteString(escaped)
		builder.WriteString(`">`)
		builder.WriteString(escaped)
		builder.WriteString("</a>\n")
	}
	builder.WriteString("</body></html>\n")

	manifestPath := filepath.Join(manifestDir, "index.html")
	if err := os.WriteFile(manifestPath, []byte(builder.String()), 0o644); err != nil {
		return "", nil, fmt.Errorf("write manifest page: %w", err)
	}

	cleanup := func() {
		_ = os.RemoveAll(manifestDir)
	}
	return fmt.Sprintf("http://127.0.0.1:%s/__linkcheck__/", port), cleanup, nil
}

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

func isMuffetReportExit(err error) bool {
	var exitErr *exec.ExitError
	return errors.As(err, &exitErr)
}

func validateJSONReport(path string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	if len(strings.TrimSpace(string(content))) == 0 {
		return errors.New("empty report")
	}

	var out any
	if err := json.Unmarshal(content, &out); err != nil {
		return err
	}
	return nil
}

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

func stopCommand(cmd *exec.Cmd, done <-chan error, timeout time.Duration) error {
	if cmd == nil || cmd.Process == nil {
		return nil
	}

	if runtime.GOOS == "windows" {
		_ = cmd.Process.Kill()
	} else {
		_ = cmd.Process.Signal(syscall.SIGTERM)
	}

	timer := time.NewTimer(timeout)
	defer timer.Stop()

	select {
	case <-done:
		return nil
	case <-timer.C:
		return cmd.Process.Kill()
	}
}

type processState struct {
	done chan struct{}
	mu   sync.Mutex
	errv error
}

func newProcessState() *processState {
	return &processState{done: make(chan struct{})}
}

func (p *processState) setErr(err error) {
	p.mu.Lock()
	p.errv = err
	p.mu.Unlock()
	close(p.done)
}

func (p *processState) err() error {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.errv
}
