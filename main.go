package main

import (
	"fmt"
	"github.com/fatih/color"
)

func main() {
	color.Blue("PGN Helper")

	pgnFile := "./test/test.pgn"

	games, err := parsePGNFile(pgnFile)
	if err != nil {
		fmt.Println("Fehler beim Parsen der PGN-Datei:", err)
		return
	}

	// TODO: better output
	printAllGames(games)
}

func Add(a int, b int) int {
	return a + b
}
