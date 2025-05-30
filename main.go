package main

import ( 
	//"fmt"
	"strings"
)

func cleanInput(text string) []string { 
	lowerText := strings.ToLower(text)
	words := strings.Fields(lowerText)
	//fmt.Println(words)
	return words
}


func main(){
	text1 := "Hello, World!"
	cleanInput(text1)
}

