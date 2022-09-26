package jlog

import (
	"io"
	"os"
)

type jlog struct {
	location string // Folder with log files
}

// Create new jLog.
// The location variable sets the folder with log files.
func Init(location string) *jlog {
	return &jlog{
		location: location,
	}
}

func (j *jlog) Info(message string) {
	stdout(message)
}

func (j *jlog) Warning(message string) {
	stdout(message)
}

func (j *jlog) Error(message string) {
	stdout(message)
}

func stdout(message string) {
	io.WriteString(os.Stdout, message)
}
