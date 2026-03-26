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

func TestBuildLocalURLEscapesReservedPathCharacters(t *testing.T) {
	got := buildLocalURL("3002", "/docs/grafana/latest/manage-rbac-roles#create-custom-roles/")
	want := "http://127.0.0.1:3002/docs/grafana/latest/manage-rbac-roles%23create-custom-roles/"
	if got != want {
		t.Fatalf("buildLocalURL() = %q, want %q", got, want)
	}
}
