package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log/slog"
	"os"
)

type DBConfig struct {
	Url      string `json:"url,omitempty" yaml:"url"`
	Port     string `json:"port,omitempty" yaml:"port"`
	Username string `json:"username,omitempty" yaml:"username"`
	Password string `json:"password,omitempty" yaml:"password"`
	FullUrl  string `json:"full_url,omitempty" yaml:"full_url"`
}

func (c *DBConfig) Load(path string) error {
	const op = "DBConfig.Load"

	if _, err := os.Stat(path); os.IsNotExist(err) {
		slog.Error(op + ": " + err.Error())
	}

	if err := cleanenv.ReadConfig(path, c); err != nil {
		slog.Error(op + ": " + err.Error())
	}

	return nil
}
