// Package cli provides a command-line interface for the Pokedex application.
// It handles user input, command execution, and provides a help message.
// Author: Ashby Marriott
package cli

import (
	"fmt"
	"os"
	"strings"
)

// cliCommand represents a command that can be executed in the CLI.
// It contains the command name, description, and a callback function to execute.
type cliCommand struct {
	name        string
	description string
	callback    func() error
}

// CommandRegistry is a map of command names to their corresponding cliCommand definitions.
// It is initialized in the init function with available commands.
var CommandRegistry map[string]cliCommand

// init initializes the CommandRegistry with all available CLI commands.
func init() {
	CommandRegistry = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}

// CleanInput normalizes and splits the input text by whitespace.
// It converts the text to lowercase and splits it into words, removing any
// leading, trailing, or consecutive whitespace characters.
func CleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	splitText := strings.Fields(lowerText)
	return splitText
}

// CallCommand executes the specified command if it exists in the CommandRegistry.
// If the command is not found, it prints an "Unknown Command" message.
func CallCommand(command string) {
	if _, ok := CommandRegistry[command]; !ok {
		fmt.Println("Unknown Command")
	} else {
		CommandRegistry[command].callback()
	}
}

// commandExit terminates the program with a goodbye message.
// It prints a closing message and exits with status code 0.
func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

// commandHelp displays a help message listing all available commands.
// It iterates through the CommandRegistry and prints each command's name and description.
func commandHelp() error {
	helpString := "Welcome to the Pokedex!\nUsage:\n\n"
	for _, command := range CommandRegistry {
		helpString += fmt.Sprintf("%s: %s\n", command.name, command.description)
	}
	fmt.Print(helpString)
	return nil
}
