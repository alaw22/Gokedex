package main

import 
(
	"fmt"
	"strings"
	"bufio"
	"os"
)

func cleanInput(text string) []string {
	clean_input := []string{}
	split := strings.Split(strings.ToLower(text)," ")

	for _, word := range split {
		if word != ""{
			clean_input = append(clean_input,word)
		}
	}
	return clean_input
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var input string
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input = scanner.Text()
		clean_input := cleanInput(input)
		if len(clean_input) == 0{
			continue
		}
		fmt.Printf("Your command was: %v\n",clean_input[0])
	}
}
