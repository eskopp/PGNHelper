package main

import (
	"encoding/json"
	"reflect"
	"testing"
)

// TestCreateGame tests the creation of a Game struct.
func TestCreateGame(t *testing.T) {
	tags := map[string]string{
		"Event": "Test Event",
		"Site":  "Test Site",
	}
	moves := "1.e4 e5 2.Nf3 Nc6"

	game := Game{
		Tags:  tags,
		Moves: moves,
	}

	if !reflect.DeepEqual(game.Tags, tags) {
		t.Errorf("Expected Tags %v, got %v", tags, game.Tags)
	}

	if game.Moves != moves {
		t.Errorf("Expected Moves %v, got %v", moves, game.Moves)
	}
}

// TestGameSerialization tests the JSON serialization of a Game struct.
func TestGameSerialization(t *testing.T) {
	game := Game{
		Tags: map[string]string{
			"Event": "Test Event",
			"Site":  "Test Site",
		},
		Moves: "1.e4 e5 2.Nf3 Nc6",
	}

	expectedJSON := `{"tags":{"Event":"Test Event","Site":"Test Site"},"moves":"1.e4 e5 2.Nf3 Nc6"}`

	data, err := json.Marshal(game)
	if err != nil {
		t.Fatalf("Failed to serialize Game: %v", err)
	}

	if string(data) != expectedJSON {
		t.Errorf("Expected JSON %v, got %v", expectedJSON, string(data))
	}
}
