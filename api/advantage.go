package api

import (
	"fmt"
	"net/http"
)

// type Pokemon struct {
// 	Name   string        `json:"name"`
// 	Ptypes []pokemonType `json:"types"`
// }

// type pokemonType struct {
// 	Slot int            `json:"slot"`
// 	Type pokemonSubType `json:"type"`
// }

// type pokemonSubType struct {
// 	Name string `json:"name"`
// 	URL  string `json:"url"`
// }

func (app *App) comparePokemons(w http.ResponseWriter, r *http.Request) {
	// Extract pokemon names from url.
	pokemon1Keys, ok := r.URL.Query()["pokemon1"]
	if !ok {
		respondWithError(w, 400, "Url param 'pokemon1' is missing.")
		return
	}

	pokemon2Keys, ok := r.URL.Query()["pokemon2"]
	if !ok {
		respondWithError(w, 400, "Url param 'pokemon2' is missing.")
		return
	}

	pokemon1 := pokemon1Keys[0]
	pokemon2 := pokemon2Keys[0]

	fmt.Printf("Pokemon 1: %s\n", pokemon1)
	fmt.Printf("Pokemon 2: %s\n", pokemon2)
	// =====================================================

	// Get pokemon types.

	// url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemon1)
	// response, err := http.Get(url)
	// if err != nil {
	// 	respondWithError(w, response.StatusCode, response.Status)
	// 	return
	// }

	// defer response.Body.Close()

	// var pokemon = Pokemon{}
	// if err = json.NewDecoder(response.Body).Decode(&pokemon); err != nil {
	// 	panic(err)
	// }

	// =====================================================

	respondWithJSON(w, 200, "Comparing Pokemons!")
}
