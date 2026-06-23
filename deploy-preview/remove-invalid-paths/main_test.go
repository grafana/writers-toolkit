package main

import (
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"testing"
)

func TestHasInvalidArtifactPathCharacters(t *testing.T) {
	t.Parallel()

	for _, invalidCharacter := range []string{"\"", ":", "<", ">", "|", "*", "?", "\r", "\n"} {
		invalidCharacter := invalidCharacter

		t.Run("detects "+invalidCharacter, func(t *testing.T) {
			t.Parallel()

			path := "docs/grafana/tempo" + invalidCharacter + "/index.html"
			if !hasInvalidArtifactPathCharacters(path) {
				t.Fatalf("expected invalid character %q to be detected in %q", invalidCharacter, path)
			}
		})
	}

	for _, tc := range []struct {
		name string
		path string
		want bool
	}{
		{name: "valid path", path: "docs/grafana/v11/datasources/tempo/index.html", want: false},
		{name: "invalid less than", path: "docs/grafana/<GRAFANA_VERSION>/index.html", want: true},
	} {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := hasInvalidArtifactPathCharacters(tc.path)
			if got != tc.want {
				t.Fatalf("hasInvalidArtifactPathCharacters(%q) = %v, want %v", tc.path, got, tc.want)
			}
		})
	}
}

func TestRemoveInvalidFiles(t *testing.T) {
	t.Parallel()

	root := t.TempDir()

	validPath := filepath.Join(root, "docs/grafana/v11/datasources/tempo/configure-tempo-data-source/index.html")
	invalidDirPath := filepath.Join(root, "docs/grafana/<GRAFANA_VERSION>/datasources/tempo/configure-tempo-data-source/#additional-settings/index.html")
	invalidFilePath := filepath.Join(root, "docs/grafana/v11/datasources/tempo/configure-tempo-data-source/question?/index.html")

	for _, path := range []string{validPath, invalidDirPath, invalidFilePath} {
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			t.Fatalf("mkdir %s: %v", filepath.Dir(path), err)
		}

		if err := os.WriteFile(path, []byte("ok"), 0o644); err != nil {
			t.Fatalf("write file %s: %v", path, err)
		}
	}

	removed, err := removeInvalidFiles(root)
	if err != nil {
		t.Fatalf("removeInvalidFiles() error = %v", err)
	}

	sort.Strings(removed)
	want := []string{
		"docs/grafana/<GRAFANA_VERSION>/datasources/tempo/configure-tempo-data-source/#additional-settings/index.html",
		"docs/grafana/v11/datasources/tempo/configure-tempo-data-source/question?/index.html",
	}

	if !reflect.DeepEqual(removed, want) {
		t.Fatalf("removed = %v, want %v", removed, want)
	}

	if _, err := os.Stat(validPath); err != nil {
		t.Fatalf("valid file should remain: %v", err)
	}

	if _, err := os.Stat(invalidDirPath); !os.IsNotExist(err) {
		t.Fatalf("invalid file should be removed, stat err = %v", err)
	}

	if _, err := os.Stat(invalidFilePath); !os.IsNotExist(err) {
		t.Fatalf("invalid file should be removed, stat err = %v", err)
	}

	if _, err := os.Stat(filepath.Dir(invalidDirPath)); !os.IsNotExist(err) {
		t.Fatalf("empty invalid directory should be removed, stat err = %v", err)
	}
}
