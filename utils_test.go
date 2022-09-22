package jlog

import (
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
