package api

import (
	"net/http"
	"strconv"

	"github.com/humbertoatondo/pokemon-api/pokemon"

	"github.com/humbertoatondo/pokemon-api/helpers"
)

// Stores the results obtained on the comparePokemonMoves function.
// It stores the list of pokemons, the list of common moves between the
// pokemons and the desired language of the moves names.
type commonMoves struct {
	Pokemons []string           `json:"pokemons"`
	Moves    []pokemon.MoveData `json:"moves"`
	Lang     string             `json:"lang"`
}

// Receives a list of pokemon names, a language and a limit as arguments.
// Gets the list of common moves among the list of pokemons, it can limit the
// amount of desired common moves to get with the limit argument.
// It can also translate the name of the common moves to the languages listed
// int LaguagesMap located in the helpers package.
// Returns a json containing the values in commonMoves struct.
func (app *App) comparePokemonMoves(w http.ResponseWriter, r *http.Request) {
	// Extract arguments from url
	keys, ok := r.URL.Query()["pokemon"]
	if !ok || len(keys) < 1 {
		helpers.RespondWithError(w, http.StatusBadRequest, "Url param 'pokemon' is missing.")
		return
	} else if len(keys) == 1 {
		helpers.RespondWithError(w, http.StatusBadRequest, "Please provide at least 2 pokemon names.")
		return
	}

	lang, ok := helpers.ParseKeyFromURL("lang", r)
	if !ok {
		lang = "en"
	}

	if _, ok := helpers.LanguageMap[lang]; !ok {
		lang = "en"
	}

	limitStr, ok := helpers.ParseKeyFromURL("limit", r)
	if !ok {
		limitStr = "10"
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error while parsing limit of moves.")
		return
	}

	pokemons, err := pokemon.GetPokemonsFromListOfNames(keys, httpGet)
	if err != nil {
		helpers.RespondWithError(w, http.StatusForbidden, err.Error())
		return
	}

	commonMovesForPokemons := pokemon.GetCommonMovesForPokemons(pokemons, limit)

	translatedCommonMovesForPokemons, err := pokemon.TranslatePokemonMoves(commonMovesForPokemons, lang, httpGet)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	result := commonMoves{
		Pokemons: keys,
		Moves:    translatedCommonMovesForPokemons,
		Lang:     lang,
	}

	helpers.RespondWithJSON(w, 200, result)
}
