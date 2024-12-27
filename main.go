package main

import (
	"log"
	"rune/cmd"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
