package main

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"strings"
)

func createfile() {
	// Create a PGN File
	var inputfile string

	// Prüfe, ob zu viele oder wenige Args gesetzt sind
	if len(os.Args) != 3 {
		color.Red("Error counting args")
		os.Exit(1)
	}

	// Checke Input
	inputfile = os.Args[2]

	// Prüfe ob die Datei existiert
	if _, err := os.Stat(inputfile); os.IsExist(err) {
		color.Red(inputfile + " exist")
		os.Exit(1)
	}

	// Check File
	fmt.Println(inputfile)

	// Test File
	if _, err := os.Stat(inputfile); err == nil {
		// Überprüfen, ob die Dateiendung ".json" oder ".pgn" ist
		if strings.HasSuffix(inputfile, ".json") || strings.HasSuffix(inputfile, ".pgn") {
			fmt.Println(inputfile, "ist entweder eine JSON- oder eine PGN-Datei.")
			_, err := os.Create(inputfile)
			if err != nil {
				return
			}
		}
	}
}

func deletefile() {
	// Create a PGN File
	var inputfile string

	// Prüfe, ob zu viele oder wenige Args gesetzt sind
	if len(os.Args) != 3 {
		color.Red("Error counting args")
		os.Exit(1)
	}

	// Checke Input
	inputfile = os.Args[2]

	// Prüfe ob die Datei existiert
	if _, err := os.Stat(inputfile); os.IsNotExist(err) {
		color.Red(inputfile + " File dosent exist")
		os.Exit(1)
	}

	// Magic
	if _, err := os.Stat(inputfile); err == nil {
		if strings.HasSuffix(inputfile, ".json") || strings.HasSuffix(inputfile, ".pgn") {
			err := os.Remove(inputfile)
			if err != nil {
				return
			}
		} else {
			fmt.Println("Fehler beim Überprüfen von", inputfile+":", err)
		}
	}
}
