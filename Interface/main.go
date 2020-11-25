package main

import (
	"fmt"
)

type EnglishBot struct {
}

type SpanishBot struct {
}

func (EnglishBot) getGreeting() string {
	return "Hello"
}

func (SpanishBot) getGreeting() string {
	return "Halo"
}

func printGreeting(b Bot) {
	fmt.Println(b.getGreeting())
}

type Bot interface {
	getGreeting() string
}

func main() {
	englishBot := EnglishBot{}
	spanishBot := SpanishBot{}
	printGreeting(englishBot)
	printGreeting(spanishBot)
}
