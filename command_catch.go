package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, argument string) error {
	pokemonResp, err := cfg.pokeapiClient.PokemonGet(argument)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", argument)

	random := rand.Intn(pokemonResp.BaseExperience)
	threshold := 20

	if random < threshold {
		cfg.caughtPokemon[pokemonResp.Name] = pokemonResp
		fmt.Printf("%s was caught!\n", argument)
	} else {

		fmt.Printf("%s escaped!\n", argument)
	}

	return nil
}
