package main

import (
	"github.com/google/uuid"
	"strings"
)

// Generate a UUID
func generateUUID() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}
