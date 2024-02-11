package main

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

// Game-Struct, das die Details eines Schachspiels hält
type Game struct {
	Event   string
	Site    string
	Date    string
	Round   string
	White   string
	Black   string
	Result  string
	Opening string
	ECO     string
	Moves   string
}

// parsePGNFile liest die angegebene PGN-Datei und extrahiert die Schachspiele
func parsePGNFile(filePath string) ([]Game, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	pgnText := strings.Join(lines, "\n")
	return parsePGN(pgnText), nil
}

// parsePGN verarbeitet den gegebenen PGN-Text und extrahiert jedes Spiel
func parsePGN(pgn string) []Game {
	var games []Game
	gameStrings := strings.Split(pgn, "\n\n\n")
	for _, gameString := range gameStrings {
		var game Game
		metaRe := regexp.MustCompile(`\[(\w+) "(.*?)"\]`)
		metaMatches := metaRe.FindAllStringSubmatch(gameString, -1)

		for _, match := range metaMatches {
			switch match[1] {
			case "Event":
				game.Event = match[2]
			case "Site":
				game.Site = match[2]
			case "Date":
				game.Date = match[2]
			case "Round":
				game.Round = match[2]
			case "White":
				game.White = match[2]
			case "Black":
				game.Black = match[2]
			case "Result":
				game.Result = match[2]
			case "Opening":
				game.Opening = match[2]
			case "ECO":
				game.ECO = match[2]
			}
		}

		movesStartIndex := metaRe.FindAllStringIndex(gameString, -1)
		if len(movesStartIndex) > 0 {
			lastMatch := movesStartIndex[len(movesStartIndex)-1]
			game.Moves = strings.TrimSpace(gameString[lastMatch[1]:])
		}

		games = append(games, game)
	}
	return games
}
