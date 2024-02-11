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
	color.Red("try pgn --help ?")
}

func parsepgn() {
	fmt.Println("Parsing PGN-Datei:")
}

func parsejson() {
	fmt.Println("Parsing JSON-Datei:")
}

func unknownflag() {
	color.Red("Unbekannte Flag: %s\n")
}
