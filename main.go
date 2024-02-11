// main.go

package main

func main() {
	pgnFilePath := "./test/test.pgn"
	jsonFilePath := "./test/games.json"
	outputFile := "./test/parse.pgn"

	if err := parsePGNFile(pgnFilePath, jsonFilePath); err != nil {
		panic(err)
	}
	if err := parseJSONFile(jsonFilePath, outputFile); err != nil {
		panic(err)
	}

	println("true")
}

// Add only for Workflow
func Add(a, b int) int {
	return a + b
}
