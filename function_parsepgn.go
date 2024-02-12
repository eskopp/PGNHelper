package main

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"strings"
)

func parsepgn() {
	var inputpgn string
	var outjson string

	fmt.Println("Parsing PGN-Datei:")
	if len(os.Args) < 3 {
		color.Red("No Args")
		os.Exit(1)
	} else {
		inputpgn = os.Args[2]
	}

	// Prüft, ob die Inputdatei vorhanden ist
	fmt.Println(inputpgn)
	if _, err := os.Stat(inputpgn); os.IsNotExist(err) {
		color.Red(inputpgn + " File not exist")
		os.Exit(1)
	}

	// no json name
	if len(os.Args) < 4 {
		outjson = os.Args[3]
	} else {
		outjson = inputpgn[:len(inputpgn)-3] + "json"
	}

	// Prüft ob input pgn ist
	if !strings.HasSuffix(inputpgn, "pgn") {
		color.Red(inputpgn + " is no pgn file")
		os.Exit(1)
	}

	// Magic
	if err := parsePGNFile(inputpgn, outjson); err != nil {
		panic(err)
	}
}
