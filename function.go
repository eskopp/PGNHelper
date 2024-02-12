package main

import (
	"fmt"
	"github.com/fatih/color"
	"os"
)

func pgnHelp() {
	color.Red("PGNHelper Nutzung:")
	fmt.Println("  -help: Zeigt diese Hilfe an.")
	fmt.Println("  Ohne Argumente: Gibt einen Standardtext aus.")
	fmt.Println("  -parsePGN \"datei.pgn\": Echo einer Nachricht (Platzhalter für PGN-Parsing-Logik).")
	fmt.Println("  -parseJSON \"datei.json\": Echo einer Nachricht (Platzhalter für JSON-Parsing-Logik).")
	os.Exit(0)
}

func noflag() {
	color.Red("try pgn -help ?")
}

func parsepgn() {
	var inputpgn string
	var outjson string

	fmt.Println("Parsing PGN-Datei:")
	if len(os.Args) < 4 {
		color.Red("No Args")
		os.Exit(1)
	} else {
		inputpgn = os.Args[2]
		outjson = os.Args[3]
	}

	// Prüft, ob die Inputdatei vorhanden ist
	fmt.Println(inputpgn)
	if _, err := os.Stat(inputpgn); os.IsNotExist(err) {
		color.Red(inputpgn + " File not exist")
		os.Exit(1)
	}

	// Hier findet die eigentliche Verarbeitung statt
	if err := parsePGNFile(inputpgn, outjson); err != nil {
		panic(err)
	}
}

func parsejson() {
	var inputjson string
	var outpgn string

	fmt.Println("Parsing PGN-Datei:")
	if len(os.Args) < 4 {
		color.Red("No Args")
		os.Exit(1)
	} else {
		inputjson = os.Args[2]
		outpgn = os.Args[3]
	}

	// Prüft, ob die Inputdatei vorhanden ist
	fmt.Println(inputjson)
	if _, err := os.Stat(inputjson); os.IsNotExist(err) {
		color.Red(inputjson + " File not exist")
		os.Exit(1)
	}

	// Magic
	if err := parseJSONFile(inputjson, outpgn); err != nil {
		panic(err)
	}
}

func unknownflag() {
	color.Red("Unbekannte Flag: %s\n")
}
