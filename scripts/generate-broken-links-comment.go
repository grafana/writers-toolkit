package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"net/url"
	"os"
	"path"
	"sort"
	"strings"
)

type pageReport struct {
	URL   string       `json:"url"`
	Links []linkReport `json:"links"`
}

type linkReport struct {
	URL   string `json:"url"`
	Error string `json:"error"`
}

type sourceConfig struct {
	SourceDirectory string `json:"source_directory"`
	RelativePrefix  string `json:"relative_prefix"`
}

type mapping struct {
	sourceDirectory string
	relativePrefix  string
}

type brokenRow struct {
	File      string
	PageURL   string
	BrokenURL string
	Error     string
}

const (
	defaultSourceDirectory        = "docs/sources"
	generateDefaultRelativePrefix = "/docs/"
)

// main parses inputs, correlates changed files with broken links, and writes a PR comment body.
func main() {
	var linksPath string
	var changedFilesPath string
	var outputPath string
	var title string
	var repo string
	var artifactURL string
	var sourceDirectory string
	var relativePrefix string
	var sourcesJSON string
	var maxRows int

	flag.StringVar(&linksPath, "links", "links.json", "Path to links JSON report")
	flag.StringVar(&changedFilesPath, "changed-files", "changed-files.txt", "Path to changed files list")
	flag.StringVar(&outputPath, "output", "broken-links-comment.md", "Path to write comment body")
	flag.StringVar(&title, "title", "", "PR title for comment heading")
	flag.StringVar(&repo, "repo", "", "Repository slug fragment for hidden comment marker")
	flag.StringVar(&artifactURL, "artifact-url", "", "Uploaded artifact URL")
	flag.StringVar(&sourceDirectory, "source-directory", strings.TrimSpace(os.Getenv("SOURCE_DIRECTORY")), "Legacy single source directory")
	flag.StringVar(&relativePrefix, "relative-prefix", strings.TrimSpace(os.Getenv("RELATIVE_PREFIX")), "Legacy single relative prefix")
	flag.StringVar(&sourcesJSON, "sources-json", strings.TrimSpace(os.Getenv("SOURCES")), "JSON array of source mappings")
	flag.IntVar(&maxRows, "max-rows", 150, "Maximum number of table rows to render")
	flag.Parse()

	reports, err := readReports(linksPath)
	if err != nil {
		exitWithError(err)
	}

	changedFiles, err := readChangedFiles(changedFilesPath)
	if err != nil {
		exitWithError(err)
	}

	mappings := resolveMappings(sourcesJSON, sourceDirectory, relativePrefix)
	fileToCandidates := make(map[string][]string)
	for _, file := range changedFiles {
		candidates := candidatePagePaths(file, mappings)
		if len(candidates) == 0 {
			continue
		}
		fileToCandidates[file] = candidates
	}

	reportsByPath := make(map[string]pageReport, len(reports))
	for _, report := range reports {
		reportPath := reportPath(report.URL)
		if reportPath == "" {
			continue
		}
		reportsByPath[reportPath] = report
	}

	rows := make([]brokenRow, 0)
	seenPageByFile := make(map[string]map[string]struct{})
	uniqueBrokenOnChangedPages := 0
	pageSeenForChangedTotals := map[string]struct{}{}

	changedFilesWithBroken := make([]string, 0)
	for file, candidates := range fileToCandidates {
		fileHasBroken := false
		for _, candidate := range candidates {
			report, ok := reportsByPath[candidate]
			if !ok {
				continue
			}
			if _, ok := seenPageByFile[file]; !ok {
				seenPageByFile[file] = map[string]struct{}{}
			}
			if _, dup := seenPageByFile[file][candidate]; dup {
				continue
			}
			seenPageByFile[file][candidate] = struct{}{}

			if len(report.Links) == 0 {
				continue
			}
			fileHasBroken = true
			if _, seen := pageSeenForChangedTotals[candidate]; !seen {
				pageSeenForChangedTotals[candidate] = struct{}{}
				uniqueBrokenOnChangedPages += len(report.Links)
			}
			for _, link := range report.Links {
				rows = append(rows, brokenRow{
					File:      file,
					PageURL:   report.URL,
					BrokenURL: link.URL,
					Error:     link.Error,
				})
			}
		}
		if fileHasBroken {
			changedFilesWithBroken = append(changedFilesWithBroken, file)
		}
	}

	sort.Strings(changedFilesWithBroken)
	sort.Slice(rows, func(i, j int) bool {
		if rows[i].File != rows[j].File {
			return rows[i].File < rows[j].File
		}
		if rows[i].PageURL != rows[j].PageURL {
			return rows[i].PageURL < rows[j].PageURL
		}
		return rows[i].BrokenURL < rows[j].BrokenURL
	})

	totalBroken := 0
	for _, report := range reports {
		totalBroken += len(report.Links)
	}

	comment := buildComment(commentInput{
		repo:                   repo,
		title:                  title,
		artifactURL:            artifactURL,
		totalBroken:            totalBroken,
		changedDocsFileCount:   len(fileToCandidates),
		changedWithBrokenCount: len(changedFilesWithBroken),
		brokenOnChangedPages:   uniqueBrokenOnChangedPages,
		rows:                   rows,
		maxRows:                maxRows,
	})

	if err := os.WriteFile(outputPath, []byte(comment), 0o644); err != nil {
		exitWithError(fmt.Errorf("write %s: %w", outputPath, err))
	}
}

