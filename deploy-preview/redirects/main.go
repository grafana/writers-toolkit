package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func main() {
	manifestPath := flag.String("manifest", "dist/redirects.txt", "path to Hugo's generated redirects manifest")
	outputPath := flag.String("output", "dist/redirects.conf", "path to the generated nginx redirects")
	flag.Parse()

	if err := run(*manifestPath, *outputPath); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(manifestPath, outputPath string) error {
	manifest, err := os.Open(manifestPath)
	if err != nil {
		if os.IsNotExist(err) {
			output, createErr := os.Create(outputPath)
			if createErr != nil {
				return fmt.Errorf("create empty nginx redirects: %w", createErr)
			}
			if closeErr := output.Close(); closeErr != nil {
				return fmt.Errorf("close empty nginx redirects: %w", closeErr)
			}
			return nil
		}
		return fmt.Errorf("open redirect manifest: %w", err)
	}
	defer func() { _ = manifest.Close() }()

	output, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("create nginx redirects: %w", err)
	}

	if _, err := convert(manifest, output); err != nil {
		_ = output.Close()
		return fmt.Errorf("convert redirect manifest: %w", err)
	}
	if err := output.Close(); err != nil {
		return fmt.Errorf("close nginx redirects: %w", err)
	}
	return nil
}

func convert(r io.Reader, w io.Writer) (int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	count := 0
	for scanner.Scan() {
		line := strings.TrimSuffix(scanner.Text(), "\r")
		if strings.TrimSpace(line) == "" {
			continue
		}
		source, destination, ok := strings.Cut(line, "\t")
		if !ok || source == "" || destination == "" || !strings.HasPrefix(source, "/") || !strings.HasPrefix(destination, "/") {
			return count, fmt.Errorf("invalid redirect manifest line %q", line)
		}

		pattern := "^" + regexp.QuoteMeta(source) + "$"
		if strings.HasSuffix(source, "/") {
			pattern = "^" + regexp.QuoteMeta(strings.TrimSuffix(source, "/")) + "/?$"
		}
		if _, err := fmt.Fprintf(w, "rewrite %q %q permanent;\n", pattern, destination); err != nil {
			return count, err
		}
		count++
	}
	return count, scanner.Err()
}
