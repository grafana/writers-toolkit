package main

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/grafana/writers-toolkit/tools/exit"
	"github.com/grafana/writers-toolkit/tools/rwfilefs"
	"gopkg.in/yaml.v3"
)

const (
	command = "generate-documentation"
	tmpl    = `---
date: "2024-06-25"
description: A description of every Grafana Labs prose linting rule.
menuTitle: Rules
review_date: "{{ .ReviewDate }}"
title: {{ .Title }}
---

# {{ .Title }}

<!-- These are our style rules. -->
<!-- vale Grafana.We = NO -->

Vale codifies our style guide into a series of rules that can be checked against your prose.
The following is a list of all the rules that we've defined.

<!-- This page breaks a number of rules in demonstrating them. -->
{{ range .Rules -}}
<!-- vale {{ .Name }} = NO -->
{{ end -}}

## Errors

The following rules are considered errors and must be fixed.

{{ range .Rules -}}
{{ if or (eq .Level "error") -}}

### {{ .Name }}

Extends: {{ .Extends }}

{{ escapeShortcodes .Message }}

{{ with .Link }}[More information ->]({{ . }}){{ end }}

{{ end -}}
{{ end -}}

## Warnings

The following rules are warnings and may need to be fixed or otherwise require consideration.

{{ range .Rules -}}
{{ if or (eq .Level "warning") (eq .Level "") -}}
### {{ .Name }}

Extends: {{ .Extends }}

{{ escapeShortcodes .Message }}

{{ with .Link }}[More information ->]({{ . }}){{ end }}

{{ end -}}
{{ end -}}

## Suggestions

The following rules are suggestions to consider a certain point of style.

{{ range .Rules -}}
{{ if or (eq .Level "suggestion") -}}
### {{ .Name }}

Extends: {{ .Extends }}

{{ escapeShortcodes .Message }}

{{ with .Link }}[More information ->]({{ . }}){{ end }}

{{ end -}}
{{ end -}}
`
)

var errProcessingRule = errors.New("error processing rule")

type PageData struct {
	ReviewDate string `yaml:"review_date"`
	Rules      []Rule `yaml:"rules"`
	Title      string `yaml:"title"`
}

type Rule struct {
	Extends string `yaml:"extends"`
	Level   string `yaml:"level"`
	Link    string `yaml:"link"`
	Message string `yaml:"message"`
	Name    string `yaml:"name"`
}

func usage(w io.Writer, fs *flag.FlagSet) {
	fmt.Fprintf(w, "Usage of %s:\n", command)
	fs.PrintDefaults()

	fmt.Fprintln(w, "  <REPOSITORY PATH>")
	fmt.Fprintln(w, "    	Path to the repository root.")

	fmt.Fprintln(w, "  <SOURCE PATH>")
	fmt.Fprintln(w, "    	Path to the Vale rules directory, relative to <REPOSITORY PATH>.")

	fmt.Fprintln(w, "  <OUTPUT PATH>")
	fmt.Fprintln(w, "    	Path to the output file, relative to <REPOSITORY PATH>.")
}

func generate(fsys rwfilefs.RWFileFS, srcDirPath, dstPath string) error {
	var (
		rules     []Rule
		rulesErrs error
	)

	if err := fs.WalkDir(fsys, srcDirPath, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("couldn't walk Vale rules directory: %w", err)
		}

		if info.IsDir() {
			return nil
		}

		data, err := fsys.ReadFile(path)
		if err != nil {
			rulesErrs = errors.Join(rulesErrs, fmt.Errorf("%v: %w", errProcessingRule, err))
		}

		rule := Rule{
			Extends: "",
			Level:   "",
			Link:    "",
			Message: "",
			Name:    fmt.Sprintf("Grafana.%s", strings.Title(strings.TrimSuffix(filepath.Base(path), ".yml"))),
		}

		if err := yaml.Unmarshal(data, &rule); err != nil {
			rulesErrs = errors.Join(rulesErrs, fmt.Errorf("%v: %w", errProcessingRule, err))
		}

		rules = append(rules, rule)

		return nil
	}); err != nil {
		return fmt.Errorf("couldn't walk Vale rules directory: %w", err)
	}

	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)

	pageData := PageData{
		ReviewDate: time.Now().Format("2006-01-02"),
		Title:      "Vale rules",
		Rules:      rules,
	}

	funcMap := template.FuncMap{
		"escapeShortcodes": func(s string) string {
			return strings.ReplaceAll(strings.ReplaceAll(s, "{{< ", "{{</* "), " >}}", " */>}}")
		},
	}

	tmpl, err := template.New(filepath.Base(dstPath)).Funcs(funcMap).Parse(tmpl)
	if err != nil {
		return fmt.Errorf("couldn't parse template: %w", err)
	}

	if err := tmpl.Execute(w, pageData); err != nil {
		return fmt.Errorf("couldn't execute template: %w", err)
	}

	if err := w.Flush(); err != nil {
		return fmt.Errorf("couldn't flush output buffer: %w", err)
	}

	if err := fsys.WriteFile(dstPath, buf.Bytes(), os.ModePerm); err != nil {
		return fmt.Errorf("couldn't write output file: %w", err)
	}

	return rulesErrs
}

func main() {
	const (
		requiredRepoDirPath = iota
		requiredsrcDirPath
		requiredDstPath
		requiredTotal
	)

	fs := flag.NewFlagSet(command, flag.ExitOnError)
	flag.Parse()

	if flag.NArg() != requiredTotal {
		usage(os.Stderr, fs)

		os.Exit(exit.UsageError)
	}

	repoDir := flag.Arg(requiredRepoDirPath)
	srcDirPath := flag.Arg(requiredsrcDirPath)
	dstPath := flag.Arg(requiredDstPath)

	fsys := rwfilefs.NewOSDirFS(repoDir)

	if err := generate(fsys, srcDirPath, dstPath); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(exit.RuntimeError)
	}
}
