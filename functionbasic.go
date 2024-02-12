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

	// Erstelle die Datei entsprechend der Dateiendung
	var filename string
	if strings.HasSuffix(inputfile, ".pgn") || strings.HasSuffix(inputfile, ".json") {
		filename = inputfile
	} else {
		filename = inputfile + ".pgn"
	}

	// Erstelle die Datei
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		os.Exit(1)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	fmt.Println(filename, "created successfully")
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

	// Überprüfe, ob die Datei bereits existiert
	if _, err := os.Stat(inputfile); err == nil {
		color.Red(inputfile + " already exists")
		os.Exit(1)
	}

	// Erstelle die Datei entsprechend der Dateiendung
	var filename string
	if strings.HasSuffix(inputfile, ".pgn") || strings.HasSuffix(inputfile, ".json") {
		filename = inputfile
	} else {
		filename = inputfile + ".pgn"
	}

	// Erstelle die Datei
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		os.Exit(1)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	fmt.Println(filename, "created successfully")
	os.Exit(0)
}
