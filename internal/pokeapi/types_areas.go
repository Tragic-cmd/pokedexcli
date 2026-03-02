package pokeapi

// RespArea -
type RespArea struct {
    Pokemon_encounters []struct {
        Pokemon struct {
			Name string `json:"name"`
			Url string `json:"url"`
    	} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}
