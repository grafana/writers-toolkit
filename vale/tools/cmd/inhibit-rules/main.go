package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/grafana/writers-toolkit/tools/exit"
)

const command = "inhibit-rules"

func usage(w io.Writer, fs *flag.FlagSet) {
	fmt.Fprintf(w, "Usage of %s:\n", command)
	fs.SetOutput(w)
	fs.PrintDefaults()

	fmt.Fprintf(w, "\n")
	fmt.Fprintln(w, "Reads JSON formatted error messages from standard input and filters them based on rule precedence.")
}

func main() {
	fs := flag.NewFlagSet(command, flag.ExitOnError)
	flag.Parse()

	if flag.NArg() > 0 {
		usage(os.Stderr, fs)
		os.Exit(exit.UsageError)
	}

	r := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(r)
	diags := make([]diagnostic, 0)

	var errs bool
	for scanner.Scan() {
		diag := diagnostic{}
		if err := json.Unmarshal(scanner.Bytes(), &diag); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			errs = true
		}

		diags = append(diags, diag)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		errs = true
	}

	diags = inhibit(diags)
	for _, d := range diags {
		if err := json.NewEncoder(os.Stdout).Encode(d); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			errs = true
		}
	}

	if errs {
		os.Exit(exit.RuntimeError)
	}
}

// {"message": "[{{ $check }}] {{ $message | jsonEscape }}", "location": {"path": "{{ $path }}", "range": {"start": {"line": {{ $line }}, "column": {{ $col }}}}}, "severity": "{{ $error }}"}
type diagnostic struct {
	Message  string `json:"message"`
	Location struct {
		Path  string `json:"path"`
		Range struct {
			Start struct {
				Line   int `json:"line"`
				Column int `json:"column"`
			} `json:"start"`
		} `json:"range"`
	} `json:"location"`
	Severity string `json:"severity"`
}

// Each error has the rule name in square brackets at the beginning of the message.
// For example: {"message": "[Grafana.Spelling] ..." ...}
func (d diagnostic) Rule() string {
	if len(d.Message) == 0 || d.Message[0] != '[' {
		return ""
	}

	end := strings.Index(d.Message, "]")
	if end == -1 {
		return ""
	}

	return d.Message[1:end]
}

func (d diagnostic) LocationKey() string {
	return fmt.Sprintf("%s:%d:%d", d.Location.Path, d.Location.Range.Start.Line, d.Location.Range.Start.Column)
}

func inhibit(in []diagnostic) []diagnostic {
	filtered := make(map[string]diagnostic)
	precedence := map[string]int{
		"Grafana.ProductPossessives": 0,
		"Grafana.WordList":           1,
		"Grafana.Spelling":           2,
	}

	for _, d := range in {
		if _, ok := filtered[d.LocationKey()]; !ok || precedence[d.Rule()] < precedence[filtered[d.LocationKey()].Rule()] {
			filtered[d.LocationKey()] = d
		}
	}

	out := make([]diagnostic, 0, len(filtered))
	for _, d := range filtered {
		out = append(out, d)
	}

	return out
}
