// +build integration

package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComparePokemons(t *testing.T) {

	t.Run("Comparing pikachu with lucario should return [false, true, false]", func(t *testing.T) {

		urlPath := "http://localhost:5000/comparePokemons?pokemon1=pikachu&pokemon2=lucario"
		response, _ := http.Get(urlPath)

		var resBody advantage
		if err := json.NewDecoder(response.Body).Decode(&resBody); err != nil {
			t.Error(err)
		}

		jsonFilePath := "../testing_resources/advantages/pikachu_vs_lucario.json"
		fileBytes, err := ioutil.ReadFile(jsonFilePath)
		if err != nil {
			t.Error(err)
		}

		var expectedBody advantage
		json.Unmarshal(fileBytes, &expectedBody)

		assert.Equal(t, resBody.ComparisonResults.DealsDoubleDamage, expectedBody.ComparisonResults.DealsDoubleDamage)
		assert.Equal(t, resBody.ComparisonResults.ReceivesHalfDamage, expectedBody.ComparisonResults.ReceivesHalfDamage)
		assert.Equal(t, resBody.ComparisonResults.ReceivesNoDamage, expectedBody.ComparisonResults.ReceivesNoDamage)
	})

	t.Run("Comparing lucario with lucario should return [true, true, false]", func(t *testing.T) {

		urlPath := "http://localhost:5000/comparePokemons?pokemon1=lucario&pokemon2=lucario"
		response, _ := http.Get(urlPath)

		var resBody advantage
		if err := json.NewDecoder(response.Body).Decode(&resBody); err != nil {
			t.Error(err)
		}

		jsonFilePath := "../testing_resources/advantages/lucario_vs_lucario.json"
		fileBytes, err := ioutil.ReadFile(jsonFilePath)
		if err != nil {
			t.Error(err)
		}

		var expectedBody advantage
		json.Unmarshal(fileBytes, &expectedBody)

		assert.Equal(t, resBody.ComparisonResults.DealsDoubleDamage, expectedBody.ComparisonResults.DealsDoubleDamage)
		assert.Equal(t, resBody.ComparisonResults.ReceivesHalfDamage, expectedBody.ComparisonResults.ReceivesHalfDamage)
		assert.Equal(t, resBody.ComparisonResults.ReceivesNoDamage, expectedBody.ComparisonResults.ReceivesNoDamage)
	})

	t.Run("Comparing pikachu with ditto should return [false, false, false]", func(t *testing.T) {

		urlPath := "http://localhost:5000/comparePokemons?pokemon1=pikachu&pokemon2=ditto"
		response, _ := http.Get(urlPath)

		var resBody advantage
		if err := json.NewDecoder(response.Body).Decode(&resBody); err != nil {
			t.Error(err)
		}

		jsonFilePath := "../testing_resources/advantages/pikachu_vs_ditto.json"
		fileBytes, err := ioutil.ReadFile(jsonFilePath)
		if err != nil {
			t.Error(err)
		}

		var expectedBody advantage
		json.Unmarshal(fileBytes, &expectedBody)

		assert.Equal(t, resBody.ComparisonResults.DealsDoubleDamage, expectedBody.ComparisonResults.DealsDoubleDamage)
		assert.Equal(t, resBody.ComparisonResults.ReceivesHalfDamage, expectedBody.ComparisonResults.ReceivesHalfDamage)
		assert.Equal(t, resBody.ComparisonResults.ReceivesNoDamage, expectedBody.ComparisonResults.ReceivesNoDamage)
	})

}
