package main

import (
	"bufio"
	"fmt"
	"internal/pokeapi"
	"internal/pokecache"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(string) error
}

var quickMap map[string]cliCommand

// Creates a map of the possible CLI commands and their associated functions
func makeCommandMap() {
	quickMap = map[string]cliCommand{
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
			description: "Displays the names of the next 20 location areas in the Pokemon World",
			callback:    displayMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 location areas in the Pokemon World",
			callback:    displayMapBack,
		},
		"explore": {
			name:        "explore",
			description: "Displays the names of the pokemon in a specific area",
			callback:    exploreArea,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to catch a specfic pokemon",
			callback:    catchPokemon,
		},
		"inspect": {
			name:        "inspect",
			description: "Attempts to inspect a caught pokemon",
			callback:    inspectPokemon,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all captured pokemon in your pokedex",
			callback:    explorePokedex,
		},
	}
}

var _cached_storage pokecache.Cache
var _pokedex_storage map[string][]byte
var _quit_channel chan bool

// Entry point | runs main input loop
func main() {
	cached, quitChan := pokecache.NewCache(30)
	_pokedex_storage = make(map[string][]byte)
	_cached_storage = cached
	_quit_channel = quitChan
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	text := "pokedex > "
	fmt.Fprint(writer, text)
	writer.Flush()
	for scanner.Scan() {
		text += scanner.Text()
		parseCLiCommand(text)
		text = "pokedex > "
		fmt.Fprint(writer, text)
		writer.Flush()
	}
}

// Function to parse input from user and run command if found
// Commands are stored in the command map
// If command is not found, does nothing
// If command is found, calls the callback function | All callback functions don't require any arguments
func parseCLiCommand(input string) {
	input = strings.ToLower(input)
	if strings.Contains(input, "pokedex > ") {
		makeCommandMap()
		commandString, found := strings.CutPrefix(input, "pokedex > ")
		commandString, arguments, _ := strings.Cut(commandString, " ")
		if found {
			command, ok := quickMap[commandString]
			if !ok {
				fmt.Printf("Encountered error\n")
				return
			}
			fmt.Println("Executing")
			err := command.callback(arguments)
			if err != nil {
				fmt.Printf("Encountered error: %s", err)
				return
			}
		}
	}
}

// Prints the name of the command and its description held in the map
func commandHelp(arguments string) error {
	fmt.Printf("Welcome to the CLI Pokedex!\nUsage:\n\n")
	for msg := range quickMap {
		fmt.Printf("%v : %v\n", quickMap[msg].name, quickMap[msg].description)
	}
	return nil
}

// Exits the CLI application
func commandExit(arguments string) error {
	os.Exit(0)
	return nil

}

// Requests the next area of the map and displays it if found
// Moves the request forward if area has already been requested
func displayMap(arguments string) error {
	areas, err := pokeapi.GetAreaLocation(1, &_cached_storage)
	if err != nil {
		return err
	}
	for _, area := range areas {
		fmt.Printf("%v\n", area)
	}
	return nil
}

// Requests the previous area of the map and displays it if found
// Moves the request backward if area has already been requested
func displayMapBack(arguments string) error {
	areas, err := pokeapi.GetAreaLocation(-1, &_cached_storage)
	if err != nil {
		return err
	}
	for _, area := range areas {
		fmt.Printf("%v\n", area)
	}
	return nil
}
func exploreArea(arguments string) error {
	pokemon, err := pokeapi.GetPokemonInArea(arguments, &_cached_storage)
	if err != nil {
		return err
	}
	for _, mon := range pokemon {
		fmt.Printf("%v\n", mon)
	}
	return nil
}

func catchPokemon(pokemon string) error {
	success, err := pokeapi.CatchPokemon(pokemon, &_cached_storage, _pokedex_storage)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", success)
	return nil
}

func inspectPokemon(pokemon string) error {
	success, err := pokeapi.InspectPokemon(pokemon, _pokedex_storage)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", success)
	return nil
}

func explorePokedex(arguments string) error {
	success, err := pokeapi.ExplorePokedex(_pokedex_storage)
	if err != nil {
		return err
	}
	fmt.Println(success)
	return nil
}
