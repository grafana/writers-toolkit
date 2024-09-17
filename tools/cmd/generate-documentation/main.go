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
	"regexp"
	"strings"
	"text/template"
	"time"

	"github.com/grafana/writers-toolkit/tools/exit"
	"github.com/grafana/writers-toolkit/tools/rwfilefs"
	"gopkg.in/yaml.v3"
)

const command = "generate-documentation"

var errProcessingRule = errors.New("error processing rule")

type PageData struct {
	ReviewDate string `yaml:"review_date"`
	Rules      []Rule `yaml:"rules"`
	Title      string `yaml:"title"`
}

type Rule struct {
	Extends string            `yaml:"extends"`
	Level   string            `yaml:"level"`
	Link    string            `yaml:"link"`
	Message string            `yaml:"message"`
	Name    string            `yaml:"name"`
	Swap    map[string]string `yaml:"swap,omitempty"`
	Tokens  []string          `yaml:"tokens,omitempty"`
}

func usage(w io.Writer, fs *flag.FlagSet) {
	fmt.Fprintf(w, "Usage of %s:\n", command)
	fs.PrintDefaults()

	fmt.Fprintln(w, "  <REPOSITORY PATH>")
	fmt.Fprintln(w, "    	Path to the repository root.")

	fmt.Fprintln(w, "  <TEMPLATE PATH>")
	fmt.Fprintln(w, "    	Path to the Vale rules directory, relative to <REPOSITORY PATH>.")

	fmt.Fprintln(w, "  <SOURCE PATH>")
	fmt.Fprintln(w, "    	Path to the Vale rules directory, relative to <REPOSITORY PATH>.")

	fmt.Fprintln(w, "  <OUTPUT PATH>")
	fmt.Fprintln(w, "    	Path to the output file, relative to <REPOSITORY PATH>.")
}

func generate(fsys rwfilefs.RWFileFS, templateDirPath, srcDirPath, dstPath string) error {
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
		"codify": func(s string) string {
			return regexp.MustCompile("'?(<[A-Z -_]*>)'?").ReplaceAllString(s, "_`$1`_")
		},
		"escapeShortcodes": func(s string) string {
			return strings.ReplaceAll(
				strings.ReplaceAll(
					strings.ReplaceAll(
						strings.ReplaceAll(s, "{{<", "{{</*"),
						">}}", "*/>}}"),
					"{{%", "{{%/*"),
				"%}}", "*/%}}")
		},
		"escapeForTable": func(s string) string {
			return strings.ReplaceAll(s, "|", "\\|")
		},
		"format": func(s string) string {
			return strings.ReplaceAll(s, "%s", "<CURRENT TEXT>")
		},
	}

	pageTmpl, err := template.New("page.tmpl").Funcs(funcMap).ParseFS(fsys, filepath.Join(templateDirPath, "*.tmpl"))
	if err != nil {
		return fmt.Errorf("couldn't parse page template: %w", err)
	}

	if err := pageTmpl.Execute(w, pageData); err != nil {
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
		requiredTemplateDirPath
		requiredSrcDirPath
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
	templateDirPath := flag.Arg(requiredTemplateDirPath)
	srcDirPath := flag.Arg(requiredSrcDirPath)
	dstPath := flag.Arg(requiredDstPath)

	fsys := rwfilefs.NewOSDirFS(repoDir)

	if err := generate(fsys, templateDirPath, srcDirPath, dstPath); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(exit.RuntimeError)
	}
}
