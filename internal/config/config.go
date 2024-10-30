package config

import (
	"fmt"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string `yaml:"env" env-default:"local"`
	StorageURL string `yaml:"storage_url"`
	Server
}

type Server struct {
	Address string `yaml:"address" env-default:"localhost:8080"`
}

func NewConfig(nameConfig string) *Config {
	pathConfig := fmt.Sprintf("./config/%s.yaml", nameConfig)
	err := os.Setenv("CONFIG", pathConfig)
	if err != nil {
		log.Fatal("Failed to set env", err)
	}

	pathConfig = os.Getenv("CONFIG")
	if pathConfig == "" {
		log.Fatalf("Failed to get env")
	}

	if _, err := os.Stat(pathConfig); os.IsNotExist(err) {
		log.Fatalf("File is not exists: %s", err)
	}

	var cfg Config

	err = cleanenv.ReadConfig(pathConfig, &cfg)
	if err != nil {
		log.Fatal("Failed to read config: %w", err)
	}

	return &cfg

}
