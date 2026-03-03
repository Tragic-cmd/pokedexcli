package main

import (
	"fmt"
	"errors"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
        return errors.New("no pokemon to inspect.")
    }

	pokemonResp, ok := cfg.caughtPokemon[args[0]]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}

	// ... inside a function ...
	fmt.Printf("Name: %s\n", pokemonResp.Name)
	fmt.Printf("Height: %v\n", pokemonResp.Height)
	fmt.Printf("Weight: %v\n", pokemonResp.Weight)
	fmt.Println("Stats:")
	for _, s := range pokemonResp.Stats {
		fmt.Printf("  -%s: %d\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemonResp.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}
	return nil
}