package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var app App

func TestMain(m *testing.M) {

	code := m.Run()
	os.Exit(code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	app.Router.ServeHTTP(rr, req)

	return rr
}

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

		// fmt.Printf("%v\n", expectedBody)

		assert.Equal(t, resBody, expectedBody)
	})

}
