package nginx

import (
	"bufio"
	"context"
	_ "embed"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"
	"syscall"
	"time"
)

const defaultRelativePrefix = "/docs/"

//go:embed nginx.local.conf
var localTemplate string

// Server represents a running nginx process managed by this package.
type Server struct {
	cmd   *exec.Cmd
	state *processState
}

// PrepareConfig renders and writes local nginx config artifacts and logs.
func PrepareConfig(nginxPort string, relativePrefixes []string, sha string) error {
	if err := os.MkdirAll("dist", 0o755); err != nil {
		return fmt.Errorf("create dist dir: %w", err)
	}
	if err := os.MkdirAll("run", 0o755); err != nil {
		return fmt.Errorf("create run dir: %w", err)
	}

	rendered, err := renderConfig(localTemplate, nginxPort, relativePrefixes, sha)
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

// TestConfig validates the generated nginx configuration.
func TestConfig() error {
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

// Start starts nginx in foreground mode and tracks its process state.
func Start(ctx context.Context) (*Server, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("get working directory: %w", err)
	}

	cmd := exec.CommandContext(ctx, "nginx", "-p", cwd, "-c", "nginx.conf", "-g", "daemon off;")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if runtime.GOOS != "windows" {
		cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	}

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("start nginx: %w", err)
	}

	state := &processState{done: make(chan struct{})}
	go func() {
		state.setErr(cmd.Wait())
	}()

	return &Server{cmd: cmd, state: state}, nil
}

// Done returns a channel closed when nginx exits.
func (s *Server) Done() <-chan struct{} {
	if s == nil || s.state == nil {
		closed := make(chan struct{})
		close(closed)
		return closed
	}
	return s.state.done
}

// Err returns the terminal process error from nginx.
func (s *Server) Err() error {
	if s == nil || s.state == nil {
		return nil
	}
	return s.state.err()
}

// Stop attempts graceful then forced process-group shutdown.
func (s *Server) Stop(timeout time.Duration) error {
	if s == nil {
		return nil
	}
	return stopProcessGroup(s.cmd, s.state.done, timeout)
}

// WaitReady waits until nginx responds or exits/fails.
func WaitReady(ctx context.Context, readyURL string, server *Server, timeout time.Duration) error {
	client := &httpClient{timeout: 2 * time.Second}
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
		case <-server.Done():
			if err := server.Err(); err != nil {
				return fmt.Errorf("nginx exited before ready: %w", err)
			}
			return errors.New("nginx exited before ready")
		case <-ticker.C:
			status, err := client.getStatus(readyURL)
			if err != nil {
				continue
			}
			if status > 0 {
				return nil
			}
		}
	}
}

// StartLogStreaming starts asynchronous nginx error/access log tailing.
func StartLogStreaming(parent context.Context, logRequests bool, stdout, stderr io.Writer) func() {
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

// buildServerConfig builds location blocks that map relative prefixes to their dist dir.
// For example, /docs/writers-toolkit/ gets aliased to dist/docs/writers-toolkit/.
func buildServerConfig(relativePrefixes []string, distRoot, sha string) string {
	var builder strings.Builder
	_, _ = fmt.Fprintf(&builder, "add_header 'Build' '%s';\n\n", sha)

	hasDocsPrefix := false
	for _, relativePrefix := range relativePrefixes {
		normalizedPrefix := normalizeRelativePrefix(relativePrefix)
		prefixWithoutTrailingSlash := strings.TrimSuffix(normalizedPrefix, "/")
		if strings.HasPrefix(prefixWithoutTrailingSlash, "/docs/") {
			hasDocsPrefix = true
		}

		_, _ = fmt.Fprintf(&builder, `location = %s {
  return 301 %s;
}

location ^~ %s {
  alias %s%s;
}

		`, prefixWithoutTrailingSlash, normalizedPrefix, normalizedPrefix, distRoot, normalizedPrefix)
	}

	if hasDocsPrefix {
		builder.WriteString(`location ^~ /docs/ {
  proxy_pass https://grafana.com/docs/;
}
`)
	}

	return builder.String()
}

// renderConfig injects runtime values into the local check-links template.
func renderConfig(config, nginxPort string, relativePrefixes []string, sha string) (string, error) {
	replacements := []struct {
		old string
		new string
	}{
		{"listen 80;", fmt.Sprintf("listen %s;", nginxPort)},
		{"include /etc/nginx/build.conf;", buildServerConfig(relativePrefixes, "dist", sha)},
		{"include /etc/nginx/locations.conf;", "include deploy-preview/locations.conf;"},
		{"include /etc/nginx/redirects.conf;", "include deploy-preview/redirects.conf;"},
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
	}

	if runtime.GOOS == "windows" {
		_ = cmd.Process.Kill()
	} else {
		_ = syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
	}

	select {
	case <-done:
		return nil
	case <-time.After(2 * time.Second):
		return errors.New("timed out waiting for nginx to exit")
	}
}

type processState struct {
	done chan struct{}
	mu   sync.RWMutex
	errV error
}

func (p *processState) setErr(err error) {
	p.mu.Lock()
	p.errV = err
	p.mu.Unlock()
	close(p.done)
}

func (p *processState) err() error {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.errV
}

type httpClient struct {
	timeout time.Duration
}

func (c *httpClient) getStatus(url string) (int, error) {
	client := &http.Client{Timeout: c.timeout}
	response, err := client.Get(url)
	if err != nil {
		return 0, err
	}
	_ = response.Body.Close()
	return response.StatusCode, nil
}
