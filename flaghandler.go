package main

import (
	"os"
	"strings"
)

func checkForUnknownFlags() {

	knownFlags := map[string]bool{
		"-help":       true,
		"-parsePGN":   true,
		"-parseJSON":  true,
		"-EventDate":  true,
		"-createpgn":  true,
		"-deletepgn":  true,
		"-createjson": true,
		"-deletejson": true,
	}

	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "-") {
			if _, exists := knownFlags[arg]; !exists {
				unknownflag()
				os.Exit(1)
			}
		}
	}
}
