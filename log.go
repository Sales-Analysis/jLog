package jlog

import (
	"fmt"
	"io"
	"os"
)

// Status constants
const (
	info 	= "INFO"
	warning = "WARNING"
	err     = "ERROR"
)

// status colors
const (
	infoColor = "\u001B[34;1m"
	warningColor = "\u001B[33m"
	errorColor = "\u001B[31m"

	// reset color
	resetColor = "\u001B[0m"

)

type jlog struct {
	location string // Folder with log files
	format   string // date format
}

// Create new jLog.
// The location variable sets the folder with log files.
func Init(location string, format string) *jlog {
	return &jlog{
		location: location,
		format:   format,
	}
}

// Info calls stdout to print to the logger.
func (j *jlog) Info(message string) {
	j.stdout(info, message)
}

// Warning is equivalent ot Info.
func (j *jlog) Warning(message string) {
	j.stdout(warning, message)
}

// Error is equivalent ot Info.
func (j *jlog) Error(message string) {
	j.stdout(err, message)
}

// Stdout writes the output for a logging event.
func (j *jlog) stdout(prefix string, message string) {
	if !charEndOfLine(message, "\n") {
		message = message + "\n"
	}
	p := prefixColor(prefix)
	log := j.logTemplate(timeNow(j.format), p, message)
	io.WriteString(os.Stdout, log)
}

func prefixColor(prefix string) string {
	switch prefix {
	case info:
		return fmt.Sprintf("%s[%s]%s", infoColor, prefix, resetColor)
	case warning:
		return fmt.Sprintf("%s[%s]%s", warningColor, prefix, resetColor)
	case err:
		return fmt.Sprintf("%s[%s]%s", errorColor, prefix, resetColor)
	}
	return prefix
}

// logTemplate returns a string in a specific format.
func (j *jlog) logTemplate(date string, prefix string, message string) string {
	return fmt.Sprintf("[%s]%s: %s", date, prefix, message)
}
