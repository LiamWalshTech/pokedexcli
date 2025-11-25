package main

import (
	"bufio"
	"fmt"
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
		fmt.Printf("Your command was: %v\n", usersInputCleanWords[0])
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
