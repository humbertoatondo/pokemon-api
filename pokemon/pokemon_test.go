package pokemon

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPokemon(t *testing.T) {
	var httpGet = func(url string) (*http.Response, error) {

		tData := pokemonTypeData{
			Name: "normal",
			URL:  "https://pokeapi.co/api/v2/type/1/",
		}

		mData := MoveData{
			Name: "transform",
			URL:  "https://pokeapi.co/api/v2/move/144/",
		}

		typesData := []pokemonTypeData{tData}
		movesData := []MoveData{mData}
		newPokemon := createPokemon("ditto", typesData, movesData)

		reqBodyBytes := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(newPokemon)
		newPokemonBytes := reqBodyBytes.Bytes()
		response := http.Response{Body: ioutil.NopCloser(bytes.NewBuffer(newPokemonBytes))}

		return &response, nil
	}

	response, _ := httpGet("url")

	var pokemon = Pokemon{}
	if err := json.NewDecoder(response.Body).Decode(&pokemon); err != nil {
		t.Error("Expected a pokemon.")
	}

	tData := pokemonTypeData{
		Name: "normal",
		URL:  "https://pokeapi.co/api/v2/type/1/",
	}

	mData := MoveData{
		Name: "transform",
		URL:  "https://pokeapi.co/api/v2/move/144/",
	}

	typesData := []pokemonTypeData{tData}
	movesData := []MoveData{mData}
	basePokemon := createPokemon("ditto", typesData, movesData)

	assert.Equal(t, basePokemon, pokemon)
}

func TestCompareTo(t *testing.T) {
	// var httpGet = func(url string) (*http.Response, error) {

	// 	tData := pokemonTypeData{
	// 		Name: "normal",
	// 		URL:  "https://pokeapi.co/api/v2/type/1/",
	// 	}

	// 	mData := MoveData{
	// 		Name: "transform",
	// 		URL:  "https://pokeapi.co/api/v2/move/144/",
	// 	}

	// 	typesData := []pokemonTypeData{tData}
	// 	movesData := []MoveData{mData}
	// 	newPokemon := createPokemon("ditto", typesData, movesData)

	// 	reqBodyBytes := new(bytes.Buffer)
	// 	json.NewEncoder(reqBodyBytes).Encode(newPokemon)
	// 	newPokemonBytes := reqBodyBytes.Bytes()
	// 	response := http.Response{Body: ioutil.NopCloser(bytes.NewBuffer(newPokemonBytes))}

	// 	return &response, nil
	// }
}

func TestCompareDamages(t *testing.T) {
	pDamageRelations := pokemonDamageRelations{}
	pokemon2 := Pokemon{}
	value := pDamageRelations.compareDamages(pokemon2, doubleDamageDealt)
	if value {
		t.Error("Expected false")
	}
}

func createPokemon(name string, types []pokemonTypeData, moves []MoveData) Pokemon {
	// Create types and set pokemon types.
	typesSize := len(types)
	pokemonTypes := make([]pokemonType, typesSize)
	for i, pType := range types {
		pokeType := pokemonType{
			Type: pType,
		}
		pokemonTypes[i] = pokeType
	}

	// Create moves and set pokemon moves.
	movesSize := len(moves)
	pokemonMoves := make([]pokemonMove, movesSize)
	for i, pMove := range moves {
		pokeMove := pokemonMove{
			Move: pMove,
		}
		pokemonMoves[i] = pokeMove
	}

	// Set pokemon name
	newPokemon := Pokemon{
		Name:  name,
		Types: pokemonTypes,
		Moves: pokemonMoves,
	}

	return newPokemon
}
