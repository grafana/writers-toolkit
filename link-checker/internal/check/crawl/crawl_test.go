package crawl

import "testing"

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
		if got := ShouldCrawlHTMLFile(tc.name); got != tc.want {
			t.Fatalf("ShouldCrawlHTMLFile(%q) = %v, want %v", tc.name, got, tc.want)
		}
	}
}
