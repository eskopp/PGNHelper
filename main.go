package main

import (
	"fmt"
	"github.com/fatih/color"
	"log"
)

func main() {
	color.Blue("PGN Helper")

	pgnFile, jsonFile := "./test/test.pgn", "./test/games.json"

	if err := parsePGNFileToJSON(pgnFile, jsonFile); err != nil {
		log.Fatalf("Fehler beim Konvertieren der PGN-Datei in JSON: %v", err)
	}

	fmt.Println("Die PGN-Datei wurde erfolgreich in JSON konvertiert.")
}

func Add(a int, b int) int {
	return a + b
}
