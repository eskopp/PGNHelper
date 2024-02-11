package main

import (
	"encoding/json"
	"gopkg.in/freeeve/pgn.v1"
	"os"
)

// GameData hält die Daten einer Schachpartie für die JSON-Konvertierung.
type GameData struct {
	Tags  map[string]string `json:"tags"`
	Moves []string          `json:"moves"`
}

// parsePGNFileToJSON liest eine PGN-Datei und konvertiert sie in eine JSON-Datei.
func parsePGNFileToJSON(pgnFilePath, jsonFilePath string) error {
	f, err := os.Open(pgnFilePath)
	if err != nil {
		return err
	}
	defer f.Close()

	ps := pgn.NewPGNScanner(f)
	var games []GameData

	for ps.Next() {
		game, err := ps.Scan()
		if err != nil {
			return err
		}

		gameData := GameData{
			Tags: game.Tags,
			Moves: func(moves []pgn.Move) []string {
				var ms []string
				for _, m := range moves {
					ms = append(ms, m.String())
				}
				return ms
			}(game.Moves),
		}

		games = append(games, gameData)
	}

	jsonData, err := json.MarshalIndent(games, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(jsonFilePath, jsonData, 0644); err != nil {
		return err
	}

	return nil
}
