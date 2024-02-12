package main

import (
	"os"
	"testing"
)

func TestRemovecb(t *testing.T) {
	// Erstellen von Dummy-Dateien mit den Erweiterungen .pgi und .ini
	files := []string{"test1.pgi", "test2.ini", "test3.txt"}

	for _, file := range files {
		_, err := os.Create(file)
		if err != nil {
			t.Fatalf("Fehler beim Erstellen der Testdatei %s: %s", file, err)
		}
		defer func(name string) {
			err := os.Remove(name)
			if err != nil {

			}
		}(file)
	}

	// Testen der Funktion removecb
	removecb()

	// Überprüfen, ob die Dateien mit den Erweiterungen .pgi und .ini gelöscht wurden
	for _, file := range files[:2] {
		if _, err := os.Stat(file); err == nil {
			t.Errorf("Datei %s existiert immer noch, obwohl sie gelöscht werden sollte", file)
		}
	}

	// Überprüfen, ob die Datei mit der Erweiterung .txt noch existiert
	if _, err := os.Stat(files[2]); os.IsNotExist(err) {
		t.Errorf("Datei %s wurde unerwartet gelöscht", files[2])
	}
}
