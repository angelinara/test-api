package builder

import (
	"flag"
	"fmt"
	"os"
)

type Request struct {
	Name        string
	Description string
	Method      string
	URL         string
	Headers     []string
	Body        string
}

type multiFlag []string

func (m *multiFlag) String() string     { return "" }
func (m *multiFlag) Set(v string) error { *m = append(*m, v); return nil }

// builder.go sets up how tapi new reads its inputs from the command line.
// ParseFlags reads those --name, --description etc. values, checks none are missing or invalid,
// and packages them into a Request struct that the rest of the code can use.
// multiFlag trick to allow --header to be repeated multiple times

// read flags, validate, return the data.
func ParseFlags(args []string) (*Request, error) {
	fs := flag.NewFlagSet("new", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)

	// registers the flags and expected response
	var headers multiFlag
	name := fs.String("name", "", "request name (used as filename)")
	description := fs.String("description", "", "short description, max 50 chars")
	method := fs.String("method", "", "HTTP method: GET, POST, PUT, PATCH, DELETE")
	url := fs.String("url", "", "request URL")
	body := fs.String("body", "", "request body")
	fs.Var(&headers, "header", "header in 'Key: Value' format, repeatable")

	// read the flags
	if err := fs.Parse(args); err != nil {
		return nil, err
	}

	// validation
	if *name == "" {
		return nil, fmt.Errorf("--name is required")
	}
	if *description == "" {
		return nil, fmt.Errorf("--description is required")
	}
	if len(*description) > 50 {
		return nil, fmt.Errorf("--description must be 50 characters or fewer")
	}
	if *method == "" {
		return nil, fmt.Errorf("--method is required")
	}
	if *url == "" {
		return nil, fmt.Errorf("--url is required")
	}

	// Request struct to return data
	return &Request{
		Name:        *name,
		Description: *description,
		Method:      *method,
		URL:         *url,
		Headers:     headers,
		Body:        *body,
	}, nil
}
