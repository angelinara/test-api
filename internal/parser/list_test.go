package parser

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListFile(t *testing.T) {
	// create temp dir
	dir, err := os.MkdirTemp("", "requests")
	require.NoError(t, err)
	defer os.RemoveAll(dir)

	// create two .sh files into it
	file1 := "#!/bin/bash\n# Create a user\ncurl -X POST http://localhost/users\n"
	err = os.WriteFile(filepath.Join(dir, "create-user.sh"), []byte(file1), 0644)
	require.NoError(t, err)

	file2 := "#!/bin/bash\n# Get all users\ncurl -X GET http://localhost/users\n"
	err = os.WriteFile(filepath.Join(dir, "get-users.sh"), []byte(file2), 0644)
	require.NoError(t, err)

	// filepath.Join(dir, "get-users.sh") - puts file inside temp

	items, err := ListRequests(dir)
	require.NoError(t, err)

	assert.ElementsMatch(t, []ListItem{
		{Name: "create-user", Description: "Create a user"},
		{Name: "get-users", Description: "Get all users"},
	}, items)

}
func TestListRequests_EmptyDir(t *testing.T) {
	dir, err := os.MkdirTemp("", "requests")
	require.NoError(t, err)
	defer os.RemoveAll(dir)

	items, err := ListRequests(dir)

	require.NoError(t, err)
	assert.Empty(t, items)
}

func TestListRequests_IgnoresNonShFiles(t *testing.T) {
	dir, err := os.MkdirTemp("", "requests")
	require.NoError(t, err)
	defer os.RemoveAll(dir)

	// create sh file
	sh := "#!/bin/bash\n# Get all users\ncurl -X GET http://localhost/users\n"
	err = os.WriteFile(filepath.Join(dir, "get-users.sh"), []byte(sh), 0644)
	require.NoError(t, err)

	// create txt file
	err = os.WriteFile(filepath.Join(dir, "notes.txt"), []byte("ignore me"), 0644)
	require.NoError(t, err)

	items, err := ListRequests(dir)
	require.NoError(t, err)
	assert.Len(t, items, 1)
	assert.Equal(t, "get-users", items[0].Name)

}
