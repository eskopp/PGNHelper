package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// ProcessGames liest die JSON-Datei, parst die Spiele und schreibt sie als PGN in die angegebene Datei.
func parseJSONFile(jsonPath, pgnPath string) error {
	// JSON-Daten lesen mit os.ReadFile
	jsonData, err := os.ReadFile(jsonPath)
	if err != nil {
		return fmt.Errorf("error reading JSON file: %w", err)
	}

	// JSON-Daten parsen
	var games []Game
	if err := json.Unmarshal(jsonData, &games); err != nil {
		return fmt.Errorf("error parsing JSON data: %w", err)
	}

	// Spiele in PGN konvertieren
	pgnContent := convertGamesToPGN(games)

	// PGN-Daten in Datei schreiben mit os.WriteFile
	if err := os.WriteFile(pgnPath, []byte(pgnContent), 0644); err != nil {
		return fmt.Errorf("error writing PGN file: %w", err)
	}

	return nil // Kein Fehler
}

// convertGamesToPGN formatiert und konvertiert Spiel-Daten von JSON zu PGN.
func convertGamesToPGN(games []Game) string {
	var pgnBuilder strings.Builder

	for _, game := range games {
		for tag, value := range game.Tags {
			pgnBuilder.WriteString(fmt.Sprintf("[%s \"%s\"]\n", tag, value))
		}
		pgnBuilder.WriteString("\n")
		pgnBuilder.WriteString(game.Moves)
		pgnBuilder.WriteString("\n\n")
	}

	return pgnBuilder.String()
}
