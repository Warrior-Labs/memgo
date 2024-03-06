package server

import (
	"errors"
	"strings"
)

// validateKey checks if the key is valid
func (s *MemgoServer) validateKey(key string) error {
	// Trim the key
	key = strings.TrimSpace(key)

	// Check that the key is not empty
	if key == "" {
		return errors.New("key cannot be empty")
	}

	// Check that the key is not too long
	if len(key) > 255 {
		return errors.New("key cannot be longer than 255 characters")
	}

	// Check that the key is not too short
	if len(key) < 3 {
		return errors.New("key cannot be shorter than 3 characters")
	}

	// Check that the key does not contain spaces
	if strings.Contains(key, " ") {
		return errors.New("key cannot contain spaces")
	}

	// Check that the key does not contain special characters
	if strings.ContainsAny(key, "!@#$%^&*(){}|<>?") {
		return errors.New("key cannot contain special characters")
	}

	// Return nil if the key is valid
	return nil
}
