package jlog

import (
	"reflect"
	"testing"
)

func TestByteRow(t *testing.T) {
	row := byteRow("test message")
	if reflect.TypeOf(row).Elem().Kind() != reflect.Uint8 {
		t.Errorf("The row type is not unit8")
	}
}

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
