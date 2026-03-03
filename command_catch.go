package main

import (
	"fmt"
	"errors"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
        return errors.New("No pokemon to catch.")
    }

	pokemonName := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	pokemonResp, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	if pokemonResp.BaseExperience <= 0 {
		fmt.Printf("%s was caught!\n", pokemonResp.Name)
		fmt.Println("You may now inspect it with the inspect command.")
		cfg.caughtPokemon[pokemonResp.Name] = pokemonResp
		return nil
	}

	roll := rand.Intn(pokemonResp.BaseExperience)

	if roll > 40 {
		fmt.Printf("%s escaped!\n", pokemonResp.Name)
		return nil
	} else {
		fmt.Printf("%s was caught!\n", pokemonResp.Name)
		cfg.caughtPokemon[pokemonResp.Name] = pokemonResp
		return nil
	}
}