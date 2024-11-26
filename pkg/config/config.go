package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Port      int       `json:"port"`
	DebugPort int       `json:"debugPort"`
	Dosbox    Dosbox    `json:"dosbox"`
	Doors     []Door    `json:"doors"`
	DebugUser DebugUser `json:"debugUser"`
}

type Dosbox struct {
	Path      string `json:"dosboxPath"`
	Config    string `json:"configPath"`
	DrivePath string `json:"drivePath"`
	StartPort int    `json:"startPort"`
	Headless  bool   `json:"headless"`
}

type Door struct {
	Code           string `json:"code"`
	Command        string `json:"doorCmd"`
	DropFileDir    string `json:"dropFileDir"`
	Description    string `json:"description"`
	Category       string `json:"category"`
	GameTitle      string `json:"gameTitle"`
	YearCreated    int    `json:"yearCreated"`
	MultiNode      bool   `json:"multiNode"`
	RemoveLockFile string `json:"removeLockFile"`
}

type DebugUser struct {
	Name     string `json:"name"`
	Module   string `json:"module"`
	Terminal string `json:"terminal"`
}

func LoadConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, fmt.Errorf("failed to decode config file: %w", err)
	}

	return &config, nil
}
