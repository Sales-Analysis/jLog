package jlog

import (
	"io"
	"os"
)

type jlog struct {
	location string // Folder with log files
	format string // date format
}

// Create new jLog.
// The location variable sets the folder with log files.
func Init(location string, format string) *jlog {
	return &jlog{
		location: location,
		format: format,
	}
}

// Info calls stdout to print to the logger.
func (j *jlog) Info(message string) {
	j.stdout(message)
}

// Warning is equivalent ot Info.
func (j *jlog) Warning(message string) {
	j.stdout(message)
}

// Error is equivalent ot Info.
func (j *jlog) Error(message string) {
	j.stdout(message)
}

// Stdout writes the output for a logging event.
func (j *jlog) stdout(message string) {
	if !charEndOfLine(message, "\n") {
		message = message + "\n"
	}
	message = timeNow(j.format) + " " + message
	io.WriteString(os.Stdout, message)
}
