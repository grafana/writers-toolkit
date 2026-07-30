package main

import (
	"bytes"
	"embed"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed templates/*
var templates embed.FS

const command = "scaffold-docs"

func usage(w io.Writer, fs *flag.FlagSet) {
	fmt.Fprintf(w, "Usage of %s:\n", command)
	fs.PrintDefaults()

	fmt.Fprintln(w, "  <OWNER/REPO>")
	fmt.Fprintln(w, "    	GitHub repository in owner/repo format (e.g., grafana/my-repo)")
}

type options struct {
	removeOld                bool
	prTitle                  string
	prBody                   string
	branch                   string
	syncPublishMain          bool
	syncPublishNext          bool
	syncPublishRelease       bool
	syncDeployPreview        bool
	syncDocs                 bool
	syncDocsVariables        bool
	preserveTempDir          bool
}

func main() {
	var opts options

	flag.BoolVar(&opts.removeOld, "remove-old", false, "Remove older versions of template files before copying new ones")
	flag.StringVar(&opts.prTitle, "pr-title", "Add documentation infrastructure", "Title for the pull request")
	flag.StringVar(&opts.prBody, "pr-body", "This PR adds the standard documentation infrastructure files to enable building and testing documentation.", "Body for the pull request")
	flag.StringVar(&opts.branch, "branch", "add-docs-infrastructure", "Branch name for the changes")
	flag.BoolVar(&opts.syncPublishMain, "sync-publish-main", true, "Include publish-technical-documentation.yml workflow")
	flag.BoolVar(&opts.syncPublishNext, "sync-publish-next", false, "Include publish-technical-documentation-next.yml workflow")
	flag.BoolVar(&opts.syncPublishRelease, "sync-publish-release", false, "Include publish-technical-documentation-release.yml workflow")
	flag.BoolVar(&opts.syncDeployPreview, "sync-deploy-preview", false, "Include deploy-pr-preview.yml workflow")
	flag.BoolVar(&opts.syncDocs, "sync-docs", true, "Include docs directory templates (Makefile, docs.mk, make-docs)")
	flag.BoolVar(&opts.syncDocsVariables, "sync-docs-variables", true, "Include docs/variables.mk template")
	flag.BoolVar(&opts.preserveTempDir, "preserve-temp-dir", false, "Preserve temporary directory for debugging")

	flag.Usage = func() {
		usage(os.Stderr, flag.CommandLine)
	}

	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "Error: exactly one argument (owner/repo) is required")
		flag.Usage()
		os.Exit(2)
	}

	ownerRepo := flag.Arg(0)
	parts := strings.Split(ownerRepo, "/")
	if len(parts) != 2 {
		fmt.Fprintf(os.Stderr, "Error: invalid owner/repo format '%s'. Expected format: owner/repo\n", ownerRepo)
		os.Exit(1)
	}

	if err := run(ownerRepo, opts); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run(ownerRepo string, opts options) error {
	if err := checkGHCLI(); err != nil {
		return err
	}

	tempDir, err := os.MkdirTemp("", "setup-docs-*")
	if err != nil {
		return fmt.Errorf("failed to create temp directory: %w", err)
	}
	if !opts.preserveTempDir {
		defer os.RemoveAll(tempDir)
	} else {
		fmt.Printf("Temporary directory preserved at: %s\n", tempDir)
	}

	// Check if user has write access to the repository
	hasWriteAccess, err := checkWriteAccess(ownerRepo)
	if err != nil {
		return fmt.Errorf("failed to check repository access: %w", err)
	}

	var repoToClone string
	if hasWriteAccess {
		repoToClone = ownerRepo
		fmt.Printf("Cloning repository %s...\n", ownerRepo)
	} else {
		fmt.Printf("No write access to %s, creating fork...\n", ownerRepo)
		if err := createFork(ownerRepo); err != nil {
			return fmt.Errorf("failed to create fork: %w", err)
		}
		
		// Get the current user to construct fork repo name
		currentUser, err := getCurrentUser()
		if err != nil {
			return fmt.Errorf("failed to get current user: %w", err)
		}
		
		parts := strings.Split(ownerRepo, "/")
		repoName := parts[1]
		repoToClone = fmt.Sprintf("%s/%s", currentUser, repoName)
		fmt.Printf("Cloning fork %s...\n", repoToClone)
	}

	cloneCmd := exec.Command("gh", "repo", "clone", repoToClone, tempDir)
	cloneCmd.Stdout = os.Stdout
	cloneCmd.Stderr = os.Stderr
	if err := cloneCmd.Run(); err != nil {
		return fmt.Errorf("failed to clone repository: %w", err)
	}

	if err := os.Chdir(tempDir); err != nil {
		return fmt.Errorf("failed to change directory: %w", err)
	}

	fmt.Printf("Creating branch %s...\n", opts.branch)
	checkoutCmd := exec.Command("git", "checkout", "-b", opts.branch)
	if err := checkoutCmd.Run(); err != nil {
		return fmt.Errorf("failed to create branch: %w", err)
	}

	if opts.removeOld {
		fmt.Println("Removing old documentation files...")
		if err := removeOldFiles(tempDir); err != nil {
			return fmt.Errorf("failed to remove old files: %w", err)
		}
	}

	fmt.Println("Copying documentation templates...")
	if err := copyTemplates(tempDir, ownerRepo, opts); err != nil {
		return fmt.Errorf("failed to copy templates: %w", err)
	}

	fmt.Println("Staging changes...")
	addCmd := exec.Command("git", "add", ".")
	if err := addCmd.Run(); err != nil {
		return fmt.Errorf("failed to stage changes: %w", err)
	}

	fmt.Println("Committing changes...")
	commitCmd := exec.Command("git", "commit", "-m", "Add documentation infrastructure")
	commitCmd.Stdout = os.Stdout
	commitCmd.Stderr = os.Stderr
	if err := commitCmd.Run(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() == 1 {
			fmt.Println("No changes to commit - repository may already have documentation setup")
			return nil
		}
		return fmt.Errorf("failed to commit changes: %w", err)
	}

	fmt.Println("Pushing branch...")
	pushCmd := exec.Command("git", "push", "-u", "origin", opts.branch)
	pushCmd.Stdout = os.Stdout
	pushCmd.Stderr = os.Stderr
	if err := pushCmd.Run(); err != nil {
		return fmt.Errorf("failed to push branch: %w", err)
	}

	// Create draft PR
	fmt.Println("Creating draft pull request...")
	prArgs := []string{"pr", "create", "--draft", "--title", opts.prTitle, "--body", opts.prBody}
	
	// If using a fork, specify the base repository
	if !hasWriteAccess {
		prArgs = append(prArgs, "--repo", ownerRepo)
	}
	
	prCmd := exec.Command("gh", prArgs...)
	prCmd.Stdout = os.Stdout
	prCmd.Stderr = os.Stderr
	if err := prCmd.Run(); err != nil {
		return fmt.Errorf("failed to create pull request: %w", err)
	}

	fmt.Println("\nSuccess! Draft PR created.")
	return nil
}

func checkGHCLI() error {
	if _, err := exec.LookPath("gh"); err != nil {
		return fmt.Errorf("GitHub CLI (gh) is not installed. Please install it from https://cli.github.com/")
	}

	authCmd := exec.Command("gh", "auth", "status")
	if err := authCmd.Run(); err != nil {
		return fmt.Errorf("GitHub CLI is not authenticated. Please run 'gh auth login'")
	}

	return nil
}

func checkWriteAccess(ownerRepo string) (bool, error) {
	cmd := exec.Command("gh", "api", fmt.Sprintf("repos/%s", ownerRepo), "--jq", ".permissions.push")
	output, err := cmd.Output()
	if err != nil {
		return false, err
	}
	
	return strings.TrimSpace(string(output)) == "true", nil
}

func createFork(ownerRepo string) error {
	cmd := exec.Command("gh", "repo", "fork", ownerRepo, "--clone=false")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func getCurrentUser() (string, error) {
	cmd := exec.Command("gh", "api", "user", "--jq", ".login")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	
	return strings.TrimSpace(string(output)), nil
}

func removeOldFiles(repoDir string) error {
	filesToRemove := []string{
		"docs/Makefile",
		"docs/docs.mk",
		"docs/variables.mk",
		"docs/make-docs",
		".github/workflows/docs-ci.yml",
		".github/workflows/publish-technical-documentation.yml",
		".github/workflows/publish-technical-documentation-next.yml",
		".github/workflows/update-make-docs.yml",
	}

	for _, file := range filesToRemove {
		path := filepath.Join(repoDir, file)
		if _, err := os.Stat(path); err == nil {
			fmt.Printf("  Removing %s\n", file)
			if err := os.Remove(path); err != nil {
				return fmt.Errorf("failed to remove %s: %w", file, err)
			}
		}
	}

	return nil
}

func copyTemplates(repoDir, ownerRepo string, opts options) error {
	parts := strings.Split(ownerRepo, "/")
	repoName := parts[1]

	data := struct {
		RepoName  string
		OwnerRepo string
	}{
		RepoName:  repoName,
		OwnerRepo: ownerRepo,
	}

	return fs.WalkDir(templates, "templates", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		relPath, err := filepath.Rel("templates", path)
		if err != nil {
			return err
		}

		// Skip publish workflows based on options
		if strings.Contains(relPath, "publish-technical-documentation.yml") && !opts.syncPublishMain {
			return nil
		}
		if strings.Contains(relPath, "publish-technical-documentation-next.yml") && !opts.syncPublishNext {
			return nil
		}
		if strings.Contains(relPath, "publish-technical-documentation-release.yml") && !opts.syncPublishRelease {
			return nil
		}
		if strings.Contains(relPath, "deploy-pr-preview.yml") && !opts.syncDeployPreview {
			return nil
		}
		
		// Skip docs templates based on options
		if strings.HasPrefix(relPath, "docs/") && !opts.syncDocs {
			return nil
		}
		if relPath == "docs/variables.mk" && !opts.syncDocsVariables {
			return nil
		}

		content, err := templates.ReadFile(path)
		if err != nil {
			return fmt.Errorf("failed to read template %s: %w", path, err)
		}

		tmpl, err := template.New(filepath.Base(path)).Delims("⟨⟨⟨", "⟩⟩⟩").Parse(string(content))
		if err != nil {
			return fmt.Errorf("failed to parse template %s: %w", path, err)
		}

		var buf bytes.Buffer
		if err := tmpl.Execute(&buf, data); err != nil {
			return fmt.Errorf("failed to execute template %s: %w", path, err)
		}

		destPath := filepath.Join(repoDir, relPath)
		destDir := filepath.Dir(destPath)

		if err := os.MkdirAll(destDir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", destDir, err)
		}

		fmt.Printf("  Creating %s\n", relPath)
		if err := os.WriteFile(destPath, buf.Bytes(), 0644); err != nil {
			return fmt.Errorf("failed to write file %s: %w", destPath, err)
		}

		if strings.Contains(relPath, "make-docs") {
			if err := os.Chmod(destPath, 0755); err != nil {
				return fmt.Errorf("failed to make %s executable: %w", destPath, err)
			}
		}

		return nil
	})
}
