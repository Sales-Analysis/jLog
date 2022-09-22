package jlog

// byteRow converts from string to bytes.
// if row is not contains "\n", adds to the end of the line.
func byteRow(row string) []byte {
	if !charEndOfLine(row, "\n") {
		row += "\n"
	}
	return []byte(row)
}

// Check character in end of line.
func charEndOfLine(row string, char string) bool {
	return row[len(row)-1:] == char
}
