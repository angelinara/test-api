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
		description, err := ReadDescription(file)
		if err != nil {
			return nil, err
		}
		name := strings.TrimSuffix(filepath.Base(file), ".sh")
		list = append(list, ListItem{Name: name, Description: description})
	}
	return list, nil
}

type ListItem struct {
	Name        string
	Description string
}
