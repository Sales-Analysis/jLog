package jlog

import (
	"fmt"
	"io"
	"os"
	"runtime"

	"github.com/Sales-Analysis/jLog/internal/dotenv"
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

	timeColor    = "\u001b[32m"
	packageColor = "\u001b[35;1m"
	funColor     = "\u001b[34;1m"
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
func Init(location string, format string, envFile string) *jlog {
	err := dotenv.Load(envFile)
	if err != nil {
		fmt.Printf("%s.\nDefault parameters are assigned.\n", err)
		setDefaultParams()
	}
	return &jlog{
		location: location,
		format:   format,
		filename: os.Getenv("FORMAT_FILENAME"),
	}
}

// Set default parameters.
func setDefaultParams() {
	os.Setenv("FORMAT_FILENAME", defaultFilename)
}

// Info calls stdout to print to the logger.
func (j *jlog) Info(message string) {
	counter, _, _, _ := runtime.Caller(1)
	j.stdout(info, message, counter)
}

// Warning is equivalent ot Info.
func (j *jlog) Warning(message string) {
	counter, _, _, _ := runtime.Caller(1)
	j.stdout(warning, message, counter)
}

// Error is equivalent ot Info.
func (j *jlog) Error(message string) {
	counter, _, _, _ := runtime.Caller(1)
	j.stdout(err, message, counter)
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
	packageName, funName := getPackageInfo(counter)

	t := getColor(timeNow(j.format), timeColor)
	pkg := getColor(packageName, packageColor)
	fun := getColor(funName, funColor)
	p := getPrefixColor(prefix)

	log := j.logTemplate(t, pkg, fun, p, message)
	logStdout := j.logTemplateFile(timeNow(j.format), packageName, funName, prefix, message)

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

// getColor returns the colored string
func getColor(str string, color string) string {
	return fmt.Sprintf("%s[%s]%s", color, str, resetColor)
}

// logTemplate returns a string in a specific format.
func (j *jlog) logTemplate(date string, pkg string, fun string, prefix string, message string) string {
	return fmt.Sprintf("%s%s%s%s: %s", date, pkg, fun, prefix, message)
}

// logTemplateFile returns a string in a specific format.
func (j *jlog) logTemplateFile(date string, pkg string, fun string, prefix string, message string) string {
	return fmt.Sprintf("[%s][%s][%s][%s]: %s", date, pkg, fun, prefix, message)
}

// toFile write log to file
func toFile(location string, logFormat string, message string) {
	createDir(location, false)

	filename := makeFilename(logFormat)

	var path string
	if charEndOfLine(location, "/") {
		path = location + filename
	} else {
		path = location + "/" + filename
	}
	write(path, message)
}
