package main

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"os"
	"strings"
)

// Remove Eventdate from PGN Files
func eventdate() {
	// File Handler
	var inputfile string
	var filetype string
	var uuidfile string

	// legt uuid fest
	uuidfile = generateUUID()

	// Prüfe ob nicht zu viele Flags gesetzt sind
	if len(os.Args) != 3 {
		color.Red("To many args")
	}

	// Überträgt die Werte auf die Dateien
	if len(os.Args) < 3 {
		color.Red("No Args")
		os.Exit(1)
	} else {
		inputfile = os.Args[2]
	}

	// Teste ob die Datei PGN oder JSON ist
	filetype = "error"
	if strings.HasSuffix(strings.ToLower(inputfile), ".pgn") {
		filetype = "pgn"
	}
	if strings.HasSuffix(strings.ToLower(inputfile), ".json") {
		filetype = "json"
	}

	// Prüfe ob die Datei existiert
	if _, err := os.Stat(inputfile); os.IsNotExist(err) {
		color.Red(inputfile + " File not exist")
		os.Exit(1)
	}

	color.Yellow("input:" + inputfile)
	color.Yellow("filetype:" + filetype)
	color.Yellow("UUIDFile:" + uuidfile)

	// Wandelt die PGN temporär in eine JSON um
	if filetype == "pgn" {
		file, err := os.Create(uuidfile + ".json")
		if err != nil {
			fmt.Println("Fehler beim Erstellen der Datei:", err)
			return
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {

			}
		}(file)
		if err := parsePGNFile(inputfile, uuidfile+".json"); err != nil {
			panic(err)
		}
	}

	// Inhalt der JSON-Datei lesen
	var inputBytes, err = ioutil.ReadFile(uuidfile + ".json")
	if err != nil {
		color.Red("Fehler beim Lesen der JSON-Datei:", err)
		os.Exit(1)
	}

	// JSON-Daten in ein Slice von Game-Strukturen parsen
	var games []Game
	if err := json.Unmarshal(inputBytes, &games); err != nil {
		color.Red("Fehler beim Parsen der JSON-Daten:", err)
		os.Exit(1)
	}

	// Durch alle Spiele iterieren und EventDate entfernen
	for i := range games {
		delete(games[i].Tags, "EventDate")
	}

	// JSON-Daten mit den entfernten EventDate in die Ausgabedatei schreiben
	outputBytes, err := json.MarshalIndent(games, "", "    ")
	if err != nil {
		color.Red("Fehler beim Konvertieren der Daten in JSON:", err)
		os.Exit(1)
	}

	// Inhalt in die Ausgabedatei schreiben
	if err := ioutil.WriteFile(inputfile, outputBytes, 0644); err != nil {
		color.Red("Fehler beim Schreiben der JSON-Datei:", err)
		os.Exit(1)
	}

	fmt.Println("EventDate-Einträge wurden erfolgreich aus der JSON-Datei entfernt und die modifizierten Daten wurden in die Datei", inputfile, "geschrieben.")

	// Wenn die ursprungsdatei eine PGN Datei war, lösche die ursprüngliche Datei
	// und wandel es wieder als PGN um
	if filetype == "pgn" {
		err := os.Remove(inputfile)
		if err != nil {
			color.Red("Can't delete file:", err)
			return
		}
		if err := parseJSONFile(uuidfile+".json", inputfile); err != nil {
			panic(err)
		}
	}
}
