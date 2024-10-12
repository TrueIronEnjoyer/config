package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log/slog"
	"os"
)

type Config struct {
	DB   DBConfig   `json:"DB" yaml:"db"`
	Host HostConfig `json:"Host" yaml:"host"`
}

func (c *Config) Load(path string) error {
	const op = "Config.Load"

	if _, err := os.Stat(path); os.IsNotExist(err) {
		slog.Error(op + ": " + err.Error())
	}

	if err := cleanenv.ReadConfig(path, c); err != nil {
		slog.Error(op + ": " + err.Error())
	}

	return nil
}
