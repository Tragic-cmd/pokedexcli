package main

import "strings"

func CleanInput(text string) []string {
	lowered := strings.ToLower(text)
	words_list := strings.Fields(lowered)
    return words_list
}

