package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/gohugoio/hugo/parser/pageparser"
	"github.com/grafana/writers-toolkit/tools/exit"
	"github.com/grafana/writers-toolkit/tools/hugo"
	"github.com/grafana/writers-toolkit/tools/rwfilefs"
)

const command = "update-front-matter"

func usage(w io.Writer, fs *flag.FlagSet) {
	fmt.Fprintf(w, "Usage of %s:\n", command)
	fs.PrintDefaults()

	fmt.Fprintln(w, "  <PATH>")
	fmt.Fprintln(w, "    	Path to the documentation source file")

	fmt.Fprintln(w, "  [<KEY>=<VALUE>...]")
	fmt.Fprintln(w, "    	Updates to make to the front matter. Only string values are supported.")
}

func update(fsys rwfilefs.RWFileFS, path string, updates map[string]string) error {
	data, err := fsys.ReadFile(path)
	if err != nil {
		return fmt.Errorf("couldn't read file: %w", err)
	}

	cfm, err := pageparser.ParseFrontMatterAndContent(bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("couldn't parse file: %w", err)
	}

	for key, value := range updates {
		cfm.FrontMatter[key] = value
	}

	data, err = hugo.Unparse(cfm)
	if err != nil {
		return fmt.Errorf("couldn't unparse file: %w", err)
	}

	if err := fsys.WriteFile(path, data, os.ModePerm); err != nil {
		return fmt.Errorf("couldn't write file: %w", err)
	}

	return nil
}

func main() {
	const (
		requiredPath = iota
		requiredTotal
	)

	const (
		argKey = iota
		argValue
		argParts
	)

	fs := flag.NewFlagSet(command, flag.ExitOnError)
	flag.Parse()

	if flag.NArg() < requiredTotal {
		usage(os.Stderr, fs)

		os.Exit(exit.UsageError)
	}

	updates := make(map[string]string)

	for i, arg := range flag.Args() {
		if i == requiredPath {
			continue
		}

		parts := strings.SplitN(arg, "=", argParts)
		if len(parts) != argParts {
			fmt.Fprintf(os.Stderr, "Error: invalid update argument: %s\n", arg)

			os.Exit(exit.InputError)
		}

		updates[parts[argKey]] = parts[argValue]
	}

	path := flag.Arg(requiredPath)

	dir := filepath.Dir(path)
	fsys := rwfilefs.NewOSDirFS(dir)

	path, err := filepath.Rel(dir, path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(exit.RuntimeError)
	}

	if err := update(fsys, path, updates); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(exit.RuntimeError)
	}
}
