package dotenv

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Load(filename string) error {
	fmt.Println(filename)
	err := readFile(filename)
	return err
}

func readFile(filename string) error {
	r, err := os.Open(filename)
	if err != nil {
		return err
	}
	parseFile(r)
	defer r.Close()
	return nil
}

// parse reads an env file from io.Reader, returning a map of keys and values.
func parseFile(r io.Reader) error {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		return err
	}
}
