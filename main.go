package main

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	Environment string `env:"ENVIRONMENT,required"`
	Version     int    `env:"VERSION,required"`
}

func main() {
	// Loading the environment variables from '.env' file.
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("unable to load .env file: %e", err)
	}

	cfg := Config{}

	// Parse the envrionment variables into Config struct.
	err = env.Parse(&cfg)
	if err != nil {
		log.Fatalf("unable to parse ennvironment variables: %e", err)
	}

	fmt.Println("Config:")
	fmt.Printf("Environment: %s\n", cfg.Environment)
	fmt.Printf("Version: %d\n", cfg.Version)
}
