package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func (app *App) Initialize() {
	app.Router = mux.NewRouter()
	app.initializeRoutes()
}

func (app *App) Run(port string) {
	log.Fatal(http.ListenAndServe(port, app.Router))
}

func (app *App) initializeRoutes() {
	app.Router.HandleFunc("/comparePokemons", app.comparePokemons).Methods("GET")
	app.Router.HandleFunc("/comparePokemonsMoves", app.comparePokemonMoves).Methods("GET")
}
