package main

import ( 
	"fmt"
	"strings"
	"bufio"
	"os"
	"net/http"
	"io"
	"encoding/json"
	"github.com/rafael0rueda/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	name string
	description string
	callback func(areaPointer *LocationArea) error
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

func commandExit(currentArea *LocationArea) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(currentArea *LocationArea) error {
	for _, value := range listCommand {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}
	return nil
}

func commandMap(currentArea *LocationArea) error {
	var url string
	if currentArea.Next == "" {
		url = "https://pokeapi.co/api/v2/location-area?limit=20" 
	} else {
		url = currentArea.Next
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, currentArea)
	if err != nil {
		return err
	}
	for _ , location := range currentArea.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapBack(currentArea *LocationArea) error {
	if currentArea.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	
	resp, err := http.Get(currentArea.Previous)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, currentArea)
	if err != nil {
		return err
	}
	for _ , location := range currentArea.Results {
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
			description: "Displays the next 20 locations",
			callback: commandMap,
		},
		"mapb": {
			name: "map back",
			description: "Displays the previous 20 locations",
			callback: commandMapBack,
		},
	}
}

func main(){
	l := LocationArea{}
	var locationPointer *LocationArea = &l
	fmt.Println("Welcome to the Pokedex!")
	pokecache.ShowTime()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		cmd := scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Printf("Invalid input: %s", err)
		}
		if value, ok := listCommand[cmd]; ok {
			if err := value.callback(locationPointer); err != nil {
				fmt.Println(err)
			}
			} else {
			fmt.Println("Unknow command")
		}	
	}
}

