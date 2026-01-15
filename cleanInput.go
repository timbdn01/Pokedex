package main

import "strings"

func cleanInput(text string) []string {
	cleaned := []string{}
	// Implementation of cleaning the input string
	words := strings.Fields(text)
	for _, word := range words {
		word = strings.TrimSpace(word)
		if word == "" {
			continue
		}
		cleaned = append(cleaned, strings.ToLower(word))
	}
	return cleaned
}