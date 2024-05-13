package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func startRepl(cfg *config){
	reader := bufio.NewScanner(os.Stdin)
	for{

		fmt.Print("Pokedex > ")
		reader.Scan()
		words := cleanInput(reader.Text())
		if len(words) ==0{
			continue
		}
		commandName := words[0]
		command,exists := getCommands()[commandName]

		if exists{
			err:= command.callback(cfg)
			if err != nil{
				fmt.Println(err)
			}
			continue
		}else{
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string)[]string{
	output:= strings.ToLower(text)
	words := strings.Fields(output)
	return words
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
		"date":{
			name: "date",
			description: "Get current date",
			callback: commandDate,
		},
		"map":{
			name: "map",
			description: "Get 20 location name areas, display next in subsequent calls",
			callback: commandMap,
		},
		"mapb":{
			name: "mapb",
			description: "Similar to map but instead get previous 20 locations",
			callback: commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
