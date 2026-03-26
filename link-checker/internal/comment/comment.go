package comment

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path"
	"sort"
	"strings"
)

type pageReport struct {
	URL        string       `json:"url"`
	SourcePath string       `json:"source_path,omitempty"`
	Links      []linkReport `json:"links"`
}

type linkReport struct {
	ID    string `json:"id,omitempty"`
	URL   string `json:"url"`
	Raw   string `json:"raw,omitempty"`
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
	Line      int
	Column    int
	LinkID    string
	PageURL   string
	BrokenURL string
	Error     string
}

type simpleBrokenRow struct {
	PageURL   string
	BrokenURL string
}

const (
	defaultSourceDirectory        = "docs/sources"
	generateDefaultRelativePrefix = "/docs/"
)

type runOptions struct {
	linksPath       string
	changedFiles    string
	baseRef         string
	outputPath      string
	title           string
	repo            string
	artifactURL     string
	sourceDirectory string
	relativePrefix  string
	sourcesJSON     string
	maxRows         int
}

// Run parses inputs, correlates changed files with broken links, and writes a PR comment body.
func Run(args []string) error {
	options, err := parseRunOptions(args)
	if err != nil {
		return err
	}

	reports, err := readReports(options.linksPath)
	if err != nil {
		return err
	}

	mappings := resolveMappings(options.sourcesJSON, options.sourceDirectory, options.relativePrefix)
	if strings.TrimSpace(options.sourcesJSON) == "" && strings.TrimSpace(options.relativePrefix) == "" {
		changedFiles, changedErr := readChangedFiles(options.changedFiles, options.baseRef)
		if changedErr != nil {
			return changedErr
		}
		if inferred := inferRelativePrefixFromReports(reports, changedFiles, options.sourceDirectory); inferred != "" {
			mappings = []mapping{{
				sourceDirectory: normalizeSourceDirectory(options.sourceDirectory),
				relativePrefix:  inferred,
			}}
		}
	}
	rows := collectRowsForComment(reports, mappings)
	totalBroken := len(rows)

	comment := buildComment(commentInput{
		repo:        options.repo,
		title:       options.title,
		artifactURL: options.artifactURL,
		totalBroken: totalBroken,
		rows:        rows,
		maxRows:     options.maxRows,
	})

	if err := os.WriteFile(options.outputPath, []byte(comment), 0o644); err != nil {
		return fmt.Errorf("write %s: %w", options.outputPath, err)
	}
	return nil
}

