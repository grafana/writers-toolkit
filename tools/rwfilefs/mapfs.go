package rwfilefs

import (
	"io/fs"
	"testing/fstest"
	"time"
)

// MapFS is an implementation of fstest.MapFS that implements
// the RWFileFS interface.
type MapFS struct {
	fstest.MapFS
}

// WriteFile implements the RWFileFS interface.
func (fs MapFS) WriteFile(name string, data []byte, mode fs.FileMode) error {
	fs.MapFS[name] = &fstest.MapFile{
		Data:    data,
		Mode:    mode,
		ModTime: time.Now(),
		Sys:     nil,
	}

	return nil
}

// RemoveFile implements the RWFileFS interface.
func (fs MapFS) RemoveFile(name string) error {
	delete(fs.MapFS, name)

	return nil
}
