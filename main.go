package main

import (
	"flag"
)

func main() {

	// Flag für die Hilfe
	helpFlag := flag.Bool("help", false, "Zeigt die Hilfe an.")

	// Flags für das Parsen einer PGN-Datei und einer JSON-Datei, als boolesche Schalter
	parsePGNFlag := flag.Bool("parsePGN", false, "Verarbeitet eine PGN-Datei.")
	parseJSONFlag := flag.String("parseJSON", "", "Pfad zur JSON-Datei, die geparst werden soll.")

	// Eigene Hilfe-Funktion einrichten
	flag.Usage = pgnHelp

	// Parsen der Flags
	flag.Parse()

	// Help Flag
	if *helpFlag {
		pgnHelp()
		return
	}

	// Parse PGN Flag - boolescher Schalter
	if *parsePGNFlag {
		parsepgn() // Keine Argumente mehr erforderlich
		return
	}

	// Parse JSON Flag
	if *parseJSONFlag != "" {
		parsejson()
		return
	}

	// Wenn keine Flag gesetzt wird
	noflag()
}
