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
	resolver, err := filepath.Abs("../resolve-preview-paths")
	if err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct {
		name, sources, legacy string
		want                  []string
	}{
		{"multiple projects", `[
			{"repo":"k6-docs","source_directory":"docs/sources","website_directory":"content/docs/example/next"},
			{"repo":"k6-docs","source_directory":"charts/docs","website_directory":"content/docs/helm-charts/example/latest"},
			{"repo":"other","source_directory":"missing/<LATEST>","website_directory":"content/docs/other/latest"},
			{"repo":"k6-docs","ref":"v1.0","source_directory":"missing","website_directory":"content/docs/example/v1.0"}
		]`, "", []string{"docs/sources", "charts/docs"}},
		// grafana/k6-docs#2038 used this source_directory for its deploy preview.
		{"k6 versioned source on caller branch", `[{"repo":"k6-docs","ref":"feature","source_directory":"docs/sources/k6/<LATEST>","website_directory":"content/docs/k6/latest"}]`, "", []string{"docs/sources/k6/v2.10.x"}},
		{"legacy", "", "content/docs/example/latest", []string{"docs/sources"}},
		{"no PR sources", `[{"repo":"other","source_directory":"docs","website_directory":"content/docs/other/latest"}]`, "", nil},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := t.TempDir()
			for _, version := range []string{"v2.9.x", "v2.10.1", "v2.10.x", "vendor"} {
				if err := os.MkdirAll(filepath.Join(root, "source-files/docs/sources/k6", version), 0o755); err != nil {
					t.Fatal(err)
				}
			}
			env := append(os.Environ(), "PATH="+root+":"+os.Getenv("PATH"), "GITHUB_WORKSPACE="+root,
				"DOCS_CHECKER_IMAGE=checker:test", "SHA=test-sha", "REPO=k6-docs", "BRANCH=feature",
				"BASE_SOURCES="+tc.sources, "WEBSITE_DIRECTORY="+tc.legacy, "SOURCE_DIRECTORY=docs/sources",
				"TARGET_REF_NAME=main", "GITHUB_OUTPUT=/dev/stdout")
			resolve := exec.Command("bash", resolver, "--schema")
			resolve.Env = env
			output, err := resolve.CombinedOutput()
			if err != nil {
				t.Fatalf("resolve source mappings: %v\n%s", err, output)
			}
			mappings, ok := strings.CutPrefix(strings.TrimSpace(string(output)), "sources=")
			if !ok {
				t.Fatalf("missing sources output in %q", output)
			}
			// Capture the exact Docker arguments without contacting a daemon.
			if err := os.WriteFile(filepath.Join(root, "docker"), []byte("#!/bin/sh\nprintf '%s\\n' \"$@\" > \"$GITHUB_WORKSPACE/args\"\n"), 0o755); err != nil {
				t.Fatal(err)
			}
			cmd := exec.Command("bash", helper)
			cmd.Env = append(env, "SOURCES="+mappings)
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
			if sources == nil || len(sources) != len(tc.want) {
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

func TestBuildPublishesMountedVersionsForLinkChecks(t *testing.T) {
	build, err := filepath.Abs("../build")
	if err != nil {
		t.Fatal(err)
	}
	root := t.TempDir()
	for _, dir := range []string{"src/k6-docs/docs/v1.9.x", "src/k6-docs/docs/v1.10.x", "src/k6-docs__release/docs/v1.9.x"} {
		if err := os.MkdirAll(filepath.Join(root, dir), 0o755); err != nil {
			t.Fatal(err)
		}
	}
	if err := os.WriteFile(filepath.Join(root, "docker"), []byte("#!/bin/sh\nprintf '%s\\n' \"$@\"\n"), 0o755); err != nil {
		t.Fatal(err)
	}
	sourceMappings := `[
		{"repo":"k6-docs","source_directory":"docs/<LATEST>","website_directory":"content/docs/k6/next"},
		{"repo":"k6-docs","ref":"release","source_directory":"docs/<LATEST>","website_directory":"content/docs/k6/latest"}
	]`
	output := filepath.Join(root, "build-output")
	cmd := exec.Command("bash", build)
	cmd.Dir = root
	cmd.Env = append(os.Environ(), "PATH="+root+":"+os.Getenv("PATH"), "TMPDIR="+root, "GITHUB_OUTPUT="+output,
		"WEBSITE_DIRECTORY=", "WHATS_NEW_DIR="+root+"/whats-new", "BRANCH=feature", "SOURCES="+sourceMappings)
	log, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("build: %v\n%s", err, log)
	}
	data, err := os.ReadFile(output)
	if err != nil {
		t.Fatal(err)
	}
	var mappings []struct {
		Source string `json:"source_directory"`
		Target string `json:"website_directory"`
	}
	if err := json.Unmarshal([]byte(strings.TrimPrefix(strings.TrimSpace(string(data)), "sources=")), &mappings); err != nil {
		t.Fatal(err)
	}
	if len(mappings) != 2 || mappings[0].Source != "docs/v1.10.x" || mappings[1].Source != "docs/v1.9.x" {
		t.Fatalf("unexpected build output: %s", data)
	}
	for _, m := range mappings {
		if !strings.Contains(string(log), "/"+m.Source+":/hugo/"+m.Target+":z") {
			t.Errorf("published source %q differs from Docker mount for %q", m.Source, m.Target)
		}
	}
}
