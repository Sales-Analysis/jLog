package jlog

import (
	"bytes"
	"log"
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	location := "./data/test/logger"

	j := Init(location, "./.env")
	// test struct
	if location != j.location {
		t.Errorf("Location is not equal")
	}

	j.Info("This is info")
	j.Warning("This is warning")
	j.Error("This is error")
}

func TestLog(t *testing.T) {
	location := "./data/test/logger"

	j := Init(location, "./.env")

	message := "This is simple row."
	j.Dummy(message)
}

func captureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stderr)
	return buf.String()
}
