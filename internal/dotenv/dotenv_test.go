package dotenv

import (
	"os"
	"testing"
)

func TestInitWithoutExceptions(t *testing.T) {
	var exceptions []string
	Load("./.env_test", exceptions)
	if os.Getenv("FORMAT_FILENAME") != "20060102" {
		t.Errorf("FORMAT_FILENAME is not equal")
	}
	if os.Getenv("SEPARATOR") != "[]" {
		t.Errorf("SEPARATOR is not equal")
	}
}
