package main

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestSchemaCheckUsesPRSourceMappings(t *testing.T) {
	helper, err := filepath.Abs("../schema-check")
	if err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct {
		name, sources, legacy string
		want                  []string
	}{
		{"multiple projects", `[
			{"repo":"example","source_directory":"docs/sources","website_directory":"content/docs/example/next"},
			{"repo":"example","source_directory":"charts/docs","website_directory":"content/docs/helm-charts/example/latest"},
			{"repo":"other","source_directory":"missing","website_directory":"content/docs/other/latest"},
			{"repo":"example","ref":"v1.0","source_directory":"missing","website_directory":"content/docs/example/v1.0"}
		]`, "", []string{"docs/sources", "charts/docs"}},
		{"caller branch", `[{"repo":"example","ref":"feature","source_directory":"docs/<LATEST>/sources","website_directory":"content/docs/example/next"}]`, "", []string{"docs/<LATEST>/sources"}},
		{"legacy", "", "content/docs/example/latest", []string{"docs/sources"}},
		{"no PR sources", `[{"repo":"other","source_directory":"docs","website_directory":"content/docs/other/latest"}]`, "", nil},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := t.TempDir()
			if err := os.MkdirAll(filepath.Join(root, "lint-reports"), 0o755); err != nil {
				t.Fatal(err)
			}
			if err := os.WriteFile(filepath.Join(root, "lint-reports/lint.txt"), []byte("No findings.\n"), 0o644); err != nil {
				t.Fatal(err)
			}
			// Capture the exact Docker arguments without contacting a daemon.
			if err := os.WriteFile(filepath.Join(root, "docker"), []byte("#!/bin/sh\nprintf '%s\\n' \"$@\" > \"$GITHUB_WORKSPACE/args\"\n"), 0o755); err != nil {
				t.Fatal(err)
			}
			cmd := exec.Command("bash", helper)
			cmd.Env = append(os.Environ(), "PATH="+root+":"+os.Getenv("PATH"), "GITHUB_WORKSPACE="+root,
				"DOCS_CHECKER_IMAGE=checker:test", "SHA=test-sha", "REPO=example", "BRANCH=feature",
				"SOURCES="+tc.sources, "WEBSITE_DIRECTORY="+tc.legacy, "SOURCE_DIRECTORY=docs/sources")
			if out, err := cmd.CombinedOutput(); err != nil {
				t.Fatalf("helper: %v\n%s", err, out)
			}
			data, err := os.ReadFile(filepath.Join(root, "args"))
			if err != nil {
				t.Fatal(err)
			}
			args := strings.Split(strings.TrimSpace(string(data)), "\n")
			flag := func(name string) string {
				t.Helper()
				for i := range len(args) - 1 {
					if args[i] == name {
						return args[i+1]
					}
				}
				t.Fatalf("missing %s in %v", name, args)
				return ""
			}
			if flag("--workdir") != "/work/source-files" || flag("-mode") != "changes" || flag("-changed-files-json") != "/work/changed-files.json" || flag("-output-dir") != "/work/lint-reports" || flag("-sha") != "test-sha" || flag("checker:test") != "schema-check" {
				t.Fatalf("wrong container paths/command: %v", args)
			}
			var sources []struct {
				Source string `json:"source_directory"`
				Target string `json:"website_directory"`
			}
			if err := json.Unmarshal([]byte(flag("-sources-json")), &sources); err != nil {
				t.Fatal(err)
			}
			if len(sources) != len(tc.want) {
				t.Fatalf("sources = %+v, want %v", sources, tc.want)
			}
			for i, source := range sources {
				if source.Source != tc.want[i] || !strings.HasPrefix(source.Target, "content/docs/") {
					t.Fatalf("source = %+v, want %s", source, tc.want[i])
				}
			}
		})
	}
}
