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

var errDiffFetch = errors.New("couldn't get pull request diff")

// filterSARIFByPatch filters the results of a SARIF file so that it only includes those that are present in the patch data.
// The boolean return value indicates whether any results were present in the patch.
func filterSARIFByPatch(sarifFile sarif.File, patchData []byte) (sarif.File, bool) {
	set, err := patch.Parse(patchData)
	if err != nil {
		return sarifFile, false
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
			Results: results,
		})
	}

	return filtered, hasResults
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

	filtered, hasResults := filterSARIFByPatch(sarifFile, []byte(pr))

	return filtered, hasResults, nil
}
