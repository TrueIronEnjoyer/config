package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log/slog"
	"os"
)

type HostConfig struct {
	Url  string `json:"url,omitempty" yaml:"url"`
	Port string `json:"port,omitempty" yaml:"port"`
}

func (c *HostConfig) Load(path string) error {
	const op = "Config.Load"

	if _, err := os.Stat(path); os.IsNotExist(err) {
		slog.Error(op + ": " + err.Error())
	}

	if err := cleanenv.ReadConfig(path, c); err != nil {
		slog.Error(op + ": " + err.Error())
	}

	return nil
}
