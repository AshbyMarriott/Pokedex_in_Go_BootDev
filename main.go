package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/AshbyMarriott/Pokedex_in_Go_BootDev/cli"
)

// main is the entry point of the Pokedex application.
// It starts an interactive REPL (Read-Eval-Print Loop) that reads user input,
// processes commands, and executes them until the user exits.
func main() {
	pokeScanner := bufio.NewScanner(os.Stdin)
	for true {
		fmt.Print("Pokedex >")
		pokeScanner.Scan()
		text := pokeScanner.Text()
		cleaned := cli.CleanInput(text)
		if len(cleaned) == 0 {
			fmt.Println("No command found, enter a command")
		} else {
			cli.CallCommand(cleaned[0])
		}
	}
}
