package utils

import (
	"testing"
)

func TestGenerateSalt(t *testing.T) {
	salt, err := GenerateSalt(SaltLength)

	if err != nil {
		t.Fatalf("Error generating salt: %v", err)
	}

	if uint32(len(salt)) != SaltLength {
		t.Errorf("Expected salt length of %d, got %d", SaltLength, len(salt))
	}

	anotherSalt, err := GenerateSalt(SaltLength)
	if err != nil {
		t.Fatalf("Error generating another salt: %v", err)
	}

	if string(salt) == string(anotherSalt) {
		t.Errorf("Generated salts are identical, expected unique salts")
	}
}
