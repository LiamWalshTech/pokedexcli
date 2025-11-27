package main

import (
	"time"

	"github.com/LiamWalshTech/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		caughtPokemon: map[string]pokeapi.RespShallowPokemon{},
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
