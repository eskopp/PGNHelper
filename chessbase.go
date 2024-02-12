package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func removecb() {
	// Remove all pgi and ini files
	directory := "."

	extensions := []string{".pgi", ".ini"}

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		for _, ext := range extensions {
			if filepath.Ext(path) == ext {
				if err := os.Remove(path); err != nil {
					return err
				}
				fmt.Printf("Deleted: %s\n", path)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

}
