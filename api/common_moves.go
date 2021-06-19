package api

import (
	"fmt"
	"net/http"

	"github.com/humbertoatondo/pokemon-api/pokemon"

	"github.com/humbertoatondo/pokemon-api/helpers"
)

func (app *App) comparePokemonMoves(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["pokemon"]
	if !ok || len(keys) < 1 {
		helpers.RespondWithError(w, http.StatusBadRequest, "Url param 'pokemon' is missing.")
		return
	} else if len(keys) == 1 {
		helpers.RespondWithError(w, http.StatusBadRequest, "Please provide at least 2 pokemon names.")
		return
	}

	pokemons, err := pokemon.GetPokemonsFromListOfNames(keys)
	if err != nil {
		helpers.RespondWithError(w, http.StatusForbidden, err.Error())
		return
	}

	// Get common moves
	commonMoves := pokemon.GetCommonMovesForPokemons(pokemons)

	fmt.Printf("%v\n", commonMoves)

	helpers.RespondWithJSON(w, 200, "Common pokemon moves!")
}
