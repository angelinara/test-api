package builder

import (
	"testing"
)

func TestParseFlags_ValidGET(t *testing.T) {
	r, err := ParseFlags([]string{
		"--name", "get-users",
		"--description", "Fetch all users",
		"--method", "GET",
		"--url", "http://localhost:3000/users",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if r.Name != "get-users" {
		t.Errorf("expected name 'get-users', got '%s'", r.Name)
	}
	if r.Body != "" {
		t.Errorf("expected empty body for GET, got '%s'", r.Body)
	}
	if len(r.Headers) != 0 {
		t.Errorf("expected no headers, got %v", r.Headers)
	}
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
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(r.Headers) != 2 {
		t.Errorf("expected 2 headers, got %d", len(r.Headers))
	}
	if r.Body == "" {
		t.Error("expected body to be set")
	}
}

func TestParseFlags_MissingName(t *testing.T) {
	_, err := ParseFlags([]string{
		"--description", "Fetch all users",
		"--method", "GET",
		"--url", "http://localhost:3000/users",
	})
	if err == nil {
		t.Error("expected error for missing --name")
	}
}

func TestParseFlags_MissingDescription(t *testing.T) {
	_, err := ParseFlags([]string{
		"--name", "get-users",
		"--method", "GET",
		"--url", "http://localhost:3000/users",
	})
	if err == nil {
		t.Error("expected error for missing --description")
	}
}

func TestParseFlags_DescriptionTooLong(t *testing.T) {
	_, err := ParseFlags([]string{
		"--name", "get-users",
		"--description", "This description is way too long and exceeds the fifty character limit",
		"--method", "GET",
		"--url", "http://localhost:3000/users",
	})
	if err == nil {
		t.Error("expected error for description over 50 chars")
	}
}

func TestParseFlags_MissingMethod(t *testing.T) {
	_, err := ParseFlags([]string{
		"--name", "get-users",
		"--description", "Fetch all users",
		"--url", "http://localhost:3000/users",
	})
	if err == nil {
		t.Error("expected error for missing --method")
	}
}

func TestParseFlags_MissingURL(t *testing.T) {
	_, err := ParseFlags([]string{
		"--name", "get-users",
		"--description", "Fetch all users",
		"--method", "GET",
	})
	if err == nil {
		t.Error("expected error for missing --url")
	}
}
