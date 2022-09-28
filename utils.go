package jlog
import "time"

// Check character in end of line.
func charEndOfLine(row string, char string) bool {
	return row[len(row)-1:] == char
}

// timeNow returns the current local time.
// Variable format to represent the date format.
func timeNow(format string) string {
	return time.Now().Format(format)
}
