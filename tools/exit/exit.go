// Package exit provides consistent exit codes for command-line tools.
package exit

const (
	// Return codes inspired by buildifier tool.
	Success      = 0
	InputError   = 1
	UsageError   = 2
	RuntimeError = 3
	CheckError   = 4
)
