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

func TestCompareDamages(t *testing.T) {

	t.Run("Pikachu can't deal double damage to lucario", func(t *testing.T) {
		response, _ := httpGetPokemonDamageRelations("pikachu")
		var pDamageRelations = PDamageRelations{}
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
		var pDamageRelations = PDamageRelations{}
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
		var pDamageRelations = PDamageRelations{}
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

func TestGetCommonMovesForPokemons(t *testing.T) {
	response, _ := httpGetPokemon("lucario")
	var pokemon1 = Pokemon{}
	if err := json.NewDecoder(response.Body).Decode(&pokemon1); err != nil {
		t.Error("Error decoding response from 'httpGetPokemon' request.")
	}

	t.Run("Pokemon List containig one pokemon should return all the moves from that pokemon", func(t *testing.T) {
		pokemonList := []Pokemon{pokemon1}
		commonMoves := GetCommonMovesForPokemons(pokemonList, 10)

		movesSize := len(pokemon1.Moves)
		pokemon1Moves := make([]MoveData, movesSize)
		for i, pMove := range pokemon1.Moves {
			pokemon1Moves[i] = pMove.Move
		}

		assert.ElementsMatch(t, commonMoves, pokemon1Moves)
	})

	response, _ = httpGetPokemon("ditto")
	var pokemon2 = Pokemon{}
	if err := json.NewDecoder(response.Body).Decode(&pokemon2); err != nil {
		t.Error("Error decoding response from 'httpGetPokemon' request.")
	}

	t.Run("Common moves from lucario and ditto should return an empty list", func(t *testing.T) {
		pokemonList := []Pokemon{pokemon1, pokemon2}
		commonMoves := GetCommonMovesForPokemons(pokemonList, 10)

		pokemonMoves := make([]MoveData, 0)

		assert.ElementsMatch(t, commonMoves, pokemonMoves)
	})

	response, _ = httpGetPokemon("pikachu")
	var pokemon3 = Pokemon{}
	if err := json.NewDecoder(response.Body).Decode(&pokemon3); err != nil {
		t.Error("Error decoding response from 'httpGetPokemon' request.")
	}

	t.Run("fdsgsfd", func(t *testing.T) {
		pokemonList := []Pokemon{pokemon1, pokemon3}
		commonMoves := GetCommonMovesForPokemons(pokemonList, 10)

		pokemonMoves := []MoveData{
			MoveData{Name: "thunder-punch", URL: "https://pokeapi.co/api/v2/move/9/"},
			MoveData{Name: "headbutt", URL: "https://pokeapi.co/api/v2/move/29/"},
		}

		assert.ElementsMatch(t, commonMoves, pokemonMoves)
	})

}

func TestTranslatePokemonMoves(t *testing.T) {

	t.Run("Empty list of pokemon moves should return an empty list", func(t *testing.T) {
		pokemonMoves := []MoveData{}
		translatedMoves, _ := TranslatePokemonMoves(pokemonMoves, "en", httpTranslateMove)

		assert.ElementsMatch(t, translatedMoves, pokemonMoves)
	})

	t.Run("Translation of [thunder-punch] to spanish must return [Puño Trueno]", func(t *testing.T) {
		pokemonMoves := []MoveData{
			MoveData{Name: "thunder-punch", URL: "thunder-punch"},
		}

		translatedMoves, _ := TranslatePokemonMoves(pokemonMoves, "es", httpTranslateMove)

		expectedMoves := []MoveData{
			MoveData{Name: "Puño Trueno", URL: "thunder-punch"},
		}

		assert.ElementsMatch(t, translatedMoves, expectedMoves)
	})

	t.Run("Translation of [thunder-punch] to japanese must return [かみなりパンチ]", func(t *testing.T) {
		pokemonMoves := []MoveData{
			MoveData{Name: "thunder-punch", URL: "thunder-punch"},
		}

		translatedMoves, _ := TranslatePokemonMoves(pokemonMoves, "ja", httpTranslateMove)

		expectedMoves := []MoveData{
			MoveData{Name: "かみなりパンチ", URL: "thunder-punch"},
		}

		assert.ElementsMatch(t, translatedMoves, expectedMoves)
	})

}

// HELPER VARIABLES AND FUNCTIONS ==================================================

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

	case "pikachu":
		typesData = []pokemonTypeData{
			pokemonTypeData{Name: "electric", URL: "https://pokeapi.co/api/v2/type/13/"},
		}

		movesData = []MoveData{
			MoveData{Name: "pay-day", URL: "https://pokeapi.co/api/v2/move/6/"},
			MoveData{Name: "thunder-punch", URL: "https://pokeapi.co/api/v2/move/9/"},
			MoveData{Name: "slam", URL: "https://pokeapi.co/api/v2/move/21/"},
			MoveData{Name: "mega-kick", URL: "https://pokeapi.co/api/v2/move/25/"},
			MoveData{Name: "headbutt", URL: "https://pokeapi.co/api/v2/move/29/"},
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

	pDamageRelations := PDamageRelations{
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

var httpTranslateMove = func(pokeMove string) (*http.Response, error) {
	var tMoves TransMoves

	switch pokeMove {
	case "thunder-punch":
		tMoves = TransMoves{
			Names: []MoveData{
				MoveData{Name: "かみなりパンチ", URL: "https://pokeapi.co/api/v2/move/9/"},
				MoveData{Name: "번개펀치", URL: "https://pokeapi.co/api/v2/move/9/"},
				MoveData{Name: "雷電拳", URL: "https://pokeapi.co/api/v2/move/9/"},
				MoveData{Name: "Poing Éclair", URL: "https://pokeapi.co/api/v2/move/9/"},
				MoveData{Name: "Donnerschlag", URL: "https://pokeapi.co/api/v2/move/9/"},
				MoveData{Name: "Puño Trueno", URL: "https://pokeapi.co/api/v2/move/9/"},
				MoveData{Name: "Tuonopugno", URL: "https://pokeapi.co/api/v2/move/9/"},
				MoveData{Name: "Thunder Punch", URL: "https://pokeapi.co/api/v2/move/9/"},
				MoveData{Name: "かみなりパンチ", URL: "https://pokeapi.co/api/v2/move/9/"},
				MoveData{Name: "雷电拳", URL: "https://pokeapi.co/api/v2/move/9/"},
			},
		}
	}

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(tMoves)
	newtMovesBytes := reqBodyBytes.Bytes()
	response := http.Response{Body: ioutil.NopCloser(bytes.NewBuffer(newtMovesBytes))}

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
	pokemonMoves := make([]PokemonMove, movesSize)
	for i, pMove := range moves {
		pokeMove := PokemonMove{
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
