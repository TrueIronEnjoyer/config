package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log/slog"
	"os"
)

// DBConfig holds database connection details
type DBConfig struct {
	Url      string `json:"url,omitempty" yaml:"url"`           // Database URL
	Port     string `json:"port,omitempty" yaml:"port"`         // Database port
	Username string `json:"username,omitempty" yaml:"username"` // Database username
	Password string `json:"password,omitempty" yaml:"password"` // Database password
	FullUrl  string `json:"full_url,omitempty" yaml:"full_url"` // Full connection URL
}

// Load loads DBConfig from a YAML or JSON file at the specified path
func (c *DBConfig) Load(path string) error {
	const op = "DBConfig.Load" // Operation name for logging

	// Check if the config file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// Log and return error if the file doesn't exist
		slog.Error(op + ": config file not found: " + err.Error())
		return err
	}

	// Read and parse the config file into DBConfig struct
	if err := cleanenv.ReadConfig(path, c); err != nil {
		// Log and return error if reading config fails
		slog.Error(op + ": error reading config: " + err.Error())
		return err
	}

	return nil // Successfully loaded the configuration
}
