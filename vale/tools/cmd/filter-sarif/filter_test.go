package main

import (
	"context"
	"os"
	"testing"

	"github.com/google/go-github/v72/github"
	"github.com/grafana/writers-toolkit/vale/tools/cmd/filter-sarif/sarif"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFilterSARIFByPatch(t *testing.T) {
	t.Parallel()

	unfiltered, err := sarif.NewFromFile("testdata/102183.sarif")
	require.NoError(t, err)

	patch, err := os.ReadFile("testdata/102183.patch")
	require.NoError(t, err)

	_, hasRelevantErrors, err := filterSARIFByPatch(unfiltered, patch)
	require.NoError(t, err) //nolint:testifylint
	assert.True(t, hasRelevantErrors)
}

func TestFilterSARIFByPR(t *testing.T) {
	t.Parallel()

	unfiltered, err := sarif.NewFromFile("testdata/102941.sarif")
	require.NoError(t, err)

	want, err := sarif.NewFromFile("testdata/102941-filtered.sarif")
	require.NoError(t, err)

	client := github.NewClient(nil)
	ctx := context.Background()

	filtered, hasRelevantErrors, err := filterSARIFByPR(ctx, client, unfiltered, "grafana", "grafana", 102941)
	require.NoError(t, err)
	assert.True(t, hasRelevantErrors)
	assert.Equal(t, want, filtered)
}

func TestInhibitResults(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string

		results []sarif.Result
		want    []sarif.Result
	}{
		{
			"Grafana.WordList inhibits Grafana.Spelling",
			[]sarif.Result{
				{
					RuleID: "Grafana.Spelling",
					Level:  "error",
					Locations: []sarif.Location{{
						PhysicalLocation: sarif.PhysicalLocation{
							ArtifactLocation: sarif.ArtifactLocation{
								URI: "file:///path/to/file",
							},
							Region: sarif.Region{StartLine: 1, StartColumn: 1},
						},
					},
					},
				},
				{
					RuleID: "Grafana.WordList",
					Level:  "warning",
					Locations: []sarif.Location{{
						PhysicalLocation: sarif.PhysicalLocation{
							ArtifactLocation: sarif.ArtifactLocation{
								URI: "file:///path/to/file",
							},
							Region: sarif.Region{StartLine: 1, StartColumn: 1},
						},
					},
					},
				},
			},
			[]sarif.Result{{
				RuleID: "Grafana.WordList",
				Level:  "warning",
				Locations: []sarif.Location{{
					PhysicalLocation: sarif.PhysicalLocation{
						ArtifactLocation: sarif.ArtifactLocation{
							URI: "file:///path/to/file",
						},
						Region: sarif.Region{StartLine: 1, StartColumn: 1},
					},
				}},
			}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := inhibitResults(tt.results)
			assert.Equal(t, tt.want, got)
		})
	}
}
