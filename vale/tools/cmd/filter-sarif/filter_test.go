package main

import (
	"context"
	"testing"

	"github.com/google/go-github/v70/github"
	"github.com/grafana/writers-toolkit/vale/tools/cmd/filter-sarif/sarif"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFilterSarifByPR(t *testing.T) {
	t.Parallel()

	unfiltered, err := sarif.NewFromFile("testdata/102941.sarif")
	require.NoError(t, err)

	want, err := sarif.NewFromFile("testdata/102941-filtered.sarif")
	require.NoError(t, err)

	client := github.NewClient(nil)
	ctx := context.Background()

	filtered, hasRelevantErrors, err := filterSARIFByPR(ctx, client, unfiltered, "grafana", "grafana", 102941)
	assert.NoError(t, err) //nolint:testifylint
	assert.True(t, hasRelevantErrors)
	assert.Equal(t, want, filtered)
}
