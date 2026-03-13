package main

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"regexp"
	"runtime"
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

	relativePrefix := resolveRelativePrefix()
	if opts.baseURL == "" {
		opts.baseURL = baseURLForPrefix(relativePrefix, opts.nginxPort)
	}
	if opts.include == "" {
		opts.include = includePatternForPrefix(relativePrefix, opts.nginxPort)
	}
	if opts.exclude == "" {
		opts.exclude = defaultExcludePattern(opts.nginxPort)
	}

	return opts
}

func resolveRelativePrefix() string {
	if prefix := strings.TrimSpace(os.Getenv("RELATIVE_PREFIX")); prefix != "" {
		return normalizeRelativePrefix(prefix)
	}

	sources := strings.TrimSpace(os.Getenv("SOURCES"))
	if sources == "" {
		return defaultRelativePrefix
	}

	var configs []sourceConfig
	if err := json.Unmarshal([]byte(sources), &configs); err != nil {
		return defaultRelativePrefix
	}
	for _, config := range configs {
		if strings.TrimSpace(config.RelativePrefix) != "" {
			return normalizeRelativePrefix(config.RelativePrefix)
		}
	}

	return defaultRelativePrefix
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

func includePatternForPrefix(prefix, port string) string {
	normalized := strings.TrimSuffix(normalizeRelativePrefix(prefix), "/")
	if normalized == "/docs" || strings.HasPrefix(normalized, "/docs/") {
		return fmt.Sprintf(
			"http://(?:127\\.0\\.0\\.1|localhost):%s/docs(?:/.*)?(?:\\?.*)?$",
			port,
		)
	}
	return fmt.Sprintf(
		"http://(?:127\\.0\\.0\\.1|localhost):%s%s(?:/.*)?(?:\\?.*)?$",
		port,
		regexp.QuoteMeta(normalized),
	)
}

func defaultExcludePattern(port string) string {
	return fmt.Sprintf("http://(?:127\\.0\\.0\\.1|localhost):%s/(?:static|media|api|connect)", port)
}

func run(ctx context.Context, opts options) error {
	if err := prepareNginxConfig(opts.nginxPort); err != nil {
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

func prepareNginxConfig(nginxPort string) error {
	if err := os.MkdirAll("dist", 0o755); err != nil {
		return fmt.Errorf("create dist dir: %w", err)
	}
	if err := os.MkdirAll("run", 0o755); err != nil {
		return fmt.Errorf("create run dir: %w", err)
	}

	buildHeader := fmt.Sprintf("add_header 'Build' '%s';", os.Getenv("SHA"))
	if err := os.WriteFile("dist/build.conf", []byte(buildHeader), 0o644); err != nil {
		return fmt.Errorf("write dist/build.conf: %w", err)
	}

	templateBytes, err := os.ReadFile("deploy-preview/nginx.conf")
	if err != nil {
		return fmt.Errorf("read deploy-preview/nginx.conf: %w", err)
	}

	rendered, err := renderLocalNginxConfig(string(templateBytes), nginxPort)
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

func renderLocalNginxConfig(config, nginxPort string) (string, error) {
	replacements := []struct {
		old string
		new string
	}{
		{"pid /run/nginx.pid;", "pid run/nginx.pid;"},
		{"access_log /var/log/nginx/access.log;", "access_log access.log;"},
		{"error_log /var/log/nginx/error.log;", "error_log error.log;"},
		{"listen 80;", fmt.Sprintf("listen %s;", nginxPort)},
		{"include /etc/nginx/build.conf;", "include dist/build.conf;"},
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
	fmt.Fprintf(os.Stdout, "running muffet crawl from: %s\n", opts.baseURL)

	return runMuffetOnce(ctx, opts, nginxState, opts.baseURL, opts.outputPath)
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
	args = append(args, "-f", "-b", "9999", "--format=json")
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