// exitWithError prints a formatted error and terminates the process.
func exitWithError(err error) {
	fmt.Fprintf(os.Stderr, "generate broken links comment failed: %v\n", err)
	os.Exit(1)
}

// readReports loads the JSON link checker output from disk.
func readReports(path string) ([]pageReport, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", path, err)
	}

	var reports []pageReport
	if err := json.Unmarshal(content, &reports); err != nil {
		return nil, fmt.Errorf("parse %s: %w", path, err)
	}
	return reports, nil
}

// readChangedFiles returns a sorted, de-duplicated list of changed file paths.
func readChangedFiles(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open %s: %w", path, err)
	}
	defer func() { _ = file.Close() }()

	scanner := bufio.NewScanner(file)
	files := make([]string, 0)
	seen := map[string]struct{}{}
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		if _, ok := seen[line]; ok {
			continue
		}
		seen[line] = struct{}{}
		files = append(files, line)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("read %s: %w", path, err)
	}
	sort.Strings(files)
	return files, nil
}

// resolveMappings builds source-directory to relative-prefix mappings from inputs.
func resolveMappings(sourcesJSON, sourceDirectory, relativePrefix string) []mapping {
	mappings := make([]mapping, 0)
	seen := map[string]struct{}{}

	if strings.TrimSpace(sourcesJSON) != "" {
		var configs []sourceConfig
		if err := json.Unmarshal([]byte(sourcesJSON), &configs); err == nil {
			for _, config := range configs {
				sd := normalizeSourceDirectory(config.SourceDirectory)
				rp := normalizeRelativePrefixForComment(config.RelativePrefix)
				key := sd + "|" + rp
				if _, ok := seen[key]; ok {
					continue
				}
				seen[key] = struct{}{}
				mappings = append(mappings, mapping{sourceDirectory: sd, relativePrefix: rp})
			}
		}
	}

	if len(mappings) == 0 {
		sd := normalizeSourceDirectory(sourceDirectory)
		rp := normalizeRelativePrefixForComment(relativePrefix)
		mappings = append(mappings, mapping{sourceDirectory: sd, relativePrefix: rp})
	}

	return mappings
}

// normalizeSourceDirectory trims and defaults a source directory path.
func normalizeSourceDirectory(directory string) string {
	directory = strings.TrimSpace(strings.Trim(directory, "/"))
	if directory == "" {
		return defaultSourceDirectory
	}
	return directory
}

// normalizeRelativePrefixForComment ensures a relative prefix has leading and trailing slashes.
func normalizeRelativePrefixForComment(prefix string) string {
	prefix = strings.TrimSpace(prefix)
	if prefix == "" {
		return generateDefaultRelativePrefix
	}
	if !strings.HasPrefix(prefix, "/") {
		prefix = "/" + prefix
	}
	if !strings.HasSuffix(prefix, "/") {
		prefix += "/"
	}
	return prefix
}

