package nginx

import (
	"strings"
	"testing"
)

func TestRenderConfigIncludesRedirects(t *testing.T) {
	rendered, err := renderConfig(localTemplate, "3002", []string{"/docs/writers-toolkit/"}, "test-sha")
	if err != nil {
		t.Fatalf("renderConfig() error = %v", err)
	}

	if !strings.Contains(rendered, "include deploy-preview/locations.conf;") {
		t.Fatalf("expected rendered config to include locations.conf, got:\n%s", rendered)
	}
	if !strings.Contains(rendered, "include deploy-preview/redirects.conf;") {
		t.Fatalf("expected rendered config to include redirects.conf, got:\n%s", rendered)
	}
}
