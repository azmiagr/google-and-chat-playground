package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvironment() error {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatalf("failed to load env, err: %v", err)
	}
	return nil
}
