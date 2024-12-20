package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"time"
)

type Config struct {
	Env            string `yaml:"env" env-default:"local"`
	Bd             `yaml:"database"`
	Server         `yaml:"server"`
	Redis          `yaml:"redis"`
	MigrationsPath string `yaml:"migrations"`
}

type Bd struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
}

type Server struct {
	Port        string        `yaml:"port"`
	Host        string        `yaml:"host"`
	Timeout     time.Duration `yaml:"timeout"`
	IdleTimeout time.Duration `yaml:"iddle_timeout"`
}

type Redis struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatalf("CONFIG_PATH is required")
	}

	if _, err := os.Stat(configPath); err != nil {
		log.Fatalf("Failed to open config file: %v", err)
	}

	c := Config{}
	file, err := os.ReadFile(configPath)

	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}
	replaced := os.ExpandEnv(string(file))
	err = yaml.Unmarshal([]byte(replaced), &c)
	if err != nil {
		log.Fatalf("Failed to parse config file: %v", err)
	}

	return &c

}

func (c *Config) DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", c.Bd.User, c.Bd.Password, c.Bd.Host, c.Bd.Port, c.Bd.Name)
}
