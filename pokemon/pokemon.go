package pokemon

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Pokemon struct {
	Name   string        `json:"name"`
	Ptypes []pokemonType `json:"types"`
}

type pokemonType struct {
	Slot int            `json:"slot"`
	Type pokemonSubType `json:"type"`
}

type pokemonSubType struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func getPokemon(pokemonName string) (Pokemon, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemonName)
	response, err := http.Get(url)
	if err != nil {
		// respondWithError(w, response.StatusCode, response.Status)
		return Pokemon{}, err
	}

	defer response.Body.Close()

	var pokemon = Pokemon{}
	if err = json.NewDecoder(response.Body).Decode(&pokemon); err != nil {
		panic(err)
	}

	return pokemon, nil
}
