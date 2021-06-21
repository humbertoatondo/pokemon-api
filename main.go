package main

import (
	"log"

	"github.com/humbertoatondo/pokemon-api/api"
	"github.com/joho/godotenv"
)

func main() {
	// envPtr := flag.String("env", "development", "a string")
	// flag.Parse()
	// env := *envPtr

	// fmt.Println(env)

	// switch env {
	// case "test":
	// 	os.Setenv("pokemon_url", "http://localhost:5000/pokemon/")
	// default:
	// 	os.Setenv("pokemon_url", "https://pokeapi.co/api/v2/pokemon/")
	// }

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error laoding .env file")
	}

	app := api.App{}
	app.Initialize()
	app.Run(":5000")
}
