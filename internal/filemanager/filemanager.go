package filemanager

import (
	"os"
)

// Write writes data to a file named by filename.
func Write(filename string, message string) {
	writeToFile(filename, []byte(message), 0644)
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

// Check that the path exist
// if path does not exist, return false
func CheckDirs(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// Creates missing folders
func CreateDir(path string, many bool) error {
	if !CheckDirs(path) {
		switch many {
		case true:
			return os.MkdirAll(path, 0777)
		default:
			return os.Mkdir(path, 0777)
		}
	}
	return nil
}