func parseRunOptions(args []string) (runOptions, error) {
	var options runOptions
	flags := flag.NewFlagSet("broken-links comment", flag.ContinueOnError)
	flags.StringVar(&options.linksPath, "links", "links.json", "Path to links JSON report")
	flags.StringVar(&options.changedFiles, "changed-files", "", "Path to changed files list (optional; defaults to git diff)")
	flags.StringVar(&options.baseRef, "base-ref", strings.TrimSpace(os.Getenv("BROKEN_LINKS_BASE_REF")), "Git base ref used to calculate changed files when -changed-files is not provided")
	flags.StringVar(&options.outputPath, "output", "broken-links-comment.md", "Path to write comment body")
	flags.StringVar(&options.title, "title", "", "PR title for comment heading")
	flags.StringVar(&options.repo, "repo", "", "Repository slug fragment for hidden comment marker")
	flags.StringVar(&options.artifactURL, "artifact-url", "", "Uploaded artifact URL")
	flags.StringVar(&options.sourceDirectory, "source-directory", strings.TrimSpace(os.Getenv("SOURCE_DIRECTORY")), "Legacy single source directory")
	flags.StringVar(&options.relativePrefix, "relative-prefix", strings.TrimSpace(os.Getenv("RELATIVE_PREFIX")), "Legacy single relative prefix")
	flags.StringVar(&options.sourcesJSON, "sources-json", strings.TrimSpace(os.Getenv("SOURCES")), "JSON array of source mappings")
	flags.IntVar(&options.maxRows, "max-rows", 150, "Maximum number of table rows to render")
	if err := flags.Parse(args); err != nil {
		return runOptions{}, err
	}
	return options, nil
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

// readChangedFiles returns changed file paths from file input or git diff.
func readChangedFiles(path, baseRef string) ([]string, error) {
	if strings.TrimSpace(path) != "" {
		return readChangedFilesFromFile(path)
	}
	return readChangedFilesFromGit(baseRef)
}

// readChangedFilesFromFile returns a sorted, de-duplicated list of changed file paths.
func readChangedFilesFromFile(path string) ([]string, error) {
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

// readChangedFilesFromGit calculates changed files from git diff commands.
func readChangedFilesFromGit(baseRef string) ([]string, error) {
	baseRef = strings.TrimSpace(baseRef)
	if baseRef == "" {
		baseRef = "origin/main"
	}

	if err := runGit("rev-parse", "--verify", baseRef); err != nil {
		return nil, fmt.Errorf("base ref %q does not exist: %w", baseRef, err)
	}

	mergeBaseOutput, err := runGitOutput("merge-base", "HEAD", baseRef)
	if err != nil {
		return nil, fmt.Errorf("compute merge base with %q: %w", baseRef, err)
	}
	mergeBase := strings.TrimSpace(mergeBaseOutput)
	if mergeBase == "" {
		return nil, fmt.Errorf("empty merge base for HEAD and %q", baseRef)
	}

	allFiles := map[string]struct{}{}
	commands := [][]string{
		{"diff", "--name-status", "--diff-filter=ACMR", mergeBase + "...HEAD"},
		{"diff", "--name-status", "--diff-filter=ACMR"},
		{"diff", "--name-status", "--cached", "--diff-filter=ACMR"},
	}
	for _, args := range commands {
		output, cmdErr := runGitOutput(args...)
		if cmdErr != nil {
			return nil, fmt.Errorf("git %s: %w", strings.Join(args, " "), cmdErr)
		}
		collectChangedFilesFromNameStatus(string(output), allFiles)
	}

	files := make([]string, 0, len(allFiles))
	for file := range allFiles {
		files = append(files, file)
	}
	sort.Strings(files)
	return files, nil
}

// collectChangedFilesFromNameStatus parses `git diff --name-status` output into file paths.
func collectChangedFilesFromNameStatus(output string, files map[string]struct{}) {
	scanner := bufio.NewScanner(strings.NewReader(output))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.Split(line, "\t")
		if len(parts) < 2 {
			continue
		}
		status := strings.TrimSpace(parts[0])
		if status == "" {
			continue
		}

		if strings.HasPrefix(status, "R") || strings.HasPrefix(status, "C") {
			if len(parts) >= 3 {
				if strings.TrimSpace(parts[1]) != "" {
					files[strings.TrimSpace(parts[1])] = struct{}{}
				}
				if strings.TrimSpace(parts[2]) != "" {
					files[strings.TrimSpace(parts[2])] = struct{}{}
				}
			}
			continue
		}

		if strings.TrimSpace(parts[1]) != "" {
			files[strings.TrimSpace(parts[1])] = struct{}{}
		}
	}
}

// runGit runs a git command and returns error if it fails.
func runGit(args ...string) error {
	cmd := exec.Command("git", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		message := strings.TrimSpace(string(output))
		if message != "" {
			return fmt.Errorf("%s", message)
		}
		return err
	}
	return nil
}

// runGitOutput runs a git command and returns stdout.
func runGitOutput(args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	output, err := cmd.Output()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok && len(exitErr.Stderr) > 0 {
			return "", fmt.Errorf("%s", strings.TrimSpace(string(exitErr.Stderr)))
		}
		return "", err
	}
	return string(output), nil
}

// filterReportsForComment removes links/pages that should not be included in PR comments.
func filterReportsForComment(reports []pageReport) []pageReport {
	return reports
}

// shouldExcludeFromComment returns true when a URL/path should be omitted from comment output.
func shouldExcludeFromComment(value string) bool {
	return strings.Contains(strings.ToLower(value), "unstyled")
}

