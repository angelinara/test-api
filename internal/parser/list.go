package parser

import (
	"os"
	"path/filepath"
	"strings"
)

func ListRequests(dir string) ([]ListItem, error) {
	//get all files in the directory
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var shFiles []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		// filter for .sh files only
		if filepath.Ext(entry.Name()) == ".sh" {
			shFiles = append(shFiles, filepath.Join(dir, entry.Name()))
		}
	}

	var list []ListItem
	// get name and description to add to list
	for _, file := range shFiles {
		//read everything, ParseFile gives everything — method, URL, headers, body, description.
		r, err := ParseFile(file)
		if err != nil {
			return nil, err
		}
		name := strings.TrimSuffix(filepath.Base(file), ".sh")
		list = append(list, ListItem{Name: name, Description: r.Description, Method: r.Method, URL: r.URL})
	}
	return list, nil
}

type ListItem struct {
	Name        string
	Description string
	Method      string
	URL         string
}
