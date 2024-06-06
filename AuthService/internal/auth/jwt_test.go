package auth

import (
	"github.com/dgrijalva/jwt-go"
	"testing"
	"time"
)

func TestGenerateToken(t *testing.T) {
	email := "test@example.com"
	token, err := GenerateToken(email)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if token == "" {
		t.Fatalf("Expected a token, got an empty string")
	}

	parsedEmail, err := ParseToken(token)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if parsedEmail != email {
		t.Fatalf("Expected email %v, got %v", email, parsedEmail)
	}
}

// TestParseToken tests the ParseToken function with a valid token
func TestParseToken(t *testing.T) {
	email := "test@example.com"
	token, err := GenerateToken(email)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	parsedEmail, err := ParseToken(token)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if parsedEmail != email {
		t.Fatalf("Expected email %v, got %v", email, parsedEmail)
	}
}

// TestParseTokenInvalid tests the ParseToken function with an invalid token
func TestParseTokenInvalid(t *testing.T) {
	email, err := ParseToken("fefe.fewf.efew")
	if email != "" && err == nil {
		t.Fatalf("Expected an error, got nil")
	}
}

func TestParseTokenValid(t *testing.T) {
	email, err := ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFtaW5hQGV4YW1wbGUuY29tIiwiZXhwIjoxNzE3NzIxMDc3fQ.WABgOupHHJdcp3h_ZKgiG4Noqv2ICe4PA59y-enk-OY")
	if email == "" || err != nil {
		t.Fatalf("Expected an error, got nil")
	}
}

// TestExpiredToken tests the ParseToken function with an expired token
func TestExpiredToken(t *testing.T) {
	email := "test@example.com"
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = time.Now().Add(-time.Hour).Unix() // Set to one hour ago

	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	_, err = ParseToken(tokenString)
	if err == nil {
		t.Fatalf("Expected an error due to expired token, got nil")
	}
}
