// Package rwfilefs extends the fs.FS interfaces to include the ability to write and remove files.
package rwfilefs

import (
	"io/fs"
)

type RemoveFileFS interface {
	fs.StatFS
	RemoveFile(name string) error
}

type WriteFileFS interface {
	fs.FS
	WriteFile(name string, data []byte, mode fs.FileMode) error
}

type RWFileFS interface {
	fs.ReadFileFS
	fs.StatFS
	RemoveFileFS
	WriteFileFS
}
