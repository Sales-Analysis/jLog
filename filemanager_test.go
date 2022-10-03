package jlog

import (
	"testing"
)

func TestCheckDirsPositive(t *testing.T) {
	path := "./data"
	if !checkDirs(path) {
		t.Errorf("The path is not exis")
	}
}

func TestCheckDirsNegative(t *testing.T) {
	path := "./isNotExitPath"
	if checkDirs(path) {
		t.Errorf("The path is exit")
	}
}

func TestCreateDir(t *testing.T) {
	path := "./data/test/test_folder"
	err := createDir(path, false)
	if err != nil {
		t.Errorf("Сouldn't create a folder")
	}
	if !checkDirs(path) {
		t.Errorf("Floder not exist")
	}
}

func TestCreateDirMany(t *testing.T) {
	path := "./data/test/test_folder_many/test_folder_many_nested"
	createDir(path, true)
	if !checkDirs(path) {
		t.Errorf("Floder not exist")
	}
}
