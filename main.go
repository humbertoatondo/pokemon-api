package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/humbertoatondo/pokemon-api/api"
)

func main() {
	envPtr := flag.String("env", "development", "a string")
	flag.Parse()
	env := *envPtr

	fmt.Println(env)

	switch env {
	case "test":
		os.Setenv("pokemon_url", "http://localhost:5000/pokemon/")
	default:
		os.Setenv("pokemon_url", "https://pokeapi.co/api/v2/pokemon/")
	}

	app := api.App{}
	app.Initialize()
	app.Run(":5000")
}
