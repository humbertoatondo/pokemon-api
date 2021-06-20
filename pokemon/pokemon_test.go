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

	t.Run("Get pokemon with one type", func(t *testing.T) {
		pokemonName := "ditto"
		pokemon, _ := GetPokemon(pokemonName, "", httpGetPokemon)

		basePokemonName := "ditto"
		typesData := []pokemonTypeData{
			pokemonTypeData{Name: "normal", URL: "https://pokeapi.co/api/v2/type/1/"},
		}
		movesData := []MoveData{
			MoveData{Name: "transform", URL: "https://pokeapi.co/api/v2/move/144/"},
		}

		basePokemon := createPokemon(basePokemonName, typesData, movesData)

		assert.Equal(t, basePokemon, pokemon)
	})

	t.Run("Get pokemon with two types", func(t *testing.T) {
		pokemonName := "lucario"
		pokemon, _ := GetPokemon(pokemonName, "", httpGetPokemon)

		basePokemonName := "lucario"
		typesData := []pokemonTypeData{
			pokemonTypeData{Name: "fighting", URL: "https://pokeapi.co/api/v2/type/2/"},
			pokemonTypeData{Name: "steel", URL: "https://pokeapi.co/api/v2/type/9/"},
		}

		movesData := []MoveData{
			MoveData{Name: "ice-punch", URL: "https://pokeapi.co/api/v2/move/8/"},
			MoveData{Name: "thunder-punch", URL: "https://pokeapi.co/api/v2/move/9/"},
			MoveData{Name: "swords-dance", URL: "https://pokeapi.co/api/v2/move/14/"},
			MoveData{Name: "headbutt", URL: "https://pokeapi.co/api/v2/move/29/"},
			MoveData{Name: "roar", URL: "https://pokeapi.co/api/v2/move/46/"},
		}

		basePokemon := createPokemon(basePokemonName, typesData, movesData)

		assert.Equal(t, basePokemon, pokemon)
	})

	t.Run("Pokemons should be different", func(t *testing.T) {
		pokemonName := "lucario"
		pokemon, _ := GetPokemon(pokemonName, "", httpGetPokemon)

		basePokemonName := "ditto"
		typesData := []pokemonTypeData{
			pokemonTypeData{Name: "normal", URL: "https://pokeapi.co/api/v2/type/1/"},
		}
		movesData := []MoveData{
			MoveData{Name: "transform", URL: "https://pokeapi.co/api/v2/move/144/"},
		}

		basePokemon := createPokemon(basePokemonName, typesData, movesData)

		assert.NotEqual(t, basePokemon, pokemon)
	})

}

func TestCompareTo(t *testing.T) {

}

func TestCompareDamages(t *testing.T) {

	t.Run("Pikachu can't deal double damage to lucario", func(t *testing.T) {
		response, _ := httpGetPokemonDamageRelations("pikachu")
		var pDamageRelations = pokemonDamageRelations{}
		if err := json.NewDecoder(response.Body).Decode(&pDamageRelations); err != nil {
			t.Error("Error decoding response from 'httpGetPokemonDamageRelations' request.")
		}

		response, _ = httpGetPokemon("lucario")
		var pokemon = Pokemon{}
		if err := json.NewDecoder(response.Body).Decode(&pokemon); err != nil {
			t.Error("Error decoding response from 'httpGetPokemon' request.")
		}

		res := pDamageRelations.compareDamages(pokemon, doubleDamageDealt)

		assert.False(t, res)
	})

	t.Run("Pikachu can receive half damage from lucario", func(t *testing.T) {
		response, _ := httpGetPokemonDamageRelations("pikachu")
		var pDamageRelations = pokemonDamageRelations{}
		if err := json.NewDecoder(response.Body).Decode(&pDamageRelations); err != nil {
			t.Error("Error decoding response from 'httpGetPokemonDamageRelations' request.")
		}

		response, _ = httpGetPokemon("lucario")
		var pokemon = Pokemon{}
		if err := json.NewDecoder(response.Body).Decode(&pokemon); err != nil {
			t.Error("Error decoding response from 'httpGetPokemon' request.")
		}

		res := pDamageRelations.compareDamages(pokemon, halfDamageReceived)

		assert.True(t, res)
	})

	t.Run("Pikachu can't receive no damage from lucario", func(t *testing.T) {
		response, _ := httpGetPokemonDamageRelations("pikachu")
		var pDamageRelations = pokemonDamageRelations{}
		if err := json.NewDecoder(response.Body).Decode(&pDamageRelations); err != nil {
			t.Error("Error decoding response from 'httpGetPokemonDamageRelations' request.")
		}

		response, _ = httpGetPokemon("lucario")
		var pokemon = Pokemon{}
		if err := json.NewDecoder(response.Body).Decode(&pokemon); err != nil {
			t.Error("Error decoding response from 'httpGetPokemon' request.")
		}

		res := pDamageRelations.compareDamages(pokemon, doubleDamageDealt)

		assert.False(t, res)
	})
}

var httpGetPokemon = func(pokemonName string) (*http.Response, error) {
	var typesData []pokemonTypeData
	var movesData []MoveData

	switch pokemonName {
	case "ditto":
		typesData = []pokemonTypeData{
			pokemonTypeData{Name: "normal", URL: "https://pokeapi.co/api/v2/type/1/"},
		}

		movesData = []MoveData{
			MoveData{Name: "transform", URL: "https://pokeapi.co/api/v2/move/144/"},
		}

	case "lucario":
		typesData = []pokemonTypeData{
			pokemonTypeData{Name: "fighting", URL: "https://pokeapi.co/api/v2/type/2/"},
			pokemonTypeData{Name: "steel", URL: "https://pokeapi.co/api/v2/type/9/"},
		}

		movesData = []MoveData{
			MoveData{Name: "ice-punch", URL: "https://pokeapi.co/api/v2/move/8/"},
			MoveData{Name: "thunder-punch", URL: "https://pokeapi.co/api/v2/move/9/"},
			MoveData{Name: "swords-dance", URL: "https://pokeapi.co/api/v2/move/14/"},
			MoveData{Name: "headbutt", URL: "https://pokeapi.co/api/v2/move/29/"},
			MoveData{Name: "roar", URL: "https://pokeapi.co/api/v2/move/46/"},
		}
	}

	newPokemon := createPokemon(pokemonName, typesData, movesData)

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(newPokemon)
	newPokemonBytes := reqBodyBytes.Bytes()
	response := http.Response{Body: ioutil.NopCloser(bytes.NewBuffer(newPokemonBytes))}

	return &response, nil
}

var httpGetPokemonDamageRelations = func(pokemonName string) (*http.Response, error) {

	// Pikachu damage relations
	doubleDamageToList := []damageTypeName{
		damageTypeName{Type: "flying"},
		damageTypeName{Type: "water"},
	}

	halfDamageFromList := []damageTypeName{
		damageTypeName{Type: "flying"},
		damageTypeName{Type: "steel"},
		damageTypeName{Type: "electric"},
	}

	noDamageFromList := []damageTypeName{}

	pDamageRelations := pokemonDamageRelations{
		DamageRelations: damageRelations{
			DoubleDamageToList: doubleDamageToList,
			HalfDamageFromList: halfDamageFromList,
			NoDamageFromList:   noDamageFromList,
		},
	}

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(pDamageRelations)
	newDamageRelationsBytes := reqBodyBytes.Bytes()
	response := http.Response{Body: ioutil.NopCloser(bytes.NewBuffer(newDamageRelationsBytes))}

	return &response, nil
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
