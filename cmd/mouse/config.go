package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	PK           string
	ADD          string
	RPC_ENDPOINT string
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
	ADD = getEnv("ADD")
	RPC_ENDPOINT = getEnv("RPC_ENDPOINT")
}
