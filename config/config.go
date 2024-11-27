package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"time"
)

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type ServerConfig struct {
	StartTime        time.Time `yaml:"start_time"`
	ApiKey           string    `yaml:"api_key"`
	SharePoolAddress string    `yaml:"share_pool_address"`
	TrackingHash     string    `yaml:"tracking_hash"`
}

type config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
}

var Config config

func Init() {
	yamlFile, err := os.ReadFile("/app/config/config.yaml")
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}
