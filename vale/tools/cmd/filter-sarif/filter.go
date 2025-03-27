package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"

	"github.com/google/go-github/v70/github"
	"github.com/grafana/writers-toolkit/vale/tools/cmd/filter-sarif/sarif"
	"rsc.io/tmp/patch"
)

var (
	errDiffFetch  = errors.New("couldn't get pull request diff")
	errPatchParse = errors.New("couldn't parse patch data")
)

// filterSARIFByPatch filters the results of a SARIF file so that it only includes those that are present in the patch data.
// The boolean return value indicates whether any results were present in the patch.
func filterSARIFByPatch(sarifFile sarif.File, patchData []byte) (sarif.File, bool, error) {
	set, err := patch.Parse(patchData)
	if err != nil {
		return sarifFile, false, fmt.Errorf("%w: %w", errPatchParse, err)
	}

	modifiedLines := getModifiedLines(set)

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

// getModifiedLines returns the map of paths to its set of modified line numbers for all paths in the patch set.
func getModifiedLines(set *patch.Set) map[string]map[int]struct{} {
	modifiedLines := make(map[string]map[int]struct{})
	for _, f := range set.File {
		modifiedLines[f.Dst] = make(map[int]struct{})

		if textDiff, ok := f.Diff.(patch.TextDiff); ok {
			for _, chunk := range textDiff {
				if len(chunk.New) > 0 {
					newLines := bytes.Split(chunk.New, []byte("\n"))

					for i := range newLines {
						if len(chunk.Old) == 0 {
							modifiedLines[f.Dst][chunk.Line+i] = struct{}{}
						}

						if len(chunk.Old) > 0 {
							var preexisting bool

							oldLines := bytes.Split(chunk.Old, []byte("\n"))

							for j := range oldLines {
								if bytes.Equal(oldLines[j], newLines[i]) {
									preexisting = true

									break
								}
							}

							if !preexisting {
								modifiedLines[f.Dst][chunk.Line+i] = struct{}{}
							}
						}
					}
				}
			}
		}
	}

	return modifiedLines
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
