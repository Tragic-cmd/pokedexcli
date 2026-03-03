package main

import (
	"fmt"
	"errors"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) < 1 {
        return errors.New("You have not caught any pokemon!")
    }

	fmt.Println("Your Pokedex:")

	for _, p := range cfg.caughtPokemon {
		fmt.Printf("  - %s\n", p.Name)
	}

	return nil
}