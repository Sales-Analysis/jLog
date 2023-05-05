package jlog

import (
	"bytes"
	"fmt"
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

func TestAddSepLengthOne(t *testing.T) {
	str := []string{"date", "package", "func", "message"}
	s := addSep("|", str...)
	if s != "|date|package|func: message" {
		t.Errorf("String with sep is not qual")
	}
}

func TestAddSepLengthTwo(t *testing.T) {
	str := []string{"date", "package", "func", "message"}
	s := addSep("[]", str...)
	if s != "[date][package][func]: message" {
		t.Errorf("String with sep is not qual")
	}
}

func TestInit(t *testing.T) {
	location := "./data/test/logger"

	j := Init("./.env")
	// test struct
	if location != j.location {
		fmt.Println(location, j.location)
		t.Errorf("Location is not equal")
	}

	j.Info("This is info")
	j.Warning("This is warning")
	j.Error("This is error")
}

func TestLog(t *testing.T) {
	j := Init("")
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

func TestBackupMaxBytesZero(t *testing.T) {
	path := ""
	count := "0"
	maxBytes := 0

	status := backupNew(path, count, maxBytes)
	if status {
		t.Errorf("maxBytes is not 0")
	}
}

func TestBackupSizeLessMaxBytes(t *testing.T) {
	path := "./data/test/test.log"
	count := "0"
	maxBytes := 1024

	status := backupNew(path, count, maxBytes)
	if status {
		t.Errorf("size not less maxBytes")
	}
}

func TestBackupSizeLessMaxBytesZero(t *testing.T) {
	path := "./data/test/test.log"
	count := "0"
	maxBytes := 0

	status := backupNew(path, count, maxBytes)
	if status {
		t.Errorf("size not less maxBytes")
	}
}

func TestBackupMaxBytes(t *testing.T) {
	path := "./data/test/test.log"
	count := "5"
	maxBytes := 40
	createFileTest(path)
	_ = backupNew(path, count, maxBytes)
}

func createFileTest(path string) {
	f, _ := os.Create(path)

	defer f.Close()

	_, _ = f.WriteString("some text\n")
	_, _ = f.WriteString("some text\n")
	_, _ = f.WriteString("some text\n")
	_, _ = f.WriteString("some text\n")

}
