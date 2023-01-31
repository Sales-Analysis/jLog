package jlog

import (
	"fmt"
	"os"

	"github.com/Sales-Analysis/jLog/internal/dotenv"
)

// Default parameters
var defaultParameters = map[string]string{
	// Folder with log files
	"LOCATION": "logger",
	// Format log file name
	"FORMAT_FILENAME": "20060102",
	// Size of file
	"MAX_BYTES": "0",
	// Format log
	"FORMAT_TIME_LOG": "2006-01-02 15:04:05",
	// Separator
	"SEPARATOR": "[]",
	// log to stdout
	"GOTOSTD": "true",
	// log to file
	"GOTOFILE": "true",
}

func loadDotEnv(path string) {
	if len(path) != 0 {
		err := dotenv.Load(path)
		if err != nil {
			fmt.Println(err)
		}
	}
	setDefaultParams()
}

// Set default parameters.
func setDefaultParams() {
	for key, value := range defaultParameters {
		_, found := os.LookupEnv(key)
		if !found {
			fmt.Printf("Parameter %s is not sets. The default parameter is used.\n", key)
			os.Setenv(key, value)
		}
	}
}
