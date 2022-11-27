package jlog

import (
	"fmt"
	"testing"
)

func TestCharEndOfLinePositiveCase(t *testing.T) {
	if !charEndOfLine("message\n", "\n") {
		t.Errorf("character not found end of line")
	}
}

func TestCharEndOfLineNegativeCase(t *testing.T) {
	if charEndOfLine("message", "\n") {
		t.Errorf("character in end of line")
	}
}

func TestMakeFilename(t *testing.T) {
	format := "20060102"
	filename := makeFilename("20060102")
	testFilename := fmt.Sprintf("%s.log", timeNow(format))
	if filename != testFilename {
		t.Errorf("file name is not equal test value")
	}
}