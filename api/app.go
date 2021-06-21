package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/humbertoatondo/pokemon-api/helpers"
	"github.com/humbertoatondo/pokemon-api/pokemon"
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
	app.Router.HandleFunc("/", app.home).Methods("GET")
	app.Router.HandleFunc("/comparePokemons", app.comparePokemons).Methods("GET")
	app.Router.HandleFunc("/comparePokemonsMoves", app.comparePokemonMoves).Methods("GET")

	app.Router.HandleFunc("/pokemon/{pokemon}", app.mockGetPokemon).Methods("GET")
	app.Router.HandleFunc("/move/{move}", app.mockGetMove).Methods("GET")
	app.Router.HandleFunc("/type/{type}", app.mockGetType).Methods("GET")
}

func (app *App) home(w http.ResponseWriter, r *http.Request) {
	welcomeMessage := "Pokemon Rest API"
	helpers.RespondWithJSON(w, 200, welcomeMessage)
}

func (app *App) mockGetPokemon(w http.ResponseWriter, r *http.Request) {
	fileBytes, err := parseURLAndGetFile(r, "pokemon")
	if err != nil {
		mssg := "Error while trying to open json file in /testing_resources/pokemons"
		helpers.RespondWithError(w, 500, mssg)
		return
	}

	var inter pokemon.Pokemon
	json.Unmarshal(fileBytes, &inter)

	helpers.RespondWithJSON(w, 200, inter)
}

func (app *App) mockGetMove(w http.ResponseWriter, r *http.Request) {
	fileBytes, err := parseURLAndGetFile(r, "move")
	if err != nil {
		mssg := "Error while trying to open json file in /testing_resources/moves"
		helpers.RespondWithError(w, 500, mssg)
		return
	}

	var inter pokemon.TransMoves
	json.Unmarshal(fileBytes, &inter)

	helpers.RespondWithJSON(w, 200, inter)
}

func (app *App) mockGetType(w http.ResponseWriter, r *http.Request) {
	fileBytes, err := parseURLAndGetFile(r, "type")
	if err != nil {
		mssg := "Error while trying to open json file in /testing_resources/types"
		helpers.RespondWithError(w, 500, mssg)
		return
	}

	var inter pokemon.PDamageRelations
	json.Unmarshal(fileBytes, &inter)

	helpers.RespondWithJSON(w, 200, inter)
}

func parseURLAndGetFile(r *http.Request, sType string) ([]byte, error) {
	path := r.URL.Path
	pathFragments := strings.Split(path, "/")
	targetName := pathFragments[2]

	jsonFilePath := fmt.Sprintf("testing_resources/%ss/%s_%s.json", sType, sType, targetName)
	fileBytes, err := ioutil.ReadFile(jsonFilePath)
	if err != nil {
		return nil, err
	}

	return fileBytes, nil
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
