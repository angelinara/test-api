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
