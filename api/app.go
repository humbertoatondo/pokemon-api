package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// App contains the necessary components for runnign the server.
// In this case it just stores the router.
type App struct {
	Router *mux.Router
}

// Initialize inits the app router and the routes.
func (app *App) Initialize() {
	app.Router = mux.NewRouter()
	app.initializeRoutes()
}

// Run starts listening and serving in a given port.
func (app *App) Run(port string) {
	log.Fatal(http.ListenAndServe(port, app.Router))
}

func (app *App) initializeRoutes() {
	app.Router.HandleFunc("/comparePokemons", app.comparePokemons).Methods("GET")
	app.Router.HandleFunc("/comparePokemonsMoves", app.comparePokemonMoves).Methods("GET")
}

// Decouples the http get request from the rest of the functions.
// This makes it easier when testing the code.
var httpGet = func(url string) (*http.Response, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return response, nil
}
