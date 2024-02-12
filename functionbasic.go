package main

import (
	"github.com/fatih/color"
	"os"
	"strings"
)

// PGN Block
func createpgn() {
	// Create a PGN File
	var inputfile string

	// Prüfe, ob zu viele oder wenige Args gesetzt sind
	if len(os.Args) != 3 {
		color.Red("Error counting args")
		os.Exit(1)
	}

	// Checke Input
	inputfile = os.Args[2]

	// Check FileType
	if strings.HasSuffix(strings.ToLower(inputfile), ".pgn") {
		inputfile = inputfile + ""
	} else if !strings.ContainsRune(inputfile, '.') {
		inputfile = inputfile + ".pgn"
	} else {
		color.Red("Falsches Datei Format")
		os.Exit(1)
	}

	// Magic
	var _, err = os.Create(inputfile)
	if err != nil {
		color.Red("Can't create file")
		os.Exit(1)
	}
}

func deletepgn() {
	// Create a PGN File
	var inputfile string

	// Prüfe, ob zu viele oder wenige Args gesetzt sind
	if len(os.Args) != 3 {
		color.Red("Error counting args")
		os.Exit(1)
	}

	// Checke Input
	inputfile = os.Args[2]

	// Check FileType
	if strings.HasSuffix(strings.ToLower(inputfile), ".pgn") {
		inputfile = inputfile + ""
	} else if !strings.ContainsRune(inputfile, '.') {
		inputfile = inputfile + ".pgn"
	} else {
		color.Red("Falsches Datei Format")
		os.Exit(1)
	}

	// Magic
	var err = os.Remove(inputfile)
	if err != nil {
		color.Red("Can't create file")
		os.Exit(1)
	}
}

// JSON Block
func createjson() {
	// Create a PGN File
	var inputfile string

	// Prüfe, ob zu viele oder wenige Args gesetzt sind
	if len(os.Args) != 3 {
		color.Red("Error counting args")
		os.Exit(1)
	}

	// Checke Input
	inputfile = os.Args[2]

	// Check FileType
	if strings.HasSuffix(strings.ToLower(inputfile), ".json") {
		inputfile = inputfile + ""
	} else if !strings.ContainsRune(inputfile, '.') {
		inputfile = inputfile + ".json"
	} else {
		color.Red("Falsches Datei Format")
		os.Exit(1)
	}

	// Magic
	var _, err = os.Create(inputfile)
	if err != nil {
		color.Red("Can't create file")
		os.Exit(1)
	}
}

func deletejson() {
	// Create a PGN File
	var inputfile string

	// Prüfe, ob zu viele oder wenige Args gesetzt sind
	if len(os.Args) != 3 {
		color.Red("Error counting args")
		os.Exit(1)
	}

	// Checke Input
	inputfile = os.Args[2]

	// Check FileType
	if strings.HasSuffix(strings.ToLower(inputfile), ".json") {
		inputfile = inputfile + ""
	} else if !strings.ContainsRune(inputfile, '.') {
		inputfile = inputfile + ".json"
	} else {
		color.Red("Falsches Datei Format")
		os.Exit(1)
	}

	// Magic
	var err = os.Remove(inputfile)
	if err != nil {
		color.Red("Can't create file")
		os.Exit(1)
	}
}
