package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/grafana/writers-toolkit/tools/exit"
	"gopkg.in/yaml.v3"
)

const command = "review"

func usage(w io.Writer, fs *flag.FlagSet) {
	fmt.Fprintf(w, "Usage of %s:\n", command)
	fs.PrintDefaults()

	fmt.Fprintln(w, "  <DIRECTORY>")
	fmt.Fprintln(w, "    	Path to documentation directory")
}

var errNotDateString = fmt.Errorf("not a date string")

const frontMatterDelim = "---"

type section int

const (
	sectionBeforeDelim section = iota
	sectionFrontMatter
	sectionBody
	countSections
)

type DateArgument struct {
	time.Time
}

func (d *DateArgument) String() string {
	return d.Time.Format(time.DateOnly)
}

func (d *DateArgument) Set(str string) error {
	if t, err := time.Parse(time.DateOnly, str); err != nil {
		return err
	} else {
		d.Time = t
	}

	return nil
}

func (d *DateArgument) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.ScalarNode {
		return errNotDateString
	}

	return d.Set(value.Value)
}

type metadata struct {
	ReviewDate DateArgument `yaml:"review_date"`
	Headless   bool         `yaml:"headless"`
}

func toReview(fsys fs.FS, before DateArgument) ([]string, error) {
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

		if frontMatter.ReviewDate.Before(before.Time) && !frontMatter.Headless {
			pages = append(pages, path)
		}

		return nil
	}); err != nil {
		return pages, fmt.Errorf("unable to walk directory: %w", err)
	}

	return pages, nil
}

func main() {
	const (
		requiredSrcDirPath = iota
		requiredArgCount
	)

	before := DateArgument{Time: time.Now().AddDate(0, 0, -90)}

	flag.Var(&before, "before", "YYYY-MM-DD date after which to include documents, defaults to ninety days ago.")

	flag.Parse()

	if flag.NArg() != requiredArgCount {
		usage(os.Stderr, flag.CommandLine)

		os.Exit(exit.UsageError)
	}

	srcDirPath := flag.Arg(requiredSrcDirPath)
	fsys := os.DirFS(srcDirPath)

	pages, err := toReview(fsys, before)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)

		os.Exit(exit.RuntimeError)
	}

	for _, page := range pages {
		fmt.Println(filepath.Join(srcDirPath, page) + ":1:1")
	}
}
