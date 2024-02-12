package main

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"strings"
)

func parsejson() {
	var inputjson string
	var outpgn string

	fmt.Println("Parsing PGN-Datei:")
	if len(os.Args) < 3 {
		color.Red("No Args")
		os.Exit(1)
	} else {
		inputjson = os.Args[2]
	}

	// Prüft, ob die Inputdatei vorhanden ist
	fmt.Println(inputjson)
	if _, err := os.Stat(inputjson); os.IsNotExist(err) {
		color.Red(inputjson + " File not exist")
		os.Exit(1)
	}

	// no json name
	if len(os.Args) < 4 {
		outpgn = os.Args[3]
	} else {
		outpgn = inputjson[:len(inputjson)-3] + "json"
	}

	// Prüft ob input json ist
	if !strings.HasSuffix(inputjson, "json") {
		color.Red(inputjson + " is no json file")
		os.Exit(1)
	}

	// Magic
	if err := parseJSONFile(inputjson, outpgn); err != nil {
		panic(err)
	}
}
