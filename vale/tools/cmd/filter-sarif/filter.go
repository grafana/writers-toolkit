package main

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/google/go-github/v72/github"
	"github.com/grafana/writers-toolkit/vale/tools/cmd/filter-sarif/sarif"
	"github.com/sourcegraph/go-diff/diff"
)

var (
	errDiffFetch  = errors.New("couldn't get pull request diff")
	errPatchParse = errors.New("couldn't parse patch data")
)

// filterSARIFByPatch filters the results of a SARIF file so that it only includes those that are present in the patch data.
// The boolean return value indicates whether any results were present in the patch.
func filterSARIFByPatch(sarifFile sarif.File, patchData []byte) (sarif.File, bool, error) {
	diffs, err := diff.ParseMultiFileDiff(patchData)
	if err != nil {
		return sarifFile, false, fmt.Errorf("%w: %w", errPatchParse, err)
	}

	modifiedLines := make(map[string]map[int]struct{})
	for _, diff := range diffs {
		path := strings.TrimPrefix(diff.NewName, "b/")
		modifiedLines[path] = make(map[int]struct{})

		for _, hunk := range diff.Hunks {
			lineNum := int(hunk.NewStartLine)
			lines := strings.Split(string(hunk.Body), "\n")

			for _, line := range lines {
				if len(line) == 0 {
					continue
				}

				switch line[0] {
				case '+':
					modifiedLines[path][lineNum] = struct{}{}
					lineNum++
				case ' ':
					lineNum++
				}
			}
		}
	}

	var (
		filtered = sarif.File{
			Version: sarifFile.Version,
			Schema:  sarifFile.Schema,
			Runs:    []sarif.Run{},
		}
		hasResults bool
	)

	for _, run := range sarifFile.Runs {
		results := []sarif.Result{}

		for _, result := range run.Results {
			location := result.Locations[0]
			path := location.PhysicalLocation.ArtifactLocation.URI
			line := location.PhysicalLocation.Region.StartLine

			if lines, ok := modifiedLines[path]; ok {
				if _, ok := lines[line]; ok {
					results = append(results, result)
					hasResults = true
				}
			}
		}

		filtered.Runs = append(filtered.Runs, sarif.Run{
			Tool:    run.Tool,
			Results: inhibitResults(results),
		})
	}

	return filtered, hasResults, nil
}

// filterSARIFByPR filters the results of a SARIF file so that it only includes those that are present in the pull request patch.
// The boolean return value indicates whether any results were present in the pull request patch.
func filterSARIFByPR(ctx context.Context, client *github.Client, sarifFile sarif.File, owner, repo string, number int) (sarif.File, bool, error) {
	pr, _, err := client.PullRequests.GetRaw(ctx, owner, repo, number, github.RawOptions{Type: github.Diff})
	if err != nil {
		return sarifFile, false, fmt.Errorf("%w: %w", errDiffFetch, err)
	}

	return filterSARIFByPatch(sarifFile, []byte(pr))
}

// locationKey returns a unique key for a SARIF location.
func locationKey(location sarif.Location) string {
	return fmt.Sprintf("%s:%d:%d",
		location.PhysicalLocation.ArtifactLocation.URI,
		location.PhysicalLocation.Region.StartLine,
		location.PhysicalLocation.Region.StartColumn,
	)
}

// inhibitResults filters the results of a SARIF file using rule precedence.
// It keeps only the highest precedence rule for a line and column.
func inhibitResults(results []sarif.Result) []sarif.Result {
	filtered := make(map[string]sarif.Result)

	precedence := map[string]int{
		"Grafana.GrafanaCom":         0,
		"Grafana.ProductPossessives": 1,
		"Grafana.WordList":           2,
		"Grafana.Spelling":           3,
	}

	for _, result := range results {
		for _, location := range result.Locations {
			if _, ok := filtered[locationKey(location)]; !ok || precedence[result.RuleID] < precedence[filtered[locationKey(location)].RuleID] {
				filtered[locationKey(location)] = result
			}
		}
	}

	out := make([]sarif.Result, 0, len(filtered))
	for _, result := range filtered {
		out = append(out, result)
	}

	return out
}
