package auth

import "testing"

func TestHashPassword(t *testing.T) {
	password := "mysecretpassword"
	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(hash) == 0 {
		t.Fatalf("Expected hash to be a non-empty string")
	}

	if hash == password {
		t.Fatalf("Expected hash to be different from the original password")
	}
}

func TestCheckPasswordHash(t *testing.T) {
	password := "mysecretpassword"
	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !CheckPasswordHash(password, hash) {
		t.Fatalf("Expected CheckPasswordHash to return true")
	}

	if CheckPasswordHash("wrongpassword", hash) {
		t.Fatalf("Expected CheckPasswordHash to return false for wrong password")
	}
}
