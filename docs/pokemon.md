package pokemon // import "github.com/humbertoatondo/pokemon-api/pokemon"


TYPES

type CompareResults struct {
        DealsDoubleDamage  bool `json:"deals_double_damage"`
        ReceivesHalfDamage bool `json:"receives_half_damage"`
        ReceivesNoDamage   bool `json:"receives_no_damage"`
}
    CompareResults stores boolean values to indicate if a certain pokemon can
    deal dobule damage, receive half damage or no damage at all.

type MoveData struct {
        Name string `json:"name"`
        URL  string `json:"url"`
}
    MoveData stores the name and the pokeapi url of a pokemon move.

func GetCommonMovesForPokemons(pokemons []Pokemon, limit int) []MoveData
    GetCommonMovesForPokemons receives a list of pokemons and returns a list
    with all the common moves between this pokemons.

func TranslatePokemonMoves(pokemonMoves []MoveData, lang string, httpGet helpers.HTTPGet) ([]MoveData, error)
    TranslatePokemonMoves receives a list of pokemon moves and a language and
    translate every move to the desired language.

type PDamageRelations struct {
        DamageRelations damageRelations `json:"damage_relations"`
}
    Stores the results obtained by the different damage realtions.

type Pokemon struct {
        Name  string        `json:"name"`
        Types []pokemonType `json:"types"`
        Moves []PokemonMove `json:"moves"`
}
    Pokemon stores the name of a pokemon, its types and its moves.

func GetPokemon(pokemonName string, pokemonURL string, httpGet helpers.HTTPGet) (Pokemon, error)
    GetPokemon receives a pokemon name and makes an http request to get that
    pokemon's data from the api pokeapi.

func GetPokemonsFromListOfNames(pokemonNames []string, pokemonURL string, httpGet helpers.HTTPGet) ([]Pokemon, error)
    GetPokemonsFromListOfNames receives a list with pokemon names and calls the
    function GetPokemon for every pokemon name in the list to get the pokemon's
    data.

func (pokemon *Pokemon) CompareTo(rivalPokemon Pokemon, httpGet helpers.HTTPGet) (CompareResults, error)
    CompareTo receives a rival pokemon and compares it with the current pokemon
    in the following categories based on their types:

        - Current pokemon can deal double damage to rival pokemon.
        - Current pokemon can receive half damage from rival pokemon.
        - Current pokemon can receive no damage from rival pokemon.

type PokemonMove struct {
        Move MoveData `json:"move"`
}
    Stores the name of a pokemon move

type TransMoves struct {
        Names []MoveData `json:"names"`
}
    Stores the transalated name of a pokemon move.