package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	PK      string
	ADDMAIN string
)

func getEnv(env string) string {
	env, exists := os.LookupEnv(env)
	if !exists {
		log.Fatalf("Failed to load env variable: %v", env)
	}

	return env
}

func init() {
	godotenv.Load()

	PK = getEnv("PK")
	ADDMAIN = getEnv("ADDMAIN")
}
