package api

import (
	"net/http"

	"github.com/humbertoatondo/pokemon-api/helpers"
	"github.com/humbertoatondo/pokemon-api/pokemon"
)

type advantage struct {
	Pokemon            pokemon.Pokemon        `json:"pokemon"`
	RivalPokemon       pokemon.Pokemon        `json:"rival_pokemon"`
	ComparisionResults pokemon.CompareResults `json:"comparision_results"`
}

func (app *App) comparePokemons(w http.ResponseWriter, r *http.Request) {

	pokemon1Name, ok := helpers.ParseKeyFromURL("pokemon1", r)
	if !ok {
		helpers.RespondWithError(w, http.StatusBadRequest, "Url param 'pokemon1' is missing.")
		return
	}

	pokemon2Name, ok := helpers.ParseKeyFromURL("pokemon2", r)
	if !ok {
		helpers.RespondWithError(w, http.StatusBadRequest, "Url param 'pokemon2' is missing.")
		return
	}

	// Get pokemons.
	pokemon1, err := pokemon.GetPokemon(pokemon1Name)
	if err != nil {
		helpers.RespondWithError(w, http.StatusForbidden, err.Error())
		return
	}

	pokemon2, err := pokemon.GetPokemon(pokemon2Name)
	if err != nil {
		helpers.RespondWithError(w, http.StatusForbidden, err.Error())
		return
	}

	// Compare pokemons
	comparisionResults, err := pokemon1.CompareTo(pokemon2)
	if err != nil {
		helpers.RespondWithError(w, http.StatusForbidden, err.Error())
		return
	}

	advantageResult := advantage{
		Pokemon:            pokemon1,
		RivalPokemon:       pokemon2,
		ComparisionResults: comparisionResults,
	}

	helpers.RespondWithJSON(w, 200, advantageResult)
}
