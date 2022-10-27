package jlog

import (
	"fmt"
	"io"
	"os"
	"runtime"
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

// default value format log file name
const defaultFilename = "20060102"

type jlog struct {
	location string // Folder with log files
	format   string // date format
	filename string // format log file name. Сan be an empty string. default value "20060102".
}

// Create new jLog.
// The location variable sets the folder with log files.
func Init(location string, format string, filename string) *jlog {
	return &jlog{
		location: location,
		format:   format,
		filename: filename,
	}
}

// Info calls stdout to print to the logger.
func (j *jlog) Info(message string) {
	counter, _, _, _ := runtime.Caller(1)
	j.stdout(dummy, message, counter)
}

// Warning is equivalent ot Info.
func (j *jlog) Warning(message string) {
	counter, _, _, _ := runtime.Caller(1)
	j.stdout(dummy, message, counter)
}

// Error is equivalent ot Info.
func (j *jlog) Error(message string) {
	counter, _, _, _ := runtime.Caller(1)
	j.stdout(dummy, message, counter)
}

// Dummy is useless log.
func (j *jlog) Dummy(message string) {
	counter, _, _, _ := runtime.Caller(1)
	j.stdout(dummy, message, counter)
}

// Stdout writes the output for a logging event.
func (j *jlog) stdout(prefix string, message string, counter uintptr) {
	if !charEndOfLine(message, "\n") {
		message = message + "\n"
	}
	packageName, funcName := getPackageInfo(counter)
	fmt.Println(packageName, funcName)
	p := getPrefixColor(prefix)
	t := getTimeColor(timeNow(j.format))
	log := j.logTemplate(t, p, message)
	logStdout := j.logTemplateFile(timeNow(j.format), prefix, message)
	io.WriteString(os.Stdout, log)
	toFile(j.location, j.filename, logStdout)
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

// logTemplateFile returns a string in a specific format.
func (j *jlog) logTemplateFile(date string, prefix string, message string) string {
	return fmt.Sprintf("[%s][%s]: %s", date, prefix, message)
}

// toFile write log to file
func toFile(location string, logFormat string, message string) {
	createDir(location, false)
	
	if logFormat == "" {
		logFormat = defaultFilename
	}
	filename := makeFilename(logFormat)
	
	var path string
	if charEndOfLine(location, "/") {
		path = location + filename
	} else {
		path = location + "/" + filename
	}
	write(path, message)
}
