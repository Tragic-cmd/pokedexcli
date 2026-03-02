package main

import (
	"fmt"
	"errors"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
        return errors.New("No area to explore.")
    }
	areaName := args[0]
	areaResp, err := cfg.pokeapiClient.ExploreLocation(areaName)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", areaName)
	fmt.Println("Found Pokemon:")
	for _, poke := range areaResp.Pokemon_encounters {
		fmt.Printf(" - %s\n", poke.Pokemon.Name)
	}
	return nil
}