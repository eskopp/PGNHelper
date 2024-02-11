package main

import (
	"fmt"
)

func printAllGames(games []Game) {
	for _, game := range games {
		fmt.Println("Event:", game.Event)
		fmt.Println("Site:", game.Site)
		fmt.Println("Date:", game.Date)
		fmt.Println("Round:", game.Round)
		fmt.Println("White:", game.White)
		fmt.Println("Black:", game.Black)
		fmt.Println("Result:", game.Result)
		fmt.Println("Opening:", game.Opening)
		fmt.Println("ECO:", game.ECO)
		fmt.Println("Moves:", game.Moves)
		fmt.Println()
	}
}
