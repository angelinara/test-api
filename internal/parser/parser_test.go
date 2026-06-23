package parser

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadDescription(t *testing.T) {
	content := "#!/bin/bash\n# Login with email and password\ncurl -s -X POST \"http://localhost\"\n"

	tmp, err := os.CreateTemp("", "*.sh")
	require.NoError(t, err)
	defer os.Remove(tmp.Name())

	_, err = tmp.WriteString(content)
	require.NoError(t, err)
	tmp.Close()

	description, err := ReadDescription(tmp.Name())
	require.NoError(t, err)
	assert.Equal(t, "Login with email and password", description)
}

func TestTokenise_Simple(t *testing.T) {
	tokens := tokenise(`curl -s -X POST "http://localhost/login"`)
	assert.Equal(t, []string{"curl", "-s", "-X", "POST", "http://localhost/login"}, tokens)
}

func TestTokenise_QuotedBody(t *testing.T) {
	tokens := tokenise(`curl -d '{"email": "a@b.com"}' "http://localhost"`)
	assert.Equal(t, []string{"curl", "-d", `{"email": "a@b.com"}`, "http://localhost"}, tokens)
}

func TestExtractFields_GET(t *testing.T) {
	tokens := []string{"curl", "-s", "-X", "GET", "http://localhost:3000/users"}
	r := extractFields(tokens)
	assert.Equal(t, "GET", r.Method)
	assert.Equal(t, "http://localhost:3000/users", r.URL)
	assert.Empty(t, r.Body)
	assert.Empty(t, r.Headers)
}

func TestExtractFields_POSTWithHeadersAndBody(t *testing.T) {
	tokens := []string{"curl", "-s", "-X", "POST", "-H", "Content-Type: application/json", "-d", `{"email": "a@b.com"}`, "http://localhost:3000/login"}
	r := extractFields(tokens)
	assert.Equal(t, "POST", r.Method)
	assert.Equal(t, "http://localhost:3000/login", r.URL)
	assert.Equal(t, []string{"Content-Type: application/json"}, r.Headers)
	assert.Equal(t, `{"email": "a@b.com"}`, r.Body)
}

func TestParseFile_GET(t *testing.T) {
	content := "#!/bin/bash\n# Get all users\ncurl -s \\\n  -X GET \\\n  \"http://localhost:3000/users\"\n"

	tmp, err := os.CreateTemp("", "*.sh")
	require.NoError(t, err)
	defer os.Remove(tmp.Name())

	_, err = tmp.WriteString(content)
	require.NoError(t, err)
	tmp.Close()

	r, err := ParseFile(tmp.Name())
	require.NoError(t, err)
	assert.Equal(t, "Get all users", r.Description)
	assert.Equal(t, "GET", r.Method)
	assert.Equal(t, "http://localhost:3000/users", r.URL)
	assert.Empty(t, r.Headers)
	assert.Empty(t, r.Body)
}

func TestParseFile_MultiLineBody(t *testing.T) {
	content := "#!/bin/bash\n# Create user\ncurl -s \\\n  -X POST \\\n  -H \"Content-Type: application/json\" \\\n  -d '{\n    \"email\": \"a@b.com\",\n    \"name\": \"Alice\"\n  }' \\\n  \"http://localhost:3000/users\"\n"

	tmp, err := os.CreateTemp("", "*.sh")
	require.NoError(t, err)
	defer os.Remove(tmp.Name())

	_, err = tmp.WriteString(content)
	require.NoError(t, err)
	tmp.Close()

	r, err := ParseFile(tmp.Name())
	require.NoError(t, err)
	assert.Equal(t, "Create user", r.Description)
	assert.Equal(t, "POST", r.Method)
	assert.Equal(t, "http://localhost:3000/users", r.URL)
	assert.Equal(t, []string{"Content-Type: application/json"}, r.Headers)
	assert.Contains(t, r.Body, `"email": "a@b.com"`)
	assert.Contains(t, r.Body, `"name": "Alice"`)
}

func TestParseFile(t *testing.T) {
	content := "#!/bin/bash\n# Login with email and password\ncurl -s \\\n  -X POST \\\n  -H \"Content-Type: application/json\" \\\n  -d '{\"email\": \"a@b.com\"}' \\\n  \"http://localhost:3000/login\"\n"

	tmp, err := os.CreateTemp("", "*.sh")
	require.NoError(t, err)
	defer os.Remove(tmp.Name())

	_, err = tmp.WriteString(content)
	require.NoError(t, err)
	tmp.Close()

	r, err := ParseFile(tmp.Name())
	require.NoError(t, err)
	assert.Equal(t, "Login with email and password", r.Description)
	assert.Equal(t, "POST", r.Method)
	assert.Equal(t, "http://localhost:3000/login", r.URL)
	assert.Equal(t, []string{"Content-Type: application/json"}, r.Headers)
	assert.Equal(t, `{"email": "a@b.com"}`, r.Body)
}
