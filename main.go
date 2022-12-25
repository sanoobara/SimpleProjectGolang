package main

import (
	"fmt"
	"lib/internal/files"
	"log"
	"time"
)

func setupSettings(confFile string) error {
	if confFile == "" {
		confFile = files.JoinRunPath("settings.json")
	}

	var err error
	settings, err = LoadSettings(confFile)
	if err != nil {
		return fmt.Errorf("unable to load settings: %w", err)
	}

	return nil
}

func main() {

	run()

}

func run() {

	start := time.Now()

	fmt.Println(start)

	if err := setupSettings("settings.json"); err != nil {
		log.Fatalf("Error loading settings: %v\n", err)
	}
	fmt.Println(settings.InputDir)
	fmt.Println(settings.filename)
	fmt.Println(settings.OutputDir)

}
