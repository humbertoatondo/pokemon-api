package pokemon

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Pokemon stores the name of a pokemon and its type.
type Pokemon struct {
	Name  string        `json:"name"`
	Types []pokemonType `json:"types"`
	Moves []pokemonMove `json:"moves"`
}

type pokemonType struct {
	Slot int             `json:"slot"`
	Type pokemonTypeData `json:"type"`
}

type pokemonTypeData struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type pokemonMove struct {
	Move pokemonMoveData `json:"move"`
}

type pokemonMoveData struct {
	Name string `json:"name"`
}

// CompareResults stores boolean values to indicate if a certain pokemon
// can deal dobule damage, receive half damage or no damage at all.
type CompareResults struct {
	DealsDoubleDamage  bool `json:"deals_double_damage"`
	ReceivesHalfDamage bool `json:"receives_half_damage"`
	ReceivesNoDamage   bool `json:"receives_no_damage"`
}

type pokemonDamageRelations struct {
	DamageRelations damageRelations `json:"damage_relations"`
}

type damageRelations struct {
	DoubleDamageToList []damageTypeName `json:"double_damage_to"`
	HalfDamageFromList []damageTypeName `json:"half_damage_from"`
	NoDamageFromList   []damageTypeName `json:"no_damage_from"`
}

type damageTypeName struct {
	Type string `json:"name"`
}

type damageType int

const (
	doubleDamageDealt damageType = iota
	halfDamageReceived
	noDamageReceived
)

// GetPokemon receives a pokemon name and makes an http request
// to get that pokemon's data from the api pokeapi.
func GetPokemon(pokemonName string) (Pokemon, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemonName)
	response, err := http.Get(url)
	if err != nil {
		return Pokemon{}, err
	}

	defer response.Body.Close()

	var pokemon = Pokemon{}
	if err = json.NewDecoder(response.Body).Decode(&pokemon); err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}

// CompareTo receives a rival pokemon and compares it with the current
// pokemon in the following categories based on their types:
//   - Current pokemon can deal double damage to rival pokemon.
//   - Current pokemon can receive half damage from rival pokemon.
//   - Current pokemon can receive no damage from rival pokemon.
func (pokemon *Pokemon) CompareTo(rivalPokemon Pokemon) (CompareResults, error) {
	var compareResults = CompareResults{}

	for _, pType := range pokemon.Types {
		url := pType.Type.URL

		response, err := http.Get(url)
		if err != nil {
			return CompareResults{}, err
		}

		var pokemonDamageRelations = pokemonDamageRelations{}
		if err = json.NewDecoder(response.Body).Decode(&pokemonDamageRelations); err != nil {
			return CompareResults{}, err
		}

		dealsDoubleDamage := pokemonDamageRelations.compareDamages(rivalPokemon, doubleDamageDealt)
		receivesHalfDamage := pokemonDamageRelations.compareDamages(rivalPokemon, halfDamageReceived)
		receivesNoDamage := pokemonDamageRelations.compareDamages(rivalPokemon, noDamageReceived)

		compareResults.DealsDoubleDamage = compareResults.DealsDoubleDamage || dealsDoubleDamage
		compareResults.ReceivesHalfDamage = compareResults.ReceivesHalfDamage || receivesHalfDamage
		compareResults.ReceivesNoDamage = compareResults.ReceivesNoDamage || receivesNoDamage
	}

	return compareResults, nil
}

// compareDamages is a function for PokemonDamageRelations and receives a rival Pokemon and a damageType
// and compares both pokemons depending on what we want to compare.
// For example:
//   - If dType is set to doubleDamageDealt then we will return true if the current pokemon
//     can deal double damage to the rival pokemon, else return false.
func (pokemonDamageRelations *pokemonDamageRelations) compareDamages(rivalPokemon Pokemon, dType damageType) bool {
	var rivalPokemonTypeList = rivalPokemon.Types
	var damageTypeNameList []damageTypeName

	switch dType {
	case doubleDamageDealt:
		damageTypeNameList = pokemonDamageRelations.DamageRelations.DoubleDamageToList
		break
	case halfDamageReceived:
		damageTypeNameList = pokemonDamageRelations.DamageRelations.HalfDamageFromList
		break
	case noDamageReceived:
		damageTypeNameList = pokemonDamageRelations.DamageRelations.NoDamageFromList
		break
	}

	for _, damage := range damageTypeNameList {
		for _, rivalPokemonType := range rivalPokemonTypeList {
			if damage.Type == rivalPokemonType.Type.Name {
				return true
			}
		}
	}
	return false
}

// GetPokemonsFromListOfNames receives a list with pokemon names and calls
// the function GetPokemon for every pokemon name in the list to get the
// pokemon's data.
func GetPokemonsFromListOfNames(pokemonNames []string) ([]Pokemon, error) {
	size := len(pokemonNames)
	pokemons := make([]Pokemon, size)

	for i, pokemonName := range pokemonNames {
		pokemon, err := GetPokemon(pokemonName)
		if err != nil {
			return make([]Pokemon, 0), err
		}
		pokemons[i] = pokemon
	}

	return pokemons, nil
}

// GetCommonMovesForPokemons receives a list of pokemons and returns
// a list with all the common moves between this pokemons.
func GetCommonMovesForPokemons(pokemons []Pokemon) []string {
	var commonMoves []string
	movesMap := make(map[string]int)

	// Build commonMoves map
	for i, pokemon := range pokemons {
		for _, pMove := range pokemon.Moves {
			pokemonName := pMove.Move.Name
			_, ok := movesMap[pokemonName]
			if i > 0 && !ok {
				continue
			} else {
				movesMap[pokemonName]++
			}
		}
	}

	// Get common moves from map
	size := len(pokemons)
	for key, value := range movesMap {
		if value == size {
			commonMoves = append(commonMoves, key)
		}
	}

	return commonMoves
}
