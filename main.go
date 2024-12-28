package main

import (
	"log"
	"rune/cmd"
	"rune/db"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
