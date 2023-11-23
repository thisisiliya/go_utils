package file

import (
	"bufio"
	"os"
)

func ReadByLine(path string) (*[]string, error) {

	var contents []string
	file, err := os.Open(path)

	if err != nil {

		return &contents, err
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		if text := scanner.Text(); text != "" {

			contents = append(contents, scanner.Text())
		}
	}

	if scanner.Err() != nil {

		return &contents, err
	}

	file.Close()
	return &contents, nil
}
