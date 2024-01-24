package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(5, 3)
	expected := 9

	if result != expected {
		t.Errorf("Add(5, 3) = %d; want %d", result, expected)
	}
}
