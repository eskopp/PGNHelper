package main

import (
	"testing"
)

// Teste die Erstellung eines Game-Structs
func TestGameCreation(t *testing.T) {
	testEvent := "Weltmeisterschaft"
	testSite := "New York"
	testDate := "2024.01.24"
	testRound := "1"

	game := Game{
		Event: testEvent,
		Site:  testSite,
		Date:  testDate,
		Round: testRound,
	}

	// Überprüfe, ob die Felder korrekt gesetzt wurden
	if game.Event != testEvent {
		t.Errorf("Erwartetes Event %s, aber erhalten %s", testEvent, game.Event)
	}
	if game.Site != testSite {
		t.Errorf("Erwartete Site %s, aber erhalten %s", testSite, game.Site)
	}
	if game.Date != testDate {
		t.Errorf("Erwartetes Date %s, aber erhalten %s", testDate, game.Date)
	}
	if game.Round != testRound {
		t.Errorf("Erwartete Round %s, aber erhalten %s", testRound, game.Round)
	}
}
