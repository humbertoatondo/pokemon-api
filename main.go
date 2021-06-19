package main

import "github.com/humbertoatondo/pokemon-api/api"

func main() {
	app := api.App{}
	app.Initialize()
	app.Run(":5000")
}
