package config

import (
	"errors"
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	POSTGRES_HOST     string `env:"POSTGRES_HOST"`
	POSTGRES_PORT     string `env:"POSTGRES_PORT"`
	POSTGRES_DB       string `env:"POSTGRES_DB"`
	POSTGRES_USER     string `env:"POSTGRES_USER"`
	POSTGRES_PASSWORD string `env:"POSTGRES_PASSWORD"`
	APP_IP            string `env:"APP_IP"`
	APP_PORT          string `env:"APP_PORT"`
}

func Init() (*Config, error) {
	logrus.Debug("Init env config")
	// Загрузка переменных из файла .env
	if err := godotenv.Load(); err != nil {
		logrus.Warn("No .env file found")
	}

	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, errors.New("Failed to parse config: " + err.Error())
	}
	return &cfg, nil
}
