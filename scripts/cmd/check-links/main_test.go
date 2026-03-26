package main

import "testing"

func TestExtractPageDataCapturesSourcePathAndIDs(t *testing.T) {
	pageURL := "http://127.0.0.1:3002/docs/writers-toolkit/contribute/"
	body := `
<html>
<head><script>window.Path="docs/writers-toolkit/contribute/index.md"</script></head>
<body>
  <a href="/docs/writers-toolkit/review/test-documentation-changes/">one</a>
  <a href="/docs/writers-toolkit/review/test-documentation-changes/">duplicate</a>
  <a href="https://grafana.com/docs/grafana/latest/developers/cla/">two</a>
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
  <div x-data='{"html":"\u003ca href=\"#resolver\"\u003elink\u003c/a\u003e"}'></div>
</body>
</html>
`

	_, links := extractPageData(pageURL, body, "3002")
	if len(links) != 0 {
		t.Fatalf("len(links) = %d, want 0; links=%#v", len(links), links)
	}
}

func TestShouldCrawlHTMLFile(t *testing.T) {
	cases := []struct {
		name string
		want bool
	}{
		{name: "index.html", want: true},
		{name: "unstyled.html", want: false},
		{name: "index.md", want: false},
	}

	for _, tc := range cases {
		if got := shouldCrawlHTMLFile(tc.name); got != tc.want {
			t.Fatalf("shouldCrawlHTMLFile(%q) = %v, want %v", tc.name, got, tc.want)
		}
	}
}
