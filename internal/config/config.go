package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env string `yaml:"env" env-default:"local"`
	Storage
	Cache
	Server
}

type Storage struct {
	Region                string `yaml:"region" env-default:"ru-central1" env-require:"true"`
	Endpoint_url          string `yaml:"endpoint_url" env-default:"https://storage.yandexcloud.net/" env-require:"true"`
	Aws_access_key_id     string `yaml:"aws_access_key_id" env-require:"true"`
	Aws_secret_access_key string `yaml:"aws_secret_access_key" env-require:"true"`
	BucketName            string `yaml:"bucketname" env-require:"true"`
}

type Cache struct {
	Address  string        `yaml:"address" end-default:"localhost:6379"`
	Password string        `yaml:"password" env-default:""`
	DB       int           `yaml:"db" env-default:"0"`
	TTL      time.Duration `yaml:"ttl" env-default:"10m"`
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
