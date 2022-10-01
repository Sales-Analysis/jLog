package jlog

import (
	"fmt"
	"io"
	"os"
)

// Status constants
const (
	info    = "INFO"
	warning = "WARNING"
	err     = "ERROR"
	dummy   = "DUMMY"
)

// status colors
const (
	infoColor    = "\u001B[34;1m"
	warningColor = "\u001B[33m"
	errorColor   = "\u001B[31m"
	dummyColor   = "\u001B[35m"

	timeColor = "\u001b[32m"

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

func (j *jlog) Dummy(message string) {
	j.stdout(dummy, message)
}

// Stdout writes the output for a logging event.
func (j *jlog) stdout(prefix string, message string) {
	if !charEndOfLine(message, "\n") {
		message = message + "\n"
	}
	p := getPrefixColor(prefix)
	t := getTimeColor(timeNow(j.format))
	log := j.logTemplate(t, p, message)
	io.WriteString(os.Stdout, log)
}

// prefixColor returns the colored status
func getPrefixColor(prefix string) string {
	switch prefix {
	case info:
		return fmt.Sprintf("%s[%s]%s", infoColor, prefix, resetColor)
	case warning:
		return fmt.Sprintf("%s[%s]%s", warningColor, prefix, resetColor)
	case err:
		return fmt.Sprintf("%s[%s]%s", errorColor, prefix, resetColor)
	case dummy:
		return fmt.Sprintf("%s[%s]%s", dummyColor, prefix, resetColor)
	default:
		return fmt.Sprintf("[%s]", prefix)
	}
}

// getTimeColor returns the colored time
func getTimeColor(time string) string {
	return fmt.Sprintf("%s[%s]%s", timeColor, time, resetColor)
}

// logTemplate returns a string in a specific format.
func (j *jlog) logTemplate(date string, prefix string, message string) string {
	return fmt.Sprintf("%s%s: %s", date, prefix, message)
}
