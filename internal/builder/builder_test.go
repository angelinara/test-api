package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// pass strings into the function being tested, 
// simulating command-line arguments a user would type
func TestParseFlags_ValidGET(t *testing.T) {
	r, err := ParseFlags([]string{
		"--name", "get-users",
		"--description", "Fetch all users",
		"--method", "GET",
		"--url", "http://localhost:3000/users",
	})
	require.NoError(t, err)
	// check that the Name and Method fields on the returned struct were parsed correctly
	assert.Equal(t, "get-users", r.Name)
	assert.Equal(t, "GET", r.Method)
	assert.Empty(t, r.Body)
	assert.Empty(t, r.Headers)
}

func TestParseFlags_ValidPOSTWithHeadersAndBody(t *testing.T) {
	r, err := ParseFlags([]string{
		"--name", "login",
		"--description", "Login with email and password",
		"--method", "POST",
		"--url", "http://localhost:3000/auth/login",
		"--header", "Content-Type: application/json",
		"--header", "Authorization: Bearer abc123",
		"--body", `{"email": "test@example.com", "password": "secret"}`,
	})
	require.NoError(t, err)
	assert.Len(t, r.Headers, 2)
	assert.NotEmpty(t, r.Body)
}

func TestParseFlags_MissingName(t *testing.T) {
	_, err := ParseFlags([]string{
		"--description", "Fetch all users",
		"--method", "GET",
		"--url", "http://localhost:3000/users",
	})
	assert.ErrorContains(t, err, "--name is required")
}

func TestParseFlags_MissingDescription(t *testing.T) {
	_, err := ParseFlags([]string{
		"--name", "get-users",
		"--method", "GET",
		"--url", "http://localhost:3000/users",
	})
	assert.ErrorContains(t, err, "--description is required")
}

func TestParseFlags_DescriptionTooLong(t *testing.T) {
	_, err := ParseFlags([]string{
		"--name", "get-users",
		"--description", "This description is way too long and exceeds the fifty character limit",
		"--method", "GET",
		"--url", "http://localhost:3000/users",
	})
	assert.ErrorContains(t, err, "--description must be 50 characters or fewer")
}

func TestParseFlags_MissingMethod(t *testing.T) {
	_, err := ParseFlags([]string{
		"--name", "get-users",
		"--description", "Fetch all users",
		"--url", "http://localhost:3000/users",
	})
	assert.ErrorContains(t, err, "--method is required")
}

func TestParseFlags_MissingURL(t *testing.T) {
	_, err := ParseFlags([]string{
		"--name", "get-users",
		"--description", "Fetch all users",
		"--method", "GET",
	})
	assert.ErrorContains(t, err, "--url is required")
}
