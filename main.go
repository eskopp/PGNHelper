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

	myGame := Game{
		Event: "Weltmeisterschaft",
		Site:  "New York",
		Date:  "2024.01.24",
		Round: "1",
	}

	// Zugriff auf die Felder des Game-Objekts
	fmt.Println("Event:", myGame.Event)
	fmt.Println("Site:", myGame.Site)
	fmt.Println("Date:", myGame.Date)
	fmt.Println("Round:", myGame.Round)

}

func Add(a int, b int) int {
	// Test Workflow
	return a + b
}
