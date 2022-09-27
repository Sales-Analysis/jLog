package jlog

import (
	"io"
	"os"
)

type jlog struct {
	location string // Folder with log files
}

// Create new jLog.
// The location variable sets the folder with log files.
func Init(location string) *jlog {
	return &jlog{
		location: location,
	}
}

// Info calls stdout to print to the logger.
func (j *jlog) Info(message string) {
	stdout(message)
}

// Warning is equivalent ot Info.
func (j *jlog) Warning(message string) {
	stdout(message)
}

// Error is equivalent ot Info.
func (j *jlog) Error(message string) {
	stdout(message)
}

// Stdout writes the output for a logging event.
func stdout(message string) {
	io.WriteString(os.Stdout, message)
}
