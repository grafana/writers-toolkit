package comment

import (
	"os"
	"path/filepath"
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
			URL:        "http://127.0.0.1:3002/docs/writers-toolkit/contribute/",
			SourcePath: "docs/sources/contribute/index.md",
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

func TestCandidatePagePathsNestedUnderscoreIndex(t *testing.T) {
	mappings := []mapping{{
		sourceDirectory: "docs/sources",
		relativePrefix:  "/docs/writers-toolkit/",
	}}

	got := candidatePagePaths("docs/sources/contribute/_index.md", mappings)
	want := []string{"/docs/writers-toolkit/contribute/"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("candidatePagePaths() = %#v, want %#v", got, want)
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

func TestSourceFileForReportPrefersSourcePath(t *testing.T) {
	mappings := []mapping{{
		sourceDirectory: "docs/sources",
		relativePrefix:  "/docs/writers-toolkit/",
	}}

	report := pageReport{
		URL:        "http://127.0.0.1:3002/docs/writers-toolkit/whats-new/",
		SourcePath: "docs/writers-toolkit/whats-new.md",
	}

	got := sourceFileForReport(report, mappings)
	want := "docs/sources/whats-new.md"
	if got != want {
		t.Fatalf("sourceFileForReport() = %q, want %q", got, want)
	}
}

func TestSourceFileForReportLeavesNonMatchingSourcePath(t *testing.T) {
	mappings := []mapping{{
		sourceDirectory: "docs/sources",
		relativePrefix:  "/docs/writers-toolkit/",
	}}

	report := pageReport{
		URL:        "http://127.0.0.1:3002/docs/writers-toolkit/whats-new/",
		SourcePath: "content/docs/other-product/index.md",
	}

	got := sourceFileForReport(report, mappings)
	want := "content/docs/other-product/index.md"
	if got != want {
		t.Fatalf("sourceFileForReport() = %q, want %q", got, want)
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

func TestBuildCommentShowsAllBrokenRows(t *testing.T) {
	comment := buildComment(commentInput{
		repo:        "writers-toolkit",
		title:       "test",
		totalBroken: 2,
		rows: []brokenRow{
			{
				File:      "docs/sources/contribute/_index.md",
				PageURL:   "http://127.0.0.1:3002/docs/writers-toolkit/contribute/",
				BrokenURL: "http://127.0.0.1:3002/docs/writers-toolkit/review/test-documentation-changes/",
				Error:     "404",
			},
			{
				File:      "docs/sources/get-started/_index.md",
				PageURL:   "http://127.0.0.1:3002/docs/writers-toolkit/get-started/",
				BrokenURL: "http://127.0.0.1:3002/docs/writers-toolkit/review/test-documentation-changes/",
				Error:     "404",
			},
		},
		maxRows: 150,
	})

	if !strings.Contains(comment, "Broken links found in this build:") {
		t.Fatalf("expected build-wide broken-links message, got:\n%s", comment)
	}
	if !strings.Contains(comment, "| File") || !strings.Contains(comment, "Broken link") {
		t.Fatalf("expected full broken-links table header, got:\n%s", comment)
	}
	if strings.Contains(comment, "| Page") {
		t.Fatalf("did not expect page column in table, got:\n%s", comment)
	}
	if strings.Contains(comment, "See more in the logs.") {
		t.Fatalf("did not expect fallback logs hint, got:\n%s", comment)
	}
}

func TestRenderMarkdownTablePadsColumns(t *testing.T) {
	got := renderMarkdownTable(
		[]string{"A", "BBB"},
		[][]string{
			{"xx", "y"},
			{"z", "long"},
		},
	)

	want := strings.Join([]string{
		"| A  | BBB  |",
		"| -- | ---- |",
		"| xx | y    |",
		"| z  | long |",
		"",
	}, "\n")
	if got != want {
		t.Fatalf("renderMarkdownTable() =\n%s\nwant:\n%s", got, want)
	}
}

func TestFilterReportsForCommentNoFiltering(t *testing.T) {
	reports := []pageReport{
		{
			URL: "http://127.0.0.1:3002/docs/writers-toolkit/contribute/",
			Links: []linkReport{
				{
					URL:   "http://127.0.0.1:3002/docs/writers-toolkit/review/unstyled.html",
					Error: "404",
				},
				{
					URL:   "http://127.0.0.1:3002/docs/writers-toolkit/review/missing-page/",
					Error: "404",
				},
			},
		},
		{
			URL: "http://127.0.0.1:3002/docs/writers-toolkit/tutorials/unstyled.html",
			Links: []linkReport{
				{
					URL:   "http://127.0.0.1:3002/docs/writers-toolkit/review/missing-page/",
					Error: "404",
				},
			},
		},
	}

	filtered := filterReportsForComment(reports)
	if len(filtered) != len(reports) {
		t.Fatalf("len(filtered) = %d, want %d", len(filtered), len(reports))
	}
	if len(filtered[0].Links) != len(reports[0].Links) {
		t.Fatalf("len(filtered[0].Links) = %d, want %d", len(filtered[0].Links), len(reports[0].Links))
	}
}

func TestSourceLinkLocatorFindsExactURL(t *testing.T) {
	dir := t.TempDir()
	filePath := filepath.Join(dir, "cla.md")
	content := "alpha\n- [CLA](https://grafana.com/docs/grafana/latest/developers/cla/)\n"
	if err := os.WriteFile(filePath, []byte(content), 0o644); err != nil {
		t.Fatalf("write file: %v", err)
	}

	locator := newSourceLinkLocator()
	line, column, ok := locator.find(filePath, linkReport{
		URL: "http://127.0.0.1:3002/docs/grafana/latest/developers/cla/",
		Raw: "https://grafana.com/docs/grafana/latest/developers/cla/",
	})
	if !ok {
		t.Fatal("expected to find source location")
	}
	if line != 2 || column != 9 {
		t.Fatalf("location = %d:%d, want 2:9", line, column)
	}
}

func TestSourceLinkLocatorMatchesGrafanaVersionPlaceholder(t *testing.T) {
	dir := t.TempDir()
	filePath := filepath.Join(dir, "task.md")
	content := "- [Create](https://grafana.com/docs/grafana/<GRAFANA_VERSION>/dashboards/build-dashboards/create-dashboard/)\n"
	if err := os.WriteFile(filePath, []byte(content), 0o644); err != nil {
		t.Fatalf("write file: %v", err)
	}

	locator := newSourceLinkLocator()
	line, column, ok := locator.find(filePath, linkReport{
		URL: "http://127.0.0.1:3002/docs/grafana/latest/dashboards/build-dashboards/create-dashboard/",
	})
	if !ok {
		t.Fatal("expected to find source location")
	}
	if line != 1 || column != 12 {
		t.Fatalf("location = %d:%d, want 1:12", line, column)
	}
}

func TestBuildCommentShowsFileLineColumn(t *testing.T) {
	comment := buildComment(commentInput{
		repo:        "writers-toolkit",
		title:       "test",
		totalBroken: 1,
		rows: []brokenRow{
			{
				File:      "docs/sources/review/cla-assistant/index.md",
				Line:      42,
				Column:    7,
				BrokenURL: "http://127.0.0.1:3002/docs/grafana/latest/developers/cla/",
				Error:     "404",
			},
		},
	})

	if !strings.Contains(comment, "`docs/sources/review/cla-assistant/index.md:42:7`") {
		t.Fatalf("expected file column to include line/column, got:\n%s", comment)
	}
}

func TestBuildCommentOmitsSourceCheckoutPrefix(t *testing.T) {
	comment := buildComment(commentInput{
		repo:        "writers-toolkit",
		title:       "test",
		totalBroken: 1,
		rows: []brokenRow{
			{
				File:      "source-files/docs/sources/contribute/_index.md",
				Line:      125,
				Column:    81,
				BrokenURL: "http://127.0.0.1:3002/docs/writers-toolkit/review/test-documentation-changes/",
				Error:     "404",
			},
		},
	})

	if strings.Contains(comment, "`source-files/docs/sources/contribute/_index.md:125:81`") {
		t.Fatalf("did not expect source checkout prefix in comment, got:\n%s", comment)
	}
	if !strings.Contains(comment, "`docs/sources/contribute/_index.md:125:81`") {
		t.Fatalf("expected repo-relative file path in comment, got:\n%s", comment)
	}
	if !strings.Contains(comment, "`/docs/writers-toolkit/review/test-documentation-changes/`") {
		t.Fatalf("expected local preview URL to render as path, got:\n%s", comment)
	}
}

func TestBuildCommentPreservesExternalBrokenURL(t *testing.T) {
	comment := buildComment(commentInput{
		repo:        "writers-toolkit",
		title:       "test",
		totalBroken: 1,
		rows: []brokenRow{
			{
				File:      "docs/sources/contribute/_index.md",
				BrokenURL: "https://grafana.com/docs/test/writers-toolkit/write/shortcodes/",
				Error:     "404",
			},
		},
	})

	if !strings.Contains(comment, "`https://grafana.com/docs/test/writers-toolkit/write/shortcodes/`") {
		t.Fatalf("expected external URL to remain unchanged, got:\n%s", comment)
	}
}
