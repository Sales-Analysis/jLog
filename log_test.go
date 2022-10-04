package jlog

import (
	"bytes"
	"log"
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	location := "./data/test"
	format := "2006-01-02 15:04:05"

	j := Init(location, format)
	// test struct
	if location != j.location {
		t.Errorf("Location is not equal")
	}

	j.Info("This is info")
	j.Warning("This is warning")
	j.Error("This is error")
	/*
		//check stdout info
		infoMessage := "This is info\n"
		logInfo := captureOutput(func() {
			j.Info(infoMessage)
		})
		if logInfo == "" {
			t.Errorf("Info message is not equal")
		}

		//check stdout warning
		warningMessage := "This is warning"
		logWarning := captureOutput(func() {
			j.Warning(warningMessage)
		})

		if logWarning == "" {
			t.Errorf("Warning message is not equal")
		}

		//check stdout error
		errorMessage := "This is error"
		logError := captureOutput(func() {
			j.Error(errorMessage)
		})

		if logError == "" {
			t.Errorf("Error message is not equal")
		}
	*/
}

func TestLog(t *testing.T) {
	location := "./data/test/logger"
	format := "2006-01-02 15:04:05"

	j := Init(location, format)

	message := "This is simple row."
	captureOutput(func() {
		j.stdout("DUMMY", message)
	})
}

func captureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stderr)
	return buf.String()
}
