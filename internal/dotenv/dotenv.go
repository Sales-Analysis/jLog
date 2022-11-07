package dotenv

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func Load(filename string) error {
	envMap, err := readFile(filename)

	if err != nil {
		return err
	}
	for key, value := range envMap {
		os.Setenv(key, value)
	}
	return nil
}

func readFile(filename string) (map[string]string, error) {
	r, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	lines, err := parseFile(r)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	return lines, nil
}

// parse reads an env file from io.Reader, returning a map of keys and values.
func parseFile(r io.Reader) (map[string]string, error) {
	values := make(map[string]string)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		str := scanner.Text()
		if strings.Contains(str, "#") {
			continue
		}
		splitStr := strings.Split(str, "=")
		key := strings.ReplaceAll(splitStr[0], " ", "")
		value := strings.ReplaceAll(splitStr[1], " ", "")
		values[key] = value
	}
	err := scanner.Err()
	if err != nil {
		return nil, err
	}
	return values, err
}
