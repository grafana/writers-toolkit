package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestInferRelativePrefixFromReports(t *testing.T) {
	reports := []pageReport{
		{URL: "http://127.0.0.1:3002/docs/writers-toolkit/whats-new/"},
		{URL: "http://127.0.0.1:3002/docs/writers-toolkit/review/test-documentation-changes/"},
	}
	changedFiles := []string{
		"docs/sources/whats-new.md",
		"docs/sources/review/test-documentation-changes/index.md",
	}

	got := inferRelativePrefixFromReports(reports, changedFiles, "docs/sources")
	want := "/docs/writers-toolkit/"
	if got != want {
		t.Fatalf("inferRelativePrefixFromReports() = %q, want %q", got, want)
	}
}

func TestInferRelativePrefixFromBrokenLinkTargets(t *testing.T) {
	reports := []pageReport{
		{
			URL: "http://127.0.0.1:3002/docs/writers-toolkit/contribute/",
			Links: []linkReport{
				{
					URL:   "http://127.0.0.1:3002/docs/writers-toolkit/review/test-documentation-changes/",
					Error: "404",
				},
			},
		},
	}
	changedFiles := []string{
		"docs/sources/review/test-documentation-changes/index.md",
	}

	got := inferRelativePrefixFromReports(reports, changedFiles, "docs/sources")
	want := "/docs/writers-toolkit/"
	if got != want {
		t.Fatalf("inferRelativePrefixFromReports() = %q, want %q", got, want)
	}
}

func TestCollectBrokenRowsIncludesMovedTargets(t *testing.T) {
	reports := []pageReport{
		{
			URL: "http://127.0.0.1:3002/docs/writers-toolkit/contribute/",
			Links: []linkReport{
				{
					URL:   "http://127.0.0.1:3002/docs/writers-toolkit/review/test-documentation-changes/",
					Error: "404",
				},
			},
		},
	}

	mappings := []mapping{{
		sourceDirectory: "docs/sources",
		relativePrefix:  "/docs/writers-toolkit/",
	}}

	fileToCandidates := map[string][]string{
		"docs/sources/review/test-documentation-changes/index.md":      candidatePagePaths("docs/sources/review/test-documentation-changes/index.md", mappings),
		"docs/sources/review/test-documentation-changes-test/index.md": candidatePagePaths("docs/sources/review/test-documentation-changes-test/index.md", mappings),
	}

	rows, filesWithBroken := collectBrokenRows(reports, fileToCandidates, mappings)

	if len(rows) != 1 {
		t.Fatalf("len(rows) = %d, want 1", len(rows))
	}
	if rows[0].File != "docs/sources/contribute/index.md" {
		t.Fatalf("row file = %q, want file containing the broken link", rows[0].File)
	}
	if rows[0].BrokenURL != "http://127.0.0.1:3002/docs/writers-toolkit/review/test-documentation-changes/" {
		t.Fatalf("row broken url = %q", rows[0].BrokenURL)
	}

	wantFiles := []string{"docs/sources/review/test-documentation-changes/index.md"}
	if !reflect.DeepEqual(filesWithBroken, wantFiles) {
		t.Fatalf("filesWithBroken = %#v, want %#v", filesWithBroken, wantFiles)
	}
}

func TestCollectBrokenRowsDeduplicatesOverlap(t *testing.T) {
	reports := []pageReport{
		{
			URL: "http://127.0.0.1:3002/docs/writers-toolkit/whats-new/",
			Links: []linkReport{
				{
					URL:   "http://127.0.0.1:3002/docs/writers-toolkit/whats-new/missing/",
					Error: "404",
				},
			},
		},
	}

	mappings := []mapping{{
		sourceDirectory: "docs/sources",
		relativePrefix:  "/docs/writers-toolkit/",
	}}
	fileToCandidates := map[string][]string{
		"docs/sources/whats-new.md": candidatePagePaths("docs/sources/whats-new.md", mappings),
	}

	rows, filesWithBroken := collectBrokenRows(reports, fileToCandidates, mappings)

	if len(rows) != 1 {
		t.Fatalf("len(rows) = %d, want 1", len(rows))
	}
	wantFiles := []string{"docs/sources/whats-new.md"}
	if !reflect.DeepEqual(filesWithBroken, wantFiles) {
		t.Fatalf("filesWithBroken = %#v, want %#v", filesWithBroken, wantFiles)
	}
}

func TestSourceFileForPageURL(t *testing.T) {
	mappings := []mapping{{
		sourceDirectory: "docs/sources",
		relativePrefix:  "/docs/writers-toolkit/",
	}}

	cases := []struct {
		name string
		url  string
		want string
	}{
		{
			name: "section path",
			url:  "http://127.0.0.1:3002/docs/writers-toolkit/contribute/",
			want: "docs/sources/contribute/index.md",
		},
		{
			name: "html path",
			url:  "http://127.0.0.1:3002/docs/writers-toolkit/contribute/unstyled.html",
			want: "docs/sources/contribute/unstyled.md",
		},
		{
			name: "root path",
			url:  "http://127.0.0.1:3002/docs/writers-toolkit/",
			want: "docs/sources/_index.md",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := sourceFileForPageURL(tc.url, mappings)
			if got != tc.want {
				t.Fatalf("sourceFileForPageURL() = %q, want %q", got, tc.want)
			}
		})
	}
}

func TestBuildCommentPluralization(t *testing.T) {
	comment := buildComment(commentInput{
		repo:                 "writers-toolkit",
		title:                "test",
		totalBroken:          1,
		changedDocsFileCount: 1,
		maxRows:              150,
	})

	if !strings.Contains(comment, "1 broken link total.") {
		t.Fatalf("expected singular 'broken link', got:\n%s", comment)
	}
}

func TestBuildCommentFallbackTableWhenNoChangedMatches(t *testing.T) {
	comment := buildComment(commentInput{
		repo:                   "writers-toolkit",
		title:                  "test",
		totalBroken:            2,
		changedDocsFileCount:   3,
		changedWithBrokenCount: 0,
		fallbackRows: []simpleBrokenRow{
			{
				PageURL:   "http://127.0.0.1:3002/docs/writers-toolkit/contribute/",
				BrokenURL: "http://127.0.0.1:3002/docs/writers-toolkit/review/test-documentation-changes/",
			},
			{
				PageURL:   "http://127.0.0.1:3002/docs/writers-toolkit/get-started/",
				BrokenURL: "http://127.0.0.1:3002/docs/writers-toolkit/review/test-documentation-changes/",
			},
		},
		maxRows: 150,
	})

	if strings.Contains(comment, "No broken links were detected on changed docs files in this PR") {
		t.Fatalf("did not expect old no-changed-match message, got:\n%s", comment)
	}
	if !strings.Contains(comment, "| Page | Broken link |") {
		t.Fatalf("expected fallback table header, got:\n%s", comment)
	}
	if !strings.Contains(comment, "See more in the logs.") {
		t.Fatalf("expected logs hint, got:\n%s", comment)
	}
}
