package main

import (
	"fmt"
)

func commandExplore(cfg *config, argument string) error {
	locationResp, err := cfg.pokeapiClient.LocationGet(argument)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", argument)
	fmt.Println("Found Pokemon:")

	for _, encounter := range locationResp.PokemonEncounters {
		fmt.Println(" - " + encounter.Pokemon.Name)
	}
	return nil
}
