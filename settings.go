package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"lib/common"
	"os"
)

var settings *Settings

type Settings struct {
	// Non-Saved settings
	filename string

	// Saved settings
	InputDir  string
	OutputDir string
	LogFile   string
	LogLevel  common.LogLevel
}

var settingsExampleFS []byte

func LoadSettings(filename string) (*Settings, error) {
	raw, err := os.ReadFile(filename)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return nil, fmt.Errorf("could not read settings file: %w", err)
		}
		raw = settingsExampleFS
	}

	var s *Settings
	err = json.Unmarshal(raw, &s)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling: %w", err)
	}
	s.filename = filename

	var logFileDir string = s.LogFile
	fmt.Printf("Log file: %s\n", logFileDir)
	if err = common.SetupLogging(s.LogLevel, logFileDir); err != nil {
		return nil, fmt.Errorf("unable to setup logger: %w", err)
	}

	return s, nil
}
