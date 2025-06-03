package main

import ( 
	"fmt"
	"strings"
	"bufio"
	"os"
	"net/http"
	"io"
	"encoding/json"
)

type cliCommand struct {
	name string
	description string
	callback func() error
}

type locationInfo struct {
	Name string
	Url string
}

type LocationArea struct {
	Count int
	Next string
	Previous string
	Results []locationInfo
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

func commandMap() error {
	resp, err := http.Get("https://pokeapi.co/api/v2/location-area?limit=20")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
//	fmt.Println(string(body))
	l := LocationArea{}
	err = json.Unmarshal(body, &l)
	if err != nil {
		return err
	}
	for _ , location := range l.Results {
		fmt.Println(location.Name)
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
		"map": {
			name: "map",
			description: "Desplays the next 20 locations",
			callback: commandMap,
		},
	}
}

func main(){
	fmt.Println("Welcome to the Pokedex!")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
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

