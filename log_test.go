package jlog

import (
	"testing"
)

func TestInit(t *testing.T) {
	location := "./data/test"

	j := Init(location)
	// test struct
	if location != j.location {
		t.Errorf("Location is not equal")
	}

	//stdout
	j.Info("This is info")
	j.Warning("This is warning")
	j.Error("This is error")
}

func TestLog(t *testing.T) {
	stdout("Это строка для стандартного вывода.\n")
}
