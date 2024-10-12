package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log/slog"
	"os"
)

// HostConfig holds the host's URL and port settings
type HostConfig struct {
	Url  string `json:"url,omitempty" yaml:"url"`   // Host URL
	Port string `json:"port,omitempty" yaml:"port"` // Host port
}

// Load loads configuration from a YAML or JSON file at the specified path
func (c *HostConfig) Load(path string) error {
	const op = "Config.Load" // Operation name for error logging

	// Check if the config file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// Log and return if the file doesn't exist
		slog.Error(op + ": config file not found: " + err.Error())
		return err
	}

	// Read and parse the config file into HostConfig struct
	if err := cleanenv.ReadConfig(path, c); err != nil {
		// Log and return if reading the config fails
		slog.Error(op + ": error reading config: " + err.Error())
		return err
	}

	return nil // Successfully loaded configuration
}
