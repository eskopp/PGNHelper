package main

import (
	"bufio"
	"encoding/json"
	"os"
	"strings"
)

// parsePGNFile convert to json
func parsePGNFile(pgnFilePath, jsonFilePath string) error {
	file, err := os.Open(pgnFilePath)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	scanner := bufio.NewScanner(file)
	var games []Game
	currentGame := Game{Tags: make(map[string]string)}
	var movesBuilder strings.Builder

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "[") {
			tagContents := line[1:strings.LastIndex(line, "]")]
			parts := strings.SplitN(tagContents, " ", 2)
			if len(parts) == 2 {
				key := parts[0]
				value := strings.Trim(parts[1], "\"")
				currentGame.Tags[key] = value
			}
		} else if line != "" {
			movesBuilder.WriteString(line + " ")
		} else if movesBuilder.Len() > 0 {
			currentGame.Moves = movesBuilder.String()
			games = append(games, currentGame)
			currentGame = Game{Tags: make(map[string]string)}
			movesBuilder.Reset()
		}
	}
	if movesBuilder.Len() > 0 || len(currentGame.Tags) > 0 {
		currentGame.Moves = movesBuilder.String()
		games = append(games, currentGame)
	}

	jsonData, err := json.MarshalIndent(games, "", "    ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(jsonFilePath, jsonData, 0644); err != nil {
		return err
	}

	return nil
}
