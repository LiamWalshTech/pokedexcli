package main

import (
	"fmt"
)

func commandInspect(cfg *config, argument string) error {
	if pokemon, ok := cfg.caughtPokemon[argument]; ok {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %v\n", pokemon.Height)
		fmt.Printf("Weight: %v\n", pokemon.Weight)
		fmt.Printf("Stats:\n")
		for _, pokemonStat := range pokemon.Stats {
			fmt.Printf("  -%s: %v\n", pokemonStat.Stat.Name, pokemonStat.BaseStat)
		}
		fmt.Printf("Types:\n")
		for _, pokemonType := range pokemon.Types {
			fmt.Printf("  -%s\n", pokemonType.Type.Name)
		}
	} else {
		fmt.Println("you have not caught that pokemon")
	}

	return nil
}