// collectRowsForComment returns all broken-link rows from reports without changed-file filtering.
func collectRowsForComment(reports []pageReport, mappings []mapping) []brokenRow {
	rows := make([]brokenRow, 0)
	seen := map[string]struct{}{}
	locator := newSourceLinkLocator()
	for _, report := range reports {
		holderFile := sourceFileForReport(report, mappings)
		for _, link := range report.Links {
			line, column, _ := locator.find(holderFile, link)
			key := holderFile + "\x00" + report.URL + "\x00" + link.URL + "\x00" + link.Error + "\x00" + fmt.Sprintf("%d:%d:%s", line, column, link.ID)
			if _, ok := seen[key]; ok {
				continue
			}
			seen[key] = struct{}{}
			rows = append(rows, brokenRow{
				File:      holderFile,
				Line:      line,
				Column:    column,
				LinkID:    link.ID,
				PageURL:   report.URL,
				BrokenURL: link.URL,
				Error:     link.Error,
			})
		}
	}

	sort.Slice(rows, func(i, j int) bool {
		if rows[i].File != rows[j].File {
			return rows[i].File < rows[j].File
		}
		if rows[i].Line != rows[j].Line {
			return rows[i].Line < rows[j].Line
		}
		if rows[i].Column != rows[j].Column {
			return rows[i].Column < rows[j].Column
		}
		if rows[i].PageURL != rows[j].PageURL {
			return rows[i].PageURL < rows[j].PageURL
		}
		return rows[i].BrokenURL < rows[j].BrokenURL
	})

	return rows
}

type sourceFileContent struct {
	text       string
	lineStarts []int
}

type sourceLinkLocator struct {
	files map[string]*sourceFileContent
}

func newSourceLinkLocator() *sourceLinkLocator {
	return &sourceLinkLocator{
		files: map[string]*sourceFileContent{},
	}
}

func (locator *sourceLinkLocator) find(filePath string, link linkReport) (int, int, bool) {
	filePath = strings.TrimSpace(filePath)
	if filePath == "" || strings.Contains(filePath, "://") {
		return 0, 0, false
	}

	content, ok := locator.load(filePath)
	if !ok {
		return 0, 0, false
	}

	for _, needle := range linkSearchNeedles(link) {
		if needle == "" {
			continue
		}
		offset := strings.Index(content.text, needle)
		if offset < 0 {
			continue
		}
		line, column := offsetToLineColumn(content.lineStarts, offset)
		return line, column, true
	}

	return 0, 0, false
}

func (locator *sourceLinkLocator) load(filePath string) (*sourceFileContent, bool) {
	if content, ok := locator.files[filePath]; ok {
		return content, content != nil
	}

	bytes, err := os.ReadFile(filePath)
	if err != nil {
		locator.files[filePath] = nil
		return nil, false
	}

	text := string(bytes)
	lineStarts := []int{0}
	for index := 0; index < len(text); index++ {
		if text[index] == '\n' {
			lineStarts = append(lineStarts, index+1)
		}
	}

	content := &sourceFileContent{
		text:       text,
		lineStarts: lineStarts,
	}
	locator.files[filePath] = content
	return content, true
}

func linkSearchNeedles(link linkReport) []string {
	needles := make([]string, 0, 12)
	seen := map[string]struct{}{}
	appendNeedle := func(value string) {
		value = strings.TrimSpace(value)
		if value == "" {
			return
		}
		if _, ok := seen[value]; ok {
			return
		}
		seen[value] = struct{}{}
		needles = append(needles, value)
	}

	appendNeedle(link.Raw)
	appendNeedle(link.URL)

	parsed, err := url.Parse(link.URL)
	if err != nil {
		return needles
	}

	pathValue := strings.TrimSpace(parsed.Path)
	latestPath := strings.ReplaceAll(pathValue, "/latest/", "/<GRAFANA_VERSION>/")
	if strings.HasPrefix(pathValue, "/docs/") {
		appendNeedle("https://grafana.com" + pathValue)
		appendNeedle("http://grafana.com" + pathValue)
		if latestPath != pathValue {
			appendNeedle("https://grafana.com" + latestPath)
			appendNeedle("http://grafana.com" + latestPath)
		}
	}
	appendNeedle(pathValue)
	appendNeedle(strings.TrimSuffix(pathValue, "/"))
	if decodedPath, decodeErr := url.PathUnescape(pathValue); decodeErr == nil {
		appendNeedle(decodedPath)
		appendNeedle(strings.TrimSuffix(decodedPath, "/"))
	}
	appendNeedle(latestPath)
	appendNeedle(strings.TrimSuffix(latestPath, "/"))

	return needles
}

