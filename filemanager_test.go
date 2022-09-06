package jlog

import "testing"

func TestCheckDirsPositive(t *testing.T) {
	path := "./data"
	if !checkDirs(path) {
		t.Errorf("The path is not exis")
	}
}

/*
func TestCheckDirsNegative(t *testing.T) {
	path := "./isNotExitPath"
	if checkDirs(path) {
		t.Errorf("The path is exit")
	}
}

func TestCreateDir(t *testing.T) {
	path := "./data/test/test_folder"
	createDir(path, false)
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
*/
