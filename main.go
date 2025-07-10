package main

import 
(
    "fmt"
    "bufio"
    "os"
)

const GokedexURL = ""

func main() {

    scanner := bufio.NewScanner(os.Stdin)
    var input string
    /* 
        This is a REPL (Read-Evaluate-Print-Loop) much like
        any shell. It reads the users input evaluates it by 
        doing some functionality and prints the result. 
    */
    for {
        // Prompt
        fmt.Print("Pokedex > ")

        // Get input
        scanner.Scan()
        input = scanner.Text()

        // Clean input
        clean_input := cleanInput(input)
        if len(clean_input) == 0{
            continue
        }

        commandName := clean_input[0]
        command, ok := commandRegistry[commandName]
        if !ok {
            fmt.Println("Unknown command")
            continue
        }

        err := command.callback()
        if err != nil{
            fmt.Println(err)
        }

    }
}
