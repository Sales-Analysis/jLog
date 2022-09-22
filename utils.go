package jlog

// Check character in end of line.
func charEndOfLine(row string, char string) bool {
	return row[len(row)-1:] == char
}
