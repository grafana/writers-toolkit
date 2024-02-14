package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

const frontMatterDelim = "---"

type section int

const (
	sectionBeforeDelim section = iota
	sectionFrontMatter
	sectionBody
	countSections
)

type Date time.Time

func (d *Date) String() string {
	return time.Time(*d).Format(time.DateOnly)
}

func (d *Date) Set(str string) error {
	t, err := time.Parse(time.DateOnly, str)
	if err != nil {
		return err
	}

	*d = Date(t)

	return nil
}

type metadata struct {
	Date time.Time `yaml:"date"`
}

func main() {
	before := Date(time.Now().AddDate(0, 0, -90))

	flag.Var(&before, "before", "YYYY-MM-DD date after which to include documents, defaults to ninety days ago.")

	flag.Parse()

	if flag.NArg() != 1 {
		command := "review"
		if len(os.Args) != 0 {
			command = os.Args[0]
		}

		fmt.Fprintln(os.Stderr, "Usage of "+command)
		flag.PrintDefaults()

		fmt.Fprintln(os.Stderr, "  <DIRECTORY>")
		fmt.Fprintln(os.Stderr, "    	Path to documentation directory")

		os.Exit(2)
	}

	srcDir := flag.Arg(0)
	fsys := os.DirFS(srcDir)
	pages := []string{}

	if err := fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		if !strings.HasSuffix(path, ".md") {
			return nil
		}

		data, err := fs.ReadFile(fsys, path)
		if err != nil {
			return fmt.Errorf("unable to read file %q: %w", path, err)
		}

		sections := bytes.SplitN(data, []byte(frontMatterDelim), int(countSections))
		if len(sections) < int(countSections) {
			fmt.Fprintf(os.Stderr, "Skipping page: %s\n", path)

			return nil
		}

		var frontMatter metadata
		if err := yaml.Unmarshal(sections[sectionFrontMatter], &frontMatter); err != nil {
			return fmt.Errorf("unable to unmarshal front matter in %q: %w", path, err)
		}

		if frontMatter.Date.Before(time.Time(before)) {
			pages = append(pages, path)
		}

		return nil
	}); err != nil {
		fmt.Fprintf(os.Stderr, "Error traversing documentation: %v\n", err)
	}

	for _, page := range pages {
		fmt.Println(filepath.Join(srcDir, page) + ":1:1")
	}
}
