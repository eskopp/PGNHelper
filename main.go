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

}

func Add(a int, b int) int {
	// Test Workflow
	return a + b
}
