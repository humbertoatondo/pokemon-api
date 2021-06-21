// +build integration

package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommonMoves(t *testing.T) {

	t.Run("Comparing pikachu's and lucario's moves should return an array with 2 moves", func(t *testing.T) {
		urlPath := "http://localhost:5000/comparePokemonsMoves?pokemon=pikachu&pokemon=lucario"
		response, _ := http.Get(urlPath)

		var resBody commonMoves
		if err := json.NewDecoder(response.Body).Decode(&resBody); err != nil {
			t.Error(err)
		}

		jsonFilePath := "../testing_resources/common_moves/pikachu_lucario_en.json"
		fileBytes, err := ioutil.ReadFile(jsonFilePath)
		if err != nil {
			t.Error(err)
		}

		var expectedBody commonMoves
		json.Unmarshal(fileBytes, &expectedBody)

		assert.ElementsMatch(t, resBody.Pokemons, expectedBody.Pokemons)
		assert.ElementsMatch(t, resBody.Moves, expectedBody.Moves)
	})

	t.Run("Comparing ditto's and lucario's moves should return an empty array of moves", func(t *testing.T) {
		urlPath := "http://localhost:5000/comparePokemonsMoves?pokemon=lucario&pokemon=ditto"
		response, _ := http.Get(urlPath)

		var resBody commonMoves
		if err := json.NewDecoder(response.Body).Decode(&resBody); err != nil {
			t.Error(err)
		}

		jsonFilePath := "../testing_resources/common_moves/lucario_ditto_en.json"
		fileBytes, err := ioutil.ReadFile(jsonFilePath)
		if err != nil {
			t.Error(err)
		}

		var expectedBody commonMoves
		json.Unmarshal(fileBytes, &expectedBody)

		assert.ElementsMatch(t, resBody.Pokemons, expectedBody.Pokemons)
		assert.ElementsMatch(t, resBody.Moves, expectedBody.Moves)
	})

	t.Run("Comparing pikachu's and lucario's moves should return an array with 2 moves in spanish", func(t *testing.T) {
		urlPath := "http://localhost:5000/comparePokemonsMoves?pokemon=lucario&pokemon=pikachu&lang=es"
		response, _ := http.Get(urlPath)

		var resBody commonMoves
		if err := json.NewDecoder(response.Body).Decode(&resBody); err != nil {
			t.Error(err)
		}

		jsonFilePath := "../testing_resources/common_moves/lucario_pikachu_es.json"
		fileBytes, err := ioutil.ReadFile(jsonFilePath)
		if err != nil {
			t.Error(err)
		}

		var expectedBody commonMoves
		json.Unmarshal(fileBytes, &expectedBody)

		assert.ElementsMatch(t, resBody.Pokemons, expectedBody.Pokemons)
		assert.ElementsMatch(t, resBody.Moves, expectedBody.Moves)
	})

}
