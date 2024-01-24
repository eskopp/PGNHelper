package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(Add(5, 3),Add(2,1))
	expected := 11

	if result != expected {
		t.Errorf("Add(5, 3) = %d; want %d", result, expected)
	}
}
