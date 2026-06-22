package parser

import (
	"bufio"
	"os"
	"strings"
)

// open file, returns description from line 2
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

type Request struct {
	Method      string
	URL         string
	Headers     []string
	Body        string
	Description string
}

// walks tokens and returns a Request struct
func extractFields(tokens []string) *Request {
	r := &Request{}
	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case "-X":
			// next token is the method
			if i+1 < len(tokens) {
				r.Method = tokens[i+1]
				i++
			}
		case "-H":
			if i+1 < len(tokens) {
				r.Headers = append(r.Headers, tokens[i+1])
				i++
			}
		case "-d", "--data":
			if i+1 < len(tokens) {
				r.Body = tokens[i+1]
				i++
			}
		default:
			// if it starts with "http", it's the URL
			if strings.HasPrefix(tokens[i], "http") {
				r.URL = tokens[i]
				i++
			}
		}
	}
	return r
}

// splits a curl command string into tokens
func tokenise(s string) []string {
	var current strings.Builder
	inQuote := false
	quoteChar := byte(0)

	var tokens []string

	for i := 0; i < len(s); i++ {
		c := s[i]
		if inQuote && c == quoteChar {
			inQuote = false
		} else if !inQuote && (c == '\'' || c == '"') {
			inQuote = true
			quoteChar = c
		} else if !inQuote && (c == ' ' || c == '\t' || c == '\\' || c == '\n') {
			if current.Len() > 0 {
				tokens = append(tokens, current.String())
				current.Reset()
			}
		} else {
			current.WriteByte(c)
		}
	}

	if current.Len() > 0 {
		tokens = append(tokens, current.String())
	}

	return tokens
}

func ParseFile(path string) (*Request, error) {
	//open file
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// read file
	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	content := strings.Join(lines[2:], " ")
	tokens := tokenise(content)
	r := extractFields(tokens)
	r.Description = strings.TrimPrefix(lines[1], "# ")

	return r, nil
}
