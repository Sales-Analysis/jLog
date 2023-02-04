package jlog

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/Sales-Analysis/jLog/internal/filemanager"
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
	location  string // Folder with log files. Default value "logger".
	format    string // date format. Default value "2006-01-02 15:04:05".
	filename  string // format log file name. Сan be an empty string. Default value "20060102".
	separator string // message log separator.
	maxBytes  int    // max size of file.
	gotostd   bool   // log to stdout.
	gotofile  bool   // log to file.
}

// Create new jLog.
// The location variable sets the folder with log files.
func Init(envFile string) *jlog {
	loadDotEnv(envFile)
	gotostd, _ := strconv.ParseBool(os.Getenv("GOTOSTD"))
	gotofile, _ := strconv.ParseBool(os.Getenv("GOTOFILE"))
	maxBytes, _ := strconv.Atoi(os.Getenv("MAX_BYTES"))
	return &jlog{
		location:  os.Getenv("LOCATION"),
		format:    os.Getenv("FORMAT_TIME_LOG"),
		filename:  os.Getenv("FORMAT_FILENAME"),
		separator: os.Getenv("SEPARATOR"),
		maxBytes:  maxBytes,
		gotostd:   gotostd,
		gotofile:  gotofile,
	}
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
	timeString := timeNow(j.format)

	t := j.getColor(timeString, timeColor)
	pkg := j.getColor(packageName, packageColor)
	fun := j.getColor(funName, funColor)
	p := j.getPrefixColor(prefix)

	if j.gotostd {
		log := j.logTemplate(t, pkg, fun, p, message)
		io.WriteString(os.Stdout, log)
	}
	if j.gotofile {
		row := []string{timeString, packageName, funName, prefix, message}
		logStdout := j.logTemplateFile(row...)
		toFile(j.location, j.filename, logStdout, j.maxBytes)
	}
}

// prefixColor returns the colored status.
func (j *jlog) getPrefixColor(prefix string) string {
	color := getStatusColor(prefix)
	p := sepStr(prefix, j.separator)
	return fmt.Sprintf("%s%s%s", color, p, resetColor)
}

// getStatusColor returns the status color.
func getStatusColor(status string) string {
	switch status {
	case info:
		return infoColor
	case warning:
		return warningColor
	case err:
		return errorColor
	default:
		return dummyColor
	}
}

// getColor returns the colored string
func (j *jlog) getColor(str string, color string) string {
	str = sepStr(str, j.separator)
	return fmt.Sprintf("%s%s%s", color, str, resetColor)
}

// logTemplate returns a string in a specific format.
func (j *jlog) logTemplate(date string, pkg string, fun string, prefix string, message string) string {
	return fmt.Sprintf("%s%s%s%s: %s", date, pkg, fun, prefix, message)
}

// logTemplateFile returns a string in a specific format.
func (j *jlog) logTemplateFile(str ...string) string {
	return addSep(j.separator, str...)
}

// addStep add separator for str.
func addSep(sep string, str ...string) string {
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

// sepStr returns a delimited string.
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
func toFile(location string, logFormat string, message string, maxBytes int) {
	filemanager.CreateDir(location, false)

	filename := makeFilename(logFormat)

	var path string
	if charEndOfLine(location, "/") {
		path = location + filename
	} else {
		path = location + "/" + filename
	}

	backup(path, maxBytes)

	filemanager.Write(path, message)
}

func backup(path string, maxBytes int) {
	if maxBytes != 0 {
		size, _ := filemanager.GetSizeOfFile(path)
		if size >= int64(maxBytes) {
			p := fmt.Sprintf("%s.backup.log", strings.Split(path, ".log")[0])
			if _, err := os.Stat(p); err == nil {
				filemanager.GetToZip(p, strings.Split(path, ".log")[0])
			}
			_ = os.Rename(path, p)
		}
	}
}
