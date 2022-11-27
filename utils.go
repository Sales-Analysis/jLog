package jlog

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

// Check character in end of line.
func charEndOfLine(row string, char string) bool {
	return row[len(row)-1:] == char
}

// timeNow returns the current local time.
// Variable format to represent the date format.
func timeNow(format string) string {
	return time.Now().Format(format)
}

// getPackageInfo returns name package and function.
func getPackageInfo(counter uintptr) (string, string){
	name := runtime.FuncForPC(counter).Name()
	strs := strings.Split(name, "/")
	info := strings.Split(strs[len(strs)-1], ".")
	return info[0], info[len(info)-1]
}

// makeFilename create name file. 
// Format <format>.log
func makeFilename(format string) string {
	t := timeNow(format)
	return fmt.Sprintf("%s.log", t)
}