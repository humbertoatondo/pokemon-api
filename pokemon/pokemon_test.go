package pokemon

import (
	"testing"
)

func TestCompareDamages(t *testing.T) {
	pDamageRelations := pokemonDamageRelations{}
	pokemon2 := Pokemon{}
	value := pDamageRelations.compareDamages(pokemon2, doubleDamageDealt)
	if value {
		t.Error("Expected false")
	}
}
