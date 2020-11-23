package main

import "fmt"

type deck []string

func (d deck) print() {
	for i := 0; i < len(d); i++ {
		fmt.Println(d[i])
	}
}

func (d deck) shuffle() {

}

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamnds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}
	return cards
}
