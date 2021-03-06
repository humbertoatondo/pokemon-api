package api

import (
	"net/http"
	"os"

	"github.com/humbertoatondo/pokemon-api/helpers"
	"github.com/humbertoatondo/pokemon-api/pokemon"
)

// Stores the results obtained on the comparePokemons function.
// It stores both pokemons data as well as the results for the damage comparision.
type advantage struct {
	Pokemon           pokemon.Pokemon        `json:"pokemon"`
	RivalPokemon      pokemon.Pokemon        `json:"rival_pokemon"`
	ComparisonResults pokemon.CompareResults `json:"comparison_results"`
}

// Receives two pokemon names as arguments and makes the following comparisions:
//   - Determine if pokemon1 can deal dobule damage to pokemon2.
//   - Determine if pokemon1 can receive half damage from pokemon2.
//   - Determine if pokemon1 can receive no damage from pokemon2.
// Returns a json containing the values in advantage struct.
func (app *App) comparePokemons(w http.ResponseWriter, r *http.Request) {
	// Extract pokemon names from url
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

	// Get pokemons
	pokemon1, err := pokemon.GetPokemon(pokemon1Name, os.Getenv("pokemon_url"), httpGet)
	if err != nil {
		helpers.RespondWithError(w, http.StatusForbidden, err.Error())
		return
	}

	pokemon2, err := pokemon.GetPokemon(pokemon2Name, os.Getenv("pokemon_url"), httpGet)
	if err != nil {
		helpers.RespondWithError(w, http.StatusForbidden, err.Error())
		return
	}

	// Compare pokemons damage relations
	comparisonResults, err := pokemon1.CompareTo(pokemon2, httpGet)
	if err != nil {
		helpers.RespondWithError(w, http.StatusForbidden, err.Error())
		return
	}

	advantageResult := advantage{
		Pokemon:           pokemon1,
		RivalPokemon:      pokemon2,
		ComparisonResults: comparisonResults,
	}

	helpers.RespondWithJSON(w, 200, advantageResult)
}
