package main

import (
	"log"

	"github.com/robbiew/go-doorserver/pkg/rlogin"

	"github.com/robbiew/go-doorserver/pkg/config"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Start RLOGIN server
	go rlogin.StartServer(cfg.Port, false)

	// Start Debug server if configured
	if cfg.DebugPort > 0 {
		go rlogin.StartServer(cfg.DebugPort, true)
	}

	// Keep the application running
	select {}
}
