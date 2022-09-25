package jlog

import (
	"io"
	"os"
)

type Jlog struct {
	location string
}

func Init(location string) *Jlog {
	return &Jlog{
		location: location,
	}
}

func stdout(message string) {
	io.WriteString(os.Stdout, message)
}
