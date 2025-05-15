package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/google/go-github/v72/github"
	"github.com/grafana/writers-toolkit/tools/exit"
	"github.com/grafana/writers-toolkit/vale/tools/cmd/filter-sarif/sarif"
)

const tool = "filter-sarif"

type (
	Arguments struct {
		owner     string
		repo      string
		prNumber  int
		sarifPath string
	}
	Options struct{}
)

// usage prints the command line tool usage information.
func usage() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", tool)
	fmt.Fprintln(flag.CommandLine.Output())
	fmt.Fprintln(flag.CommandLine.Output(), "Options:")
	flag.PrintDefaults()
	fmt.Fprintln(flag.CommandLine.Output())
	fmt.Fprintln(flag.CommandLine.Output(), "Arguments:")
	fmt.Fprintln(flag.CommandLine.Output(), "  <OWNER>")
	fmt.Fprintln(flag.CommandLine.Output(), "    	GitHub repository owner.")
	fmt.Fprintln(flag.CommandLine.Output(), "  <REPO>")
	fmt.Fprintln(flag.CommandLine.Output(), "    	GitHub repository name.")
	fmt.Fprintln(flag.CommandLine.Output(), "  <PULL REQUEST NUMBER>")
	fmt.Fprintln(flag.CommandLine.Output(), "    	GitHub pull request number.")
	fmt.Fprintln(flag.CommandLine.Output(), "  <SARIF PATH>")
	fmt.Fprintln(flag.CommandLine.Output(), "    	Path to SARIF file.")
}

func main() {
	_, args := parseCommandLine()

	sarifFile, err := sarif.NewFromFile(args.sarifPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading SARIF file: %v\n", err)
		os.Exit(exit.RuntimeError)
	}

	token := os.Getenv("GITHUB_TOKEN")
	ctx := context.Background()
	client := github.NewClient(nil).WithAuthToken(token)

	filtered, hasResults, err := filterSARIFByPR(ctx, client, sarifFile, args.owner, args.repo, args.prNumber)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error filtering SARIF: %v\n", err)
		os.Exit(exit.RuntimeError)
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")

	if err := enc.Encode(filtered); err != nil {
		fmt.Fprintf(os.Stderr, "Error encoding filtered SARIF: %v\n", err)
		os.Exit(exit.RuntimeError)
	}

	if hasResults {
		os.Exit(exit.CheckError)
	}
}

func parseCommandLine() (Options, Arguments) {
	opts := parseOptions()
	args := parseArguments()

	return opts, args
}

func parseOptions() Options {
	var opts Options

	flag.Parse()

	return opts
}

func parseArguments() Arguments {
	var args Arguments

	const (
		requiredOwner = iota
		requiredRepo
		requiredPRNumber
		requiredSARIFPath
		requiredArgTotal
	)

	if flag.NArg() != requiredArgTotal {
		fmt.Fprintf(flag.CommandLine.Output(), "Error: Missing required arguments, got %d\n", flag.NArg())
		flag.CommandLine.SetOutput(os.Stderr)
		usage()

		os.Exit(exit.UsageError)
	}

	args.owner = flag.Arg(requiredOwner)
	args.repo = flag.Arg(requiredRepo)
	args.prNumber, _ = strconv.Atoi(flag.Arg(requiredPRNumber))
	args.sarifPath = flag.Arg(requiredSARIFPath)

	return args
}
