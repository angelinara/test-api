package parser

import (
	"bufio"
	"os"
	"strings"
)

func ReadDescription(path string) (string, error) {

	// Read the file and extract the description
	// Open the .sh file, read all lines, grab line 2 and strip the # prefix.
	// Write a test that confirms the description comes back correctly.

	//read the file, return error if
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	// wraps the file and lets you read one line at a time with .Scan():
	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	description := strings.TrimPrefix(lines[1], "# ")

	return description, nil
}
