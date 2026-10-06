package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestConvert(t *testing.T) {
	manifest := "/docs/example/old/\t/docs/example/new/\n/legacy.html\t/docs/example/new/\n"
	var output bytes.Buffer

	count, err := convert(strings.NewReader(manifest), &output)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Fatalf("convert() count = %d, want 2", count)
	}
	want := "rewrite \"^/docs/example/old/?$\" \"/docs/example/new/\" permanent;\n" +
		"rewrite \"^/legacy\\\\.html$\" \"/docs/example/new/\" permanent;\n"
	if output.String() != want {
		t.Fatalf("convert() = %q, want %q", output.String(), want)
	}
}

func TestConvertRejectsMalformedManifest(t *testing.T) {
	var output bytes.Buffer
	if _, err := convert(strings.NewReader("/docs/example/old/"), &output); err == nil {
		t.Fatal("convert() error = nil, want malformed manifest error")
	}
}

func TestRunAllowsMissingManifest(t *testing.T) {
	outputPath := filepath.Join(t.TempDir(), "redirects.conf")
	if err := run(filepath.Join(t.TempDir(), "missing.txt"), outputPath); err != nil {
		t.Fatal(err)
	}
	output, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatal(err)
	}
	if len(output) != 0 {
		t.Fatalf("redirect output = %q, want empty", output)
	}
}
