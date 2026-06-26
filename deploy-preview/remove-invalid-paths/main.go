package main

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const invalidArtifactPathCharacters = "\":<>|*?\r\n"

func main() {
	distPath := flag.String("dist", "dist", "path to the built dist directory")
	flag.Parse()

	if err := run(*distPath); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(root string) error {
	removed, err := removeInvalidFiles(root)
	if err != nil {
		return err
	}

	if len(removed) == 0 {
		fmt.Printf("No invalid artifact paths found in %s\n", root)
		return nil
	}

	fmt.Printf("Removed %d file(s) with invalid artifact paths from %s\n", len(removed), root)
	for _, relPath := range removed {
		fmt.Println(relPath)
	}

	return nil
}

func removeInvalidFiles(root string) ([]string, error) {
	info, err := os.Stat(root)
	if err != nil {
		return nil, fmt.Errorf("stat dist directory: %w", err)
	}

	if !info.IsDir() {
		return nil, fmt.Errorf("dist path is not a directory: %s", root)
	}

	var toRemove []string

	err = filepath.WalkDir(root, func(path string, dirEntry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}

		if dirEntry.IsDir() {
			return nil
		}

		relPath, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}

		if hasInvalidArtifactPathCharacters(filepath.ToSlash(relPath)) {
			toRemove = append(toRemove, path)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("walk dist directory: %w", err)
	}

	sort.Strings(toRemove)

	removed := make([]string, 0, len(toRemove))
	for _, path := range toRemove {
		relPath, err := filepath.Rel(root, path)
		if err != nil {
			return nil, err
		}

		if err := os.Remove(path); err != nil && !errors.Is(err, os.ErrNotExist) {
			return nil, fmt.Errorf("remove invalid file %s: %w", path, err)
		}

		removed = append(removed, filepath.ToSlash(relPath))
	}

	if err := removeEmptyDirectories(root); err != nil {
		return nil, err
	}

	return removed, nil
}

func hasInvalidArtifactPathCharacters(path string) bool {
	return strings.ContainsAny(path, invalidArtifactPathCharacters)
}

func removeEmptyDirectories(root string) error {
	var directories []string

	err := filepath.WalkDir(root, func(path string, dirEntry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}

		if dirEntry.IsDir() && path != root {
			directories = append(directories, path)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("walk directories for cleanup: %w", err)
	}

	sort.Slice(directories, func(i, j int) bool {
		return directoryDepth(directories[i]) > directoryDepth(directories[j])
	})

	for _, dir := range directories {
		entries, err := os.ReadDir(dir)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				continue
			}

			return fmt.Errorf("read directory %s: %w", dir, err)
		}

		if len(entries) == 0 {
			if err := os.Remove(dir); err != nil && !errors.Is(err, os.ErrNotExist) {
				return fmt.Errorf("remove empty directory %s: %w", dir, err)
			}
		}
	}

	return nil
}

func directoryDepth(path string) int {
	return strings.Count(filepath.Clean(path), string(os.PathSeparator))
}
