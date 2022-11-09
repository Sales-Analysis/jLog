package dotenv

import (
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	Load("./.env_test")
	if os.Getenv("FORMAT_FILENAME") != "20060102" {
		t.Errorf("FORMAT_FILENAME is not equal")
	}
	if os.Getenv("SEPARATOR") != "[]" {
		t.Errorf("SEPARATOR is not equal")
	}
}
