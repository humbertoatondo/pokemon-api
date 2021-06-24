package main

import (
	"flag"
	"log"
	"os"

	"github.com/humbertoatondo/pokemon-api/api"
	"github.com/joho/godotenv"
)

func main() {
	envPtr := flag.String("env", "development", "environment type")
	flag.Parse()
	env := *envPtr

	switch env {
	case "test":
		os.Setenv("pokemon_url", "http://localhost:5000/pokemon/")
	default:
		os.Setenv("pokemon_url", "https://pokeapi.co/api/v2/pokemon/")
	}

	if err := godotenv.Load(".env"); err != nil {
		log.Printf("Error loading .env file.\n")
	}

	app := api.App{}
	app.Initialize()
	app.Run(":5000")
}
