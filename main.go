package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		usersInput := scanner.Text()
		usersInputClean := strings.ToLower(usersInput)
		usersInputCleanWords := strings.Fields(usersInputClean)
		command, exists := getCommands()[usersInputCleanWords[0]]
		cfg := &config{}

		if exists {
			err := command.callback(cfg)

			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandMap(cfg *config) error {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/")

	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return err
	}

	if err != nil {
		return err
	}

	data, err := json.Unmarshal(body)

	fmt.Printf("%s", data)

	return nil
}

type config struct {
    nextURL     *string
    previousURL *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "It displays the names of 20 location areas in the Pokemon world.",
			callback:    commandMap,
		},
	}
}