func offsetToLineColumn(lineStarts []int, offset int) (int, int) {
	if offset < 0 {
		return 0, 0
	}
	index := sort.Search(len(lineStarts), func(i int) bool {
		return lineStarts[i] > offset
	}) - 1
	if index < 0 {
		index = 0
	}
	line := index + 1
	column := offset - lineStarts[index] + 1
	return line, column
}

// collectBrokenRows returns broken-link rows relevant to changed files and moved targets.
func collectBrokenRows(reports []pageReport, fileToCandidates map[string][]string, mappings []mapping) ([]brokenRow, []string) {
	reportsByPath := make(map[string]pageReport, len(reports))
	for _, report := range reports {
		reportPath := reportPath(report.URL)
		if reportPath == "" {
			continue
		}
		reportsByPath[reportPath] = report
	}

	candidatePathToFiles := make(map[string][]string)
	for file, candidates := range fileToCandidates {
		for _, candidate := range candidates {
			candidatePathToFiles[candidate] = appendUniqueString(candidatePathToFiles[candidate], file)
		}
	}

	rows := make([]brokenRow, 0)
	rowSeen := map[string]struct{}{}
	changedFilesWithBrokenSet := map[string]struct{}{}
	addRow := func(changedFile, holderFile, pageURL, brokenURL, errText string) {
		if changedFile != "" {
			changedFilesWithBrokenSet[changedFile] = struct{}{}
		}
		key := holderFile + "\x00" + pageURL + "\x00" + brokenURL + "\x00" + errText
		if _, ok := rowSeen[key]; ok {
			return
		}
		rowSeen[key] = struct{}{}
		rows = append(rows, brokenRow{
			File:      holderFile,
			PageURL:   pageURL,
			BrokenURL: brokenURL,
			Error:     errText,
		})
	}

	for file, candidates := range fileToCandidates {
		for _, candidate := range candidates {
			report, ok := reportsByPath[candidate]
			if !ok {
				continue
			}
			holderFile := sourceFileForReport(report, mappings)
			for _, link := range report.Links {
				addRow(file, holderFile, report.URL, link.URL, link.Error)
			}
		}
	}

	for _, report := range reports {
		for _, link := range report.Links {
			brokenPath := reportPath(link.URL)
			if brokenPath == "" {
				continue
			}
			files, ok := candidatePathToFiles[brokenPath]
			if !ok {
				continue
			}
			holderFile := sourceFileForReport(report, mappings)
			for _, file := range files {
				addRow(file, holderFile, report.URL, link.URL, link.Error)
			}
		}
	}

	changedFilesWithBroken := make([]string, 0, len(changedFilesWithBrokenSet))
	for file := range changedFilesWithBrokenSet {
		changedFilesWithBroken = append(changedFilesWithBroken, file)
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

	return rows, changedFilesWithBroken
}

// sourceFileForReport resolves the source file path for a report.
func sourceFileForReport(report pageReport, mappings []mapping) string {
	sourcePath := strings.TrimSpace(report.SourcePath)
	if sourcePath != "" {
		sourcePath = strings.ReplaceAll(sourcePath, "\\", "/")
		sourcePath = strings.TrimPrefix(sourcePath, "/")
		if mappedPath, ok := remapSourcePath(sourcePath, mappings); ok {
			return mappedPath
		}
		return sourcePath
	}
	return sourceFileForPageURL(report.URL, mappings)
}

// remapSourcePath replaces a relative-prefix path root with the mapped source directory.
func remapSourcePath(sourcePath string, mappings []mapping) (string, bool) {
	sortedMappings := make([]mapping, 0, len(mappings))
	sortedMappings = append(sortedMappings, mappings...)
	sort.Slice(sortedMappings, func(i, j int) bool {
		return len(sortedMappings[i].relativePrefix) > len(sortedMappings[j].relativePrefix)
	})

	for _, m := range sortedMappings {
		relativeRoot := strings.Trim(normalizeRelativePrefixForComment(m.relativePrefix), "/")
		sourceRoot := strings.Trim(strings.TrimSpace(m.sourceDirectory), "/")
		if relativeRoot == "" || sourceRoot == "" {
			continue
		}

		if sourcePath == relativeRoot {
			return sourceRoot, true
		}
		prefix := relativeRoot + "/"
		if strings.HasPrefix(sourcePath, prefix) {
			return sourceRoot + "/" + strings.TrimPrefix(sourcePath, prefix), true
		}
	}

	return "", false
}

// sourceFileForPageURL maps a rendered page URL back to its most likely source markdown path.
func sourceFileForPageURL(pageURL string, mappings []mapping) string {
	reportPathValue := reportPath(pageURL)
	if reportPathValue == "" {
		return pageURL
	}

	sortedMappings := make([]mapping, 0, len(mappings))
	sortedMappings = append(sortedMappings, mappings...)
	sort.Slice(sortedMappings, func(i, j int) bool {
		return len(sortedMappings[i].relativePrefix) > len(sortedMappings[j].relativePrefix)
	})

	for _, m := range sortedMappings {
		prefix := normalizeReportPath(m.relativePrefix)
		if !strings.HasSuffix(prefix, "/") {
			prefix += "/"
		}
		if reportPathValue == strings.TrimSuffix(prefix, "/") || reportPathValue == prefix {
			return m.sourceDirectory + "/_index.md"
		}
		if !strings.HasPrefix(reportPathValue, prefix) {
			continue
		}

		rel := strings.TrimPrefix(reportPathValue, prefix)
		if rel == "" {
			return m.sourceDirectory + "/_index.md"
		}
		if strings.HasSuffix(rel, "/") {
			rel = strings.TrimSuffix(rel, "/")
			if rel == "" {
				return m.sourceDirectory + "/_index.md"
			}
			return m.sourceDirectory + "/" + rel + "/index.md"
		}
		if strings.HasSuffix(rel, ".html") {
			base := strings.TrimSuffix(rel, ".html")
			if base == "index" {
				return m.sourceDirectory + "/_index.md"
			}
			return m.sourceDirectory + "/" + base + ".md"
		}
		return m.sourceDirectory + "/" + rel + ".md"
	}

	return pageURL
}

// collectAllBrokenRows flattens all broken links from report data for fallback rendering.
func collectAllBrokenRows(reports []pageReport) []simpleBrokenRow {
	rows := make([]simpleBrokenRow, 0)
	for _, report := range reports {
		for _, link := range report.Links {
			rows = append(rows, simpleBrokenRow{
				PageURL:   report.URL,
				BrokenURL: link.URL,
			})
		}
	}
	sort.Slice(rows, func(i, j int) bool {
		if rows[i].PageURL != rows[j].PageURL {
			return rows[i].PageURL < rows[j].PageURL
		}
		return rows[i].BrokenURL < rows[j].BrokenURL
	})
	return rows
}

// appendUniqueString appends value if it is not already present.
func appendUniqueString(values []string, value string) []string {
	for _, existing := range values {
		if existing == value {
			return values
		}
	}
	return append(values, value)
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

// inferRelativePrefixFromReports infers the relative prefix from links.json for local runs.
func inferRelativePrefixFromReports(reports []pageReport, changedFiles []string, sourceDirectory string) string {
	sd := normalizeSourceDirectory(sourceDirectory)
	defaultMapping := []mapping{{
		sourceDirectory: sd,
		relativePrefix:  generateDefaultRelativePrefix,
	}}

	tails := make([]string, 0)
	seenTail := map[string]struct{}{}
	for _, file := range changedFiles {
		candidates := candidatePagePaths(file, defaultMapping)
		for _, candidate := range candidates {
			if !strings.HasPrefix(candidate, generateDefaultRelativePrefix) {
				continue
			}
			tail := strings.TrimPrefix(candidate, generateDefaultRelativePrefix)
			if tail == "" {
				continue
			}
			if _, ok := seenTail[tail]; ok {
				continue
			}
			seenTail[tail] = struct{}{}
			tails = append(tails, tail)
		}
	}
	if len(tails) == 0 {
		return ""
	}

	sort.Slice(tails, func(i, j int) bool {
		if len(tails[i]) != len(tails[j]) {
			return len(tails[i]) > len(tails[j])
		}
		return tails[i] < tails[j]
	})

	prefixHits := map[string]int{}
	matchPath := func(candidatePath string) {
		if candidatePath == "" {
			return
		}
		for _, tail := range tails {
			if !strings.HasSuffix(candidatePath, tail) {
				continue
			}
			prefix := strings.TrimSuffix(candidatePath, tail)
			if prefix == "" {
				continue
			}
			prefix = normalizeRelativePrefixForComment(prefix)
			prefixHits[prefix]++
			break
		}
	}
	for _, report := range reports {
		matchPath(reportPath(report.URL))
		for _, link := range report.Links {
			matchPath(reportPath(link.URL))
		}
	}

	bestPrefix := ""
	bestHits := 0
	for prefix, hits := range prefixHits {
		if hits > bestHits || (hits == bestHits && (bestPrefix == "" || len(prefix) > len(bestPrefix))) {
			bestPrefix = prefix
			bestHits = hits
		}
	}
	if bestHits == 0 {
		return ""
	}
	return bestPrefix
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
		if strings.HasSuffix(rel, "/_index.md") {
			suffix := strings.TrimSuffix(rel, "/_index.md")
			candidate := normalizeReportPath(path.Join(m.relativePrefix, suffix) + "/")
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
	fallbackRows           []simpleBrokenRow
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
	fmt.Fprintf(&b, ":link: Broken links report (%s): %d broken %s total.\n\n", title, in.totalBroken, pluralize(in.totalBroken, "link", "links"))

	if in.totalBroken == 0 {
		b.WriteString("No broken links found.\n\n")
	} else {
		b.WriteString("Broken links found in this build:\n\n")
		rows := in.rows
		if in.maxRows > 0 && len(rows) > in.maxRows {
			rows = rows[:in.maxRows]
		}
		tableRows := make([][]string, 0, len(rows))
		for _, row := range rows {
			fileValue := displayFilePath(row.File)
			if row.Line > 0 && row.Column > 0 {
				fileValue = fmt.Sprintf("%s:%d:%d", fileValue, row.Line, row.Column)
			}
			tableRows = append(tableRows, []string{
				fmt.Sprintf("`%s`", escapePipes(fileValue)),
				fmt.Sprintf("`%s`", escapePipes(row.BrokenURL)),
				fmt.Sprintf("`%s`", escapePipes(row.Error)),
			})
		}
		b.WriteString(renderMarkdownTable([]string{"File", "Broken link", "Error"}, tableRows))
		if in.maxRows > 0 && len(in.rows) > in.maxRows {
			fmt.Fprintf(&b, "\nShowing first %d of %d broken-link rows.\n", in.maxRows, len(in.rows))
		}
		b.WriteString("\n")
	}

	if strings.TrimSpace(in.artifactURL) != "" {
		fmt.Fprintf(&b, "[Download the full links report artifact](%s).\n", strings.TrimSpace(in.artifactURL))
	}

	return b.String()
}

func displayFilePath(filePath string) string {
	filePath = strings.TrimSpace(filePath)
	filePath = strings.TrimPrefix(filePath, "source-files/")
	return filePath
}

// renderMarkdownTable renders a markdown table with padded, equal-width columns.
func renderMarkdownTable(headers []string, rows [][]string) string {
	if len(headers) == 0 {
		return ""
	}

	widths := make([]int, len(headers))
	for i, header := range headers {
		widths[i] = len(header)
	}
	for _, row := range rows {
		for i := 0; i < len(headers) && i < len(row); i++ {
			if len(row[i]) > widths[i] {
				widths[i] = len(row[i])
			}
		}
	}

	var b strings.Builder
	writeRow := func(cells []string) {
		b.WriteString("|")
		for i := range headers {
			value := ""
			if i < len(cells) {
				value = cells[i]
			}
			fmt.Fprintf(&b, " %-*s |", widths[i], value)
		}
		b.WriteString("\n")
	}

	writeRow(headers)
	b.WriteString("|")
	for _, width := range widths {
		b.WriteString(" ")
		b.WriteString(strings.Repeat("-", width))
		b.WriteString(" |")
	}
	b.WriteString("\n")
	for _, row := range rows {
		writeRow(row)
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
