package main

import (
    "os"
    "fmt"
    "bufio"
)

func StartRepl() {

    cfg := &config{}

    commands["exit"] = cliCommand{
        name:        "exit",
        description: "Exit the Pokedex",
        callback:    commandExit,
    }
    
    commands["help"] = cliCommand{
        name:        "help",
        description: "Displays a help message",
        callback:    commandHelp,
    }

    commands["map"] = cliCommand{
        name:        "map",
        description: "Displays the names of 20 location areas, or the next page of 20, in the Pokemon world.",
        callback:    commandMap,
    }

    commands["mapb"] = cliCommand{
        name:        "mapb",
        description: "Displays the previous 20 location areas.",
        callback:    commandMapb,
    }
    

    scanner := bufio.NewScanner(os.Stdin)

    for {
        // Use fmt.Print to print the prompt Pokedex > without a newline character
        fmt.Print("Pokedex > ")
        // Use the scanner's .Scan and .Text methods to get the user's input as a string
        scanner.Scan()
        // Clean the user's input string
        text := CleanInput(scanner.Text())
        if len(text) == 0 {
            continue
        }
        cmdName := text[0]
        args := text[1:]
        cmd, ok := commands[cmdName]
        if !ok {
            fmt.Println("Not a recognized command")
            continue
        }
        if err := cmd.callback(cfg, args); err != nil {
            fmt.Println(err)
        }
    }
}
