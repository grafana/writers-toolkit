package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/grafana/writers-toolkit/link-checker/internal/check"
	"github.com/grafana/writers-toolkit/link-checker/internal/comment"
)

func main() {
	if len(os.Args) < 2 {
		printUsage(os.Stderr)
		os.Exit(2)
	}

	command := os.Args[1]
	args := os.Args[2:]

	var err error
	switch command {
	case "check":
		err = check.Run(args)
	case "comment":
		err = comment.Run(args)
	case "help", "-h", "--help":
		printUsage(os.Stdout)
		return
	default:
		fmt.Fprintf(os.Stderr, "unknown command %q\n\n", command)
		printUsage(os.Stderr)
		os.Exit(2)
	}

	if err == nil {
		return
	}
	if errors.Is(err, flag.ErrHelp) {
		return
	}

	fmt.Fprintf(os.Stderr, "broken-links %s failed: %v\n", command, err)
	os.Exit(1)
}

func printUsage(output *os.File) {
	_, _ = fmt.Fprintln(output, "Usage: broken-links <command> [flags]")
	_, _ = fmt.Fprintln(output)
	_, _ = fmt.Fprintln(output, "Commands:")
	_, _ = fmt.Fprintln(output, "  check    Crawl preview and write links report")
	_, _ = fmt.Fprintln(output, "  comment  Generate markdown comment from links report")
}
