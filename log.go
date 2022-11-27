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
	infoColor    = "\u001B[34m"
	warningColor = "\u001B[33m"
	errorColor   = "\u001B[31m"
	dummyColor   = "\u001B[35m"

	timeColor    = "\u001b[32m"
	packageColor = "\u001b[35;1m"
	funColor     = "\u001b[34;1m"
	// reset color
	resetColor = "\u001B[0m"
)

type jlog struct {
	location string // Folder with log files. Default value "logger".
	format   string // date format. Default value "2006-01-02 15:04:05".
	filename string // format log file name. Сan be an empty string. Default value "20060102".
}

// Create new jLog.
// The location variable sets the folder with log files.
func Init(envFile string) *jlog {
	err := dotenv.Load(envFile)
	if err != nil {
		fmt.Printf("%s.\nDefault parameters are assigned.\n", err)
		setDefaultParams()
	}
	return &jlog{
		location: os.Getenv("LOCATION"),
		format:   os.Getenv("FORMAT_TIME_LOG"),
		filename: os.Getenv("FORMAT_FILENAME"),
	}
}

// Set default parameters.
func setDefaultParams() {
	// Folder with log files
	os.Setenv("LOCATION", "logger")
	// Format log file name
	os.Setenv("FORMAT_FILENAME", "20060102")
	// Format log
	os.Setenv("FORMAT_TIME_LOG", "2006-01-02 15:04:05")
	// Separator
	os.Setenv("SEPARATOR", "[]")
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
	time_string := timeNow(j.format)
	t := getColor(time_string, timeColor)
	pkg := getColor(packageName, packageColor)
	fun := getColor(funName, funColor)
	p := getPrefixColor(prefix)

	log := j.logTemplate(t, pkg, fun, p, message)
	row := []string{timeNow(j.format), packageName, funName, prefix, message}
	logStdout := j.logTemplateFile(row...)

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
func (j *jlog) logTemplateFile(str ...string) string {
	sep := os.Getenv("SEPARATOR")
	return addSep(sep, str...)
}

// addStep add separator for str. 
func addSep(sep string, str ...string) string{
	row := ""
	for i, v := range str {
		if i != (len(str) - 1) {
			v = sepStr(v, sep)	
		} else {
			v = ": " + v
		}
		row += v
	}
	return row
}

// sep Str returns a delimited string.
// If the length of the separator is equal to 1,
// the separator is placed at the beginning.
// The length of the separator can not be more than two
func sepStr(str string, sep string) string {
	charSep := []rune(sep)
	s := ""
	if len(sep) > 1 { 
		s = string(charSep[0]) + str + string(charSep[1])
	} else {
		s = string(charSep[0]) + str
	}
	return s
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
