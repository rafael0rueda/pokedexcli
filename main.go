package main

import ( 
	"fmt"
	"strings"
	"bufio"
	"os"
)

type cliCommand struct {
	name string
	description string
	callback func() error
}

var listCommand map[string]cliCommand

func cleanInput(text string) []string { 
	lowerText := strings.ToLower(text)
	words := strings.Fields(lowerText)
	//fmt.Println(words)
	return words
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	for _, value := range listCommand {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}
	return nil
}

func init() {
	listCommand = map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help":{
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
	}
}

func main(){
	fmt.Println("Welcome to the Pokedex!")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Usage: ")
		scanner.Scan()
		cmd := scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Printf("Invalid input: %s", err)
		}
		if value, ok := listCommand[cmd]; ok {
			if err := value.callback(); err != nil {
				fmt.Println(err)
			}
			} else {
			fmt.Println("Unknow command")
		}	
	}
}

