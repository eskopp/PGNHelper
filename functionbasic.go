package main

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"strings"
)

func createfile() {
	// Prüfe die Anzahl der übergebenen Argumente
	if len(os.Args) != 3 {
		color.Red("Error: Incorrect number of arguments")
		os.Exit(1)
	}

	// Extrahiere den Dateinamen aus den Argumenten
	inputfile := os.Args[2]

	// Überprüfe, ob die Datei bereits existiert
	if _, err := os.Stat(inputfile); err == nil {
		color.Red(inputfile + " already exists")
		os.Exit(1)
	}

	// Überprüfe, ob die Dateiendung eine von ".pgn" oder ".json" ist
	if !strings.HasSuffix(inputfile, ".pgn") && !strings.HasSuffix(inputfile, ".json") {
		color.Red("Error: File extension must be '.pgn' or '.json'")
		os.Exit(1)
	}

	// Erstelle die Datei
	file, err := os.Create(inputfile)
	if err != nil {
		fmt.Println("Error creating file:", err)
		os.Exit(1)
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("Error closing file:", err)
		}
	}()

	fmt.Println(inputfile, "created successfully")
	os.Exit(0)
}

func deletefile() {
	// Prüfe die Anzahl der übergebenen Argumente
	if len(os.Args) != 3 {
		color.Red("Error: Incorrect number of arguments")
		os.Exit(1)
	}

	// Extrahiere den Dateinamen aus den Argumenten
	inputfile := os.Args[2]

	// Überprüfe, ob die Datei existiert
	if _, err := os.Stat(inputfile); os.IsNotExist(err) {
		color.Red(inputfile + " does not exist")
		os.Exit(1)
	}

	// Überprüfe, ob die Dateiendung ".pgn" oder ".json" ist
	if !strings.HasSuffix(inputfile, ".pgn") && !strings.HasSuffix(inputfile, ".json") {
		color.Red("Error: File extension must be '.pgn' or '.json'")
		os.Exit(1)
	}

	// Lösche die Datei
	err := os.Remove(inputfile)
	if err != nil {
		fmt.Println("Error deleting file:", err)
		os.Exit(1)
	}

	os.Exit(0)
}
