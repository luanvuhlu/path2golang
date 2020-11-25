package main

import (
	"fmt"
	"os"
)

func main() {
	cards := newDeck()
	shuffleError := cards.shuffle(11)
	if shuffleError != nil {
		fmt.Println(shuffleError)
		os.Exit(-1)
	}
	// hand, remainingCards := deal(cards, 5)
	const fileName string = "myCards"
	writeError := cards.saveToFile(fileName)
	if writeError != nil {
		fmt.Println("Write file error.", writeError)
		os.Exit(-1)
	}
	cardsFromFile, readError := readFromFile(fileName)
	if readError != nil {
		fmt.Println("Read from file error.", readError)
		os.Exit(-1)
	}
	for _, card := range cardsFromFile {
		fmt.Println(card)
	}
}
