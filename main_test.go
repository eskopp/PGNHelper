package main

import (
	"os"
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(Add(5, 3), Add(2, 1))
	expected := 11

	if result != expected {
		t.Errorf("Add(5, 3) = %d; want %d", result, expected)
	}
}

func TestFileExists(t *testing.T) {
	filePath := "./test/test.pgn"

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Errorf("Die Datei %s existiert nicht.", filePath)
	}
}
