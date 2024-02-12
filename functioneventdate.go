package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"os"
	"strings"
)

// Remove Eventdate from PGN Files
func eventdate() {
	// File Handler
	var inputfile string

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

	// Prüfe ob die Datei existiert
	if _, err := os.Stat(inputfile); os.IsNotExist(err) {
		color.Red(inputfile + " File not exist")
		os.Exit(1)
	}

	err := removeEventDateLines(inputfile)
	if err != nil {
		color.Red("Error:")
		fmt.Println(err)
		return
	}
}

func removeEventDateLines(filename string) error {
	// Öffne die Datei zum Lesen
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	// Erstelle einen temporären Dateinamen für die Ausgabe
	tempFilename := filename + ".tmp"

	// Öffne eine temporäre Datei zum Schreiben
	tempFile, err := os.Create(tempFilename)
	if err != nil {
		return err
	}
	defer func(tempFile *os.File) {
		err := tempFile.Close()
		if err != nil {

		}
	}(tempFile)

	// Lesen und Bearbeiten der Zeilen
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.Contains(line, "EventDate") {
			// Schreibe die Zeile in die temporäre Datei, wenn "EventDate" nicht enthalten ist
			_, err := fmt.Fprintln(tempFile, line)
			if err != nil {
				return err
			}
		}
	}

	// Überprüfe auf Fehler beim Scannen
	if err := scanner.Err(); err != nil {
		return err
	}

	// Schließe die ursprüngliche Datei
	if err := file.Close(); err != nil {
		return err
	}

	// Schließe die temporäre Datei
	if err := tempFile.Close(); err != nil {
		return err
	}

	// Lösche die ursprüngliche Datei
	if err := os.Remove(filename); err != nil {
		return err
	}

	// Benenne die temporäre Datei in den ursprünglichen Dateinamen um
	if err := os.Rename(tempFilename, filename); err != nil {
		return err
	}

	return nil
}
