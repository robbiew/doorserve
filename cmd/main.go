package main

import (
	"log"

	"github.com/robbiew/doorserve/internal/server"

	"github.com/robbiew/doorserve/internal/config"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Start RLOGIN server
	go server.StartServer(cfg.Port, false)

	// Start Debug server if configured
	if cfg.DebugPort > 0 {
		go server.StartServer(cfg.DebugPort, true)
	}

	// Keep the application running
	select {}
}
