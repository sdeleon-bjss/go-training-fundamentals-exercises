package pkg

import (
	"bufio"
	"io"
	"os"
)

// WriteStringsToFile writes a slice of strings to a file
//
// This will create a new file at root of project
func WriteStringsToFile(strings []string, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, str := range strings {
		_, err = io.WriteString(writer, str+"\n")
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}

// ReadStringsFromFile reads a slice of strings from a file
//
// Expects a file at root of project
func ReadStringsFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var strings []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		strings = append(strings, scanner.Text())
	}

	return strings, scanner.Err()
}
