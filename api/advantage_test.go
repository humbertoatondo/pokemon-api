package api

import (
	"os"
	"testing"
)

var app App

func TestMain(m *testing.M) {

	code := m.Run()
	os.Exit(code)
}

// func TestComparePokemons(t *testing.T) {

// }
