package main

import "strings"

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