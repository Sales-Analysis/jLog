package jlog

import (
	"bytes"
	"log"
	"os"
	"testing"
)


func TestSepLengthOne(t *testing.T) {
	s := sepStr("test", "|")
	if s != "|test" {
		t.Errorf("String with sep is not qual")
	}
}

func TestSepLengthTwo(t *testing.T) {
	s := sepStr("test", "[]")
	if s != "[test]" {
		t.Errorf("String with sep is not qual")
	}
}

func TestInit(t *testing.T) {
	location := "./data/test/logger"

	j := Init("./.env")
	// test struct
	if location != j.location {
		t.Errorf("Location is not equal")
	}

	j.Info("This is info")
	j.Warning("This is warning")
	j.Error("This is error")
}

func TestLog(t *testing.T) {
	j := Init("./.env")

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
