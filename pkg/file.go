package pkg

import (
	"bufio"
	"io"
	"os"
)

func WriteStringsToFile(strings []string, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, str := range strings {
		_, err := io.WriteString(writer, str+"\n")
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}

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
