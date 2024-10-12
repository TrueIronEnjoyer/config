package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log/slog"
	"os"
)

// Config holds the overall configuration with DB and Host sections
type Config struct {
	DB   DBConfig   `json:"db,omitempty" yaml:"db"`     // Database configuration
	Host HostConfig `json:"host,omitempty" yaml:"host"` // Host configuration
}

// Load loads the configuration from a YAML or JSON file at the specified path
func (c *Config) Load(path string) {
	const op = "Config.Load" // Operation name for logging

	// Check if the config file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// Log and return error if the file is not found
		slog.Error(op + ": config file not found: " + err.Error())
		panic(op + ": config file not found: " + err.Error())
	}

	// Read and parse the configuration file into the Config struct
	if err := cleanenv.ReadConfig(path, c); err != nil {
		// Log and return error if reading the config file fails
		slog.Error(op + ": error reading config: " + err.Error())
		panic(op + ": error reading config: " + err.Error())
	}

	return // Successfully loaded configuration
}