// candidatePagePaths maps a changed markdown file path to possible rendered page paths.
func candidatePagePaths(file string, mappings []mapping) []string {
	if !strings.HasSuffix(file, ".md") {
		return nil
	}

	candidates := make([]string, 0, 3)
	seen := map[string]struct{}{}

	for _, m := range mappings {
		prefix := m.sourceDirectory + "/"
		if file != m.sourceDirectory && !strings.HasPrefix(file, prefix) {
			continue
		}

		rel := strings.TrimPrefix(file, prefix)
		if file == m.sourceDirectory {
			rel = ""
		}
		if rel == "" {
			continue
		}

		rel = strings.TrimPrefix(rel, "/")
		if rel == "_index.md" {
			candidate := normalizeReportPath(m.relativePrefix)
			if _, ok := seen[candidate]; !ok {
				seen[candidate] = struct{}{}
				candidates = append(candidates, candidate)
			}
			continue
		}

		if strings.HasSuffix(rel, "/index.md") {
			suffix := strings.TrimSuffix(rel, "/index.md")
			candidate := normalizeReportPath(path.Join(m.relativePrefix, suffix) + "/")
			if _, ok := seen[candidate]; !ok {
				seen[candidate] = struct{}{}
				candidates = append(candidates, candidate)
			}
			continue
		}

		if strings.HasSuffix(rel, ".md") {
			base := strings.TrimSuffix(rel, ".md")
			pretty := normalizeReportPath(path.Join(m.relativePrefix, base) + "/")
			if _, ok := seen[pretty]; !ok {
				seen[pretty] = struct{}{}
				candidates = append(candidates, pretty)
			}

			htmlPath := normalizeReportPath(path.Join(m.relativePrefix, base) + ".html")
			if _, ok := seen[htmlPath]; !ok {
				seen[htmlPath] = struct{}{}
				candidates = append(candidates, htmlPath)
			}
		}
	}

	sort.Strings(candidates)
	return candidates
}

// reportPath extracts and normalizes the path portion from a report URL.
func reportPath(rawURL string) string {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return ""
	}
	return normalizeReportPath(parsed.Path)
}

// normalizeReportPath canonicalizes a page path while preserving trailing slash semantics.
func normalizeReportPath(p string) string {
	if p == "" {
		return "/"
	}
	if !strings.HasPrefix(p, "/") {
		p = "/" + p
	}
	cleaned := path.Clean(p)
	if strings.HasSuffix(p, "/") && !strings.HasSuffix(cleaned, "/") {
		cleaned += "/"
	}
	if cleaned == "." {
		return "/"
	}
	return cleaned
}

type commentInput struct {
	repo                   string
	title                  string
	artifactURL            string
	totalBroken            int
	changedDocsFileCount   int
	changedWithBrokenCount int
	brokenOnChangedPages   int
	rows                   []brokenRow
	maxRows                int
}

// buildComment renders the final Markdown body for the broken-links PR comment.
func buildComment(in commentInput) string {
	var b strings.Builder

	repo := strings.TrimSpace(in.repo)
	if repo == "" {
		repo = "unknown"
	}

	title := strings.TrimSpace(in.title)
	if title == "" {
		title = "Preview"
	}

	fmt.Fprintf(&b, "<!-- broken-links-report:%s -->\n", repo)
	fmt.Fprintf(&b, ":link: Broken links report (%s): %d broken links total.\n\n", title, in.totalBroken)

	if in.changedDocsFileCount == 0 {
		b.WriteString("No changed Markdown files matched configured docs source directories.\n\n")
	} else if in.changedWithBrokenCount == 0 {
		fmt.Fprintf(&b, "No broken links were detected on changed docs files in this PR (%d %s checked).\n\n", in.changedDocsFileCount, pluralize(in.changedDocsFileCount, "file", "files"))
	} else {
		fmt.Fprintf(&b, "Broken links on files changed in this PR: %d across %d %s.\n\n", in.brokenOnChangedPages, in.changedWithBrokenCount, pluralize(in.changedWithBrokenCount, "file", "files"))
		b.WriteString("| File | Page | Broken link | Error |\n")
		b.WriteString("| --- | --- | --- | --- |\n")
		rows := in.rows
		if in.maxRows > 0 && len(rows) > in.maxRows {
			rows = rows[:in.maxRows]
		}
		for _, row := range rows {
			fmt.Fprintf(&b, "| `%s` | `%s` | `%s` | `%s` |\n", escapePipes(row.File), escapePipes(row.PageURL), escapePipes(row.BrokenURL), escapePipes(row.Error))
		}
		if in.maxRows > 0 && len(in.rows) > in.maxRows {
			fmt.Fprintf(&b, "\nShowing first %d of %d broken-link rows.\n", in.maxRows, len(in.rows))
		}
		b.WriteString("\n")
	}

	if strings.TrimSpace(in.artifactURL) != "" {
		fmt.Fprintf(&b, "[Download the full links report artifact](%s).\n", strings.TrimSpace(in.artifactURL))
	} else {
		b.WriteString("The full links report is available as a workflow artifact.\n")
	}

	return b.String()
}

// escapePipes escapes table cell separators for Markdown table output.
func escapePipes(value string) string {
	return strings.ReplaceAll(value, "|", "\\|")
}

// pluralize returns the singular or plural form based on count.
func pluralize(count int, singular, plural string) string {
	if count == 1 {
		return singular
	}
	return plural
}
