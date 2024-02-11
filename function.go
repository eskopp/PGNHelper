package main

import (
	"fmt"
	"os"
)

func pgnHelp() {
	fmt.Println("PGNHelper Nutzung:")
	fmt.Println("  -help: Zeigt diese Hilfe an.")
	fmt.Println("  Ohne Argumente: Gibt einen Standardtext aus.")
	fmt.Println("  -parsePGN \"datei.pgn\": Echo einer Nachricht (Platzhalter für PGN-Parsing-Logik).")
	fmt.Println("  -parseJSON \"datei.json\": Echo einer Nachricht (Platzhalter für JSON-Parsing-Logik).")
	os.Exit(0)
}

func noflag() {
	// TODO: Solange das Tool nicht fertig ist, kommt hier die Hilfe ^^
	pgnHelp()
}

func parsepgn() { // Parameter hinzugefügt
	fmt.Println("Parsing PGN-Datei:")
	// Implementierung der PGN-Parsing-Logik
}

func parsejson() {
	fmt.Println("Parsing JSON-Datei:")
	// Implementierung der JSON-Parsing-Logik
}
