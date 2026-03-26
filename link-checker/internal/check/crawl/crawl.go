package crawl

import (
	"errors"
	"fmt"
	"net"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const defaultRelativePrefix = "/docs/"

// CollectSourcePageURLs collects rendered page URLs from dist html output.
func CollectSourcePageURLs(relativePrefixes []string, port string) ([]string, error) {
	urls := make([]string, 0)
	seen := map[string]struct{}{}

	for _, relativePrefix := range relativePrefixes {
		prefixRoot := filepath.Join("dist", filepath.FromSlash(strings.TrimPrefix(normalizeRelativePrefix(relativePrefix), "/")))
		err := filepath.Walk(prefixRoot, func(path string, info os.FileInfo, walkErr error) error {
			if walkErr != nil {
				return walkErr
			}
			if info.IsDir() || !ShouldCrawlHTMLFile(info.Name()) {
				return nil
			}

			urlPath, ok := htmlFilePathToURLPath(path)
			if !ok {
				return nil
			}

			sourceURL := buildLocalURL(port, urlPath)
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

// ShouldCrawlHTMLFile reports whether an HTML file should be crawled.
func ShouldCrawlHTMLFile(name string) bool {
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

// buildLocalURL builds a local preview URL and escapes reserved path characters.
func buildLocalURL(port, path string) string {
	return (&url.URL{
		Scheme: "http",
		Host:   net.JoinHostPort("127.0.0.1", port),
		Path:   path,
	}).String()
}
