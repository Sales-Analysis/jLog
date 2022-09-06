package jlog

import (
	"fmt"
	"os"
	"strings"
)

// Check that the path exist
// if path does not exist, return false
func checkDirs(path string) bool {
	_, err := os.Stat(path)
	fmt.Println(os.IsNotExist(err))
	return !os.IsNotExist(err)
}

// Creates missing folders
func createDir(path string, many bool) {
	if !checkDirs(path) {
		if many {
			_ = os.MkdirAll(path, 0777)
		} else {
			_ = os.Mkdir(path, 0777)
		}
	}
}

// writeFile writes data to a file named by filename.
// If the file does not exist, the file will be created.
// If the file exist, data will be appended to file
func writeToFile(filename string, data []byte, perm os.FileMode) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, perm)
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	if errW := f.Close(); err == nil {
		err = errW
	}
	return err
}

func byteRow(row string) []byte {
	if !strings.Contains(row, "\n") {
		row += "\n"
	}
	return []byte(row)
}

func Write(message string) {
	byteRow(message)
}
