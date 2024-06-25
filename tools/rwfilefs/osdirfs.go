//nolint:wrapcheck // This is a wrapper around os.DirFS.
package rwfilefs

import (
	"io/fs"
	"os"
	"path/filepath"
)

// OSDirFS is an implementation of os.DirFS that implements
// the RWFileFS interface.
type OSDirFS struct {
	fs.StatFS
	Root string
}

// NewOSDirFS returns a file system for the tree of files rooted at the directory dir.
//
// Note that all caveats of os.DirFS apply here, especially given the destructive
// nature of these methods.
func NewOSDirFS(root string) OSDirFS {
	dirFS := os.DirFS(root)

	return OSDirFS{
		//nolint:forcetypeassert
		// This is safe as os.DirFS implements Stat from Go 1.17.
		StatFS: dirFS.(fs.StatFS),
		Root:   root,
	}
}

// RemoveFile implements the RWFileFS interface.
func (fs OSDirFS) RemoveFile(name string) error {
	return os.Remove(filepath.Join(fs.Root, name))
}

// ReadFile implements the RWFileFS interface.
func (fs OSDirFS) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(filepath.Join(fs.Root, name))
}

// WriteFile implements the RWFileFS interface.
// It will create any parent directories of the named file before trying to write any data.
func (fs OSDirFS) WriteFile(name string, data []byte, mode fs.FileMode) error {
	path := filepath.Join(fs.Root, name)

	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		return err
	}

	return os.WriteFile(path, data, mode)
}
