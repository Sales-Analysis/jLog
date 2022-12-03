package jlog

import (
	"os"
	"testing"
)

func TestSetDefaultParams(t *testing.T) {
	setDefaultParams()
	loc := os.Getenv("LOCATION")
	if loc != "logger" {
		t.Errorf("value of var LOCATION is not equal")
	}
}
