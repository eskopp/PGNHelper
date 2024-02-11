package main

import (
	"fmt"
	"github.com/fatih/color"
)

func main() {
	fmt.Println("Hello, World!")
	color.Red("Dies ist ein roter Text")
	color.Green("Dies ist ein grüner Text")
	color.Blue("Dies ist ein blauer Text")

	pgnFile := "./test/test.pgn"

	games, err := parsePGNFile(pgnFile)
	if err != nil {
		fmt.Println("Fehler beim Parsen der PGN-Datei:", err)
		return
	}

	for _, game := range games {
		fmt.Printf("Event: %s\nSite: %s\nDate: %s\nRound: %s\nWhite: %s\nBlack: %s\nResult: %s\nOpening: %s\nECO: %s\nMoves: %s\n\n",
			game.Event, game.Site, game.Date, game.Round, game.White, game.Black, game.Result, game.Opening, game.ECO, game.Moves)
	}
}

func Add(a int, b int) int {
	// Test Workflow
	return a + b
}
