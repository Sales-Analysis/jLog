package filemanager

import (
	"os"
	"testing"
)

func TestCreateDir(t *testing.T) {
	path := "./data/test/test_folder"
	err := CreateDir(path, false)
	if err != nil {
		t.Errorf("Сouldn't create a folder")
	}
	if !CheckDirs(path) {
		t.Errorf("Floder not exist")
	}
}

func TestCreateDirMany(t *testing.T) {
	path := "./data/test/test_folder_many/test_folder_many_nested"
	CreateDir(path, true)
	if !CheckDirs(path) {
		t.Errorf("Floder not exist")
	}
}

func TestCheckDirsPositive(t *testing.T) {
	path := "./data"
	if !CheckDirs(path) {
		t.Errorf("The path is not exis")
	}
}

func TestCheckDirsNegative(t *testing.T) {
	path := "./isNotExitPath"
	if CheckDirs(path) {
		t.Errorf("The path is exit")
	}
}

func TestGetFileSize(t *testing.T) {
	path := "./data/data_test.txt"
	size, err := GetSizeOfFile(path)
	if err != nil {
		t.Errorf("%s", err)
	}
	if size != 192 {
		t.Errorf("Size of file is not equal")
	}
}

func TestGetToZip(t *testing.T) {
	GetToZip(
		"./data/data_test.txt",
		"./data/jLog",
	)
	if _, err := os.Stat("./data/jLog.zip"); err != nil {
		t.Errorf("zip file is not exist")
	}
}
