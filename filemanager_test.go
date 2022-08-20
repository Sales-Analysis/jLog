package jlog

import (
	"os"
	"testing"
)

func TestWriteFile(t *testing.T) {
	r := []byte("Some comment!")
	err := writeFile("./data/test/test_log.log", r, 0644)
	if err != nil {
		t.Errorf(err.Error())
	}
	_, errf := os.Stat("./data/test/test_log.log")
	if errf != nil {
		t.Errorf(errf.Error())
	}
}
