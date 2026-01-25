package main

import (
	"os"
	"fmt"
	"net/http"
	"io"
    "encoding/json"
)

var commands = map[string]cliCommand{}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, args []string) error
}

type config struct {
	next        string
	previous 	string
}

type locationAreaResponse struct {
    Results  []struct {
        Name string `json:"name"`
        URL  string `json:"url"`
    } `json:"results"`
    Next     *string `json:"next"`
    Previous *string `json:"previous"`
}


func commandExit(cfg *config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, args []string) error {

    fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
		
	// print help first, then exit
	if cmd, ok := commands["help"]; ok {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	if cmd, ok := commands["exit"]; ok {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

    return nil
}

func commandMap(cfg *config, args []string) error {
	var url string
	if cfg.next == "" {
        url = "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
    } else {
        url = cfg.next
    }

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("a network error occurred")
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("status code != 200")
		return err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var resp locationAreaResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return err
	}

	for _, r := range resp.Results {
		fmt.Println(r.Name)
	}

	if resp.Next != nil {
    cfg.next = *resp.Next
	} else {
		cfg.next = ""
	}

	if resp.Previous != nil {
		cfg.previous = *resp.Previous
	} else {
		cfg.previous = ""
	}

    return nil
}

func commandMapb(cfg *config, args []string) error {
	var url string
	if cfg.previous == "" {
        fmt.Println("You're on the first page")
		return nil
    } else {
        url = cfg.previous
    }
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("a network error occurred")
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("status code != 200")
		return err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var resp locationAreaResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return err
	}

	for _, r := range resp.Results {
		fmt.Println(r.Name)
	}

	if resp.Next != nil {
    cfg.next = *resp.Next
	} else {
		cfg.next = ""
	}

	if resp.Previous != nil {
		cfg.previous = *resp.Previous
	} else {
		cfg.previous = ""
	}

    return nil
}