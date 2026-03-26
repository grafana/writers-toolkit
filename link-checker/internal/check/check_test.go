package check

import "testing"

func TestExtractPageDataCapturesSourcePathAndIDs(t *testing.T) {
	pageURL := "http://127.0.0.1:3002/docs/writers-toolkit/contribute/"
	body := `
<html>
<head><script>window.Path="docs/writers-toolkit/contribute/index.md"</script></head>
<body>
  <div id="doc-article-text">
    <a href="/docs/writers-toolkit/review/test-documentation-changes/">one</a>
    <a href="/docs/writers-toolkit/review/test-documentation-changes/">duplicate</a>
    <a href="https://grafana.com/docs/grafana/latest/developers/cla/">two</a>
  </div>
</body>
</html>
`

	sourcePath, links := extractPageData(pageURL, body, "3002")
	if sourcePath != "docs/writers-toolkit/contribute/index.md" {
		t.Fatalf("sourcePath = %q, want %q", sourcePath, "docs/writers-toolkit/contribute/index.md")
	}
	if len(links) != 2 {
		t.Fatalf("len(links) = %d, want 2", len(links))
	}

	if links[0].ID != pageURL+"#1" {
		t.Fatalf("links[0].ID = %q, want %q", links[0].ID, pageURL+"#1")
	}
	if links[0].URL != "http://127.0.0.1:3002/docs/writers-toolkit/review/test-documentation-changes/" {
		t.Fatalf("links[0].URL = %q", links[0].URL)
	}
	if links[0].Raw != "/docs/writers-toolkit/review/test-documentation-changes/" {
		t.Fatalf("links[0].Raw = %q", links[0].Raw)
	}

	if links[1].ID != pageURL+"#2" {
		t.Fatalf("links[1].ID = %q, want %q", links[1].ID, pageURL+"#2")
	}
	if links[1].URL != "https://grafana.com/docs/grafana/latest/developers/cla/" {
		t.Fatalf("links[1].URL = %q", links[1].URL)
	}
}

func TestExtractPageDataIgnoresEscapedHrefInEmbeddedJSON(t *testing.T) {
	pageURL := "http://127.0.0.1:3002/docs/writers-toolkit/write/shortcodes/"
	body := `
<html>
<body>
  <div id="doc-article-text">
    <div x-data='{"html":"\u003ca href=\"#resolver\"\u003elink\u003c/a\u003e"}'></div>
  </div>
</body>
</html>
`

	_, links := extractPageData(pageURL, body, "3002")
	if len(links) != 0 {
		t.Fatalf("len(links) = %d, want 0; links=%#v", len(links), links)
	}
}

func TestExtractPageDataIgnoresBoundTemplateAttributes(t *testing.T) {
	pageURL := "http://127.0.0.1:3002/docs/grafana/latest/"
	body := `
<html>
<body>
  <div id="doc-article-text">
    <div :src="header.image.src"></div>
    <a :href="link.href">bad</a>
    <img x-bind:src="error.image.src" />
    <a x-bind:href="link.href">bad2</a>
    <a href="/docs/grafana/latest/real-link/">good</a>
  </div>
</body>
</html>
`

	_, links := extractPageData(pageURL, body, "3002")
	if len(links) != 1 {
		t.Fatalf("len(links) = %d, want 1; links=%#v", len(links), links)
	}
	if links[0].URL != "http://127.0.0.1:3002/docs/grafana/latest/real-link/" {
		t.Fatalf("links[0].URL = %q", links[0].URL)
	}
}

func TestExtractPageDataOnlyChecksDocArticleText(t *testing.T) {
	pageURL := "http://127.0.0.1:3002/docs/grafana/latest/"
	body := `
<html>
<body>
  <a href="/docs/grafana/latest/outside/">outside</a>
  <div id="doc-article-text">
    <a href="/docs/grafana/latest/inside/">inside</a>
  </div>
</body>
</html>
`

	_, links := extractPageData(pageURL, body, "3002")
	if len(links) != 1 {
		t.Fatalf("len(links) = %d, want 1; links=%#v", len(links), links)
	}
	if links[0].URL != "http://127.0.0.1:3002/docs/grafana/latest/inside/" {
		t.Fatalf("links[0].URL = %q", links[0].URL)
	}
}

func TestExtractPageDataReturnsNoLinksWithoutDocArticleText(t *testing.T) {
	pageURL := "http://127.0.0.1:3002/docs/grafana/latest/"
	body := `
<html>
<body>
  <a href="/docs/grafana/latest/outside/">outside</a>
</body>
</html>
`

	_, links := extractPageData(pageURL, body, "3002")
	if len(links) != 0 {
		t.Fatalf("len(links) = %d, want 0; links=%#v", len(links), links)
	}
}
