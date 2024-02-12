package main

import (
	"flag"
)

func main() {

	// Überprüfung auf unbekannte Flags, bevor `flag.Parse()` aufgerufen wird
	checkForUnknownFlags()

	// Flag für die Hilfe
	helpFlag := flag.Bool("help", false, "Zeigt die Hilfe an.")

	// Flags für das Parsen einer PGN-Datei und einer JSON-Datei
	parsePGNFlag := flag.Bool("parsePGN", false, "Verarbeitet eine PGN-Datei.")
	parseJSONFlag := flag.String("parseJSON", "", "Pfad zur JSON-Datei, die geparst werden soll.")
	parseEventDateFlag := flag.String("parseEventDate", "", "Pfad zur JSON-Datei, aus der das Event-Datum entfernt werden soll.")

	// Eigene Hilfe-Funktion einrichten
	flag.Usage = pgnHelp

	// Parsen der Flags
	flag.Parse()

	// Help Flag
	if *helpFlag {
		pgnHelp()
		return
	}

	// Parse PGN Flag
	if *parsePGNFlag {
		parsepgn()
		return
	}

	// Parse Event Date Flag
	if *parseEventDateFlag != "" {
		eventdate()
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
