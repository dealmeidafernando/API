package main

import "github.com/kelseyhightower/envconfig"

// Config struct
type Config struct {
	AppName     string `envconfig:"APP_NAME" default:"API-Vagas"`
	AppUser     string `envconfig:"APP_USER" required:"true"`
	AppPassword string `envconfig:"APP_PASSWORD" required:"true"`
	AppPort     string `envconfig:"APP_PORT" required:"true"`
	AppDB       string `envconfig:"APP_DB" default:"vagas"`
}

// LoadConfig load Config
func LoadConfig() (Config, error) {
	var config Config
	err := envconfig.Process("", &config)
	return config, err
}
