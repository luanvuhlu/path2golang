package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type deck []string

func (d deck) print() {
	for i := 0; i < len(d); i++ {
		fmt.Println(d[i])
	}
}

func (d deck) shuffle(times int) error {
	if times > 10 {
		return fmt.Errorf("You shuffle too many times")
	}
	if times < 1 {
		return nil
	}
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
	return d.shuffle(times - 1)
}

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamnds", "Clubs"}
	cardValues := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func readFromFile(filename string) ([]string, error) {
	bs, error := ioutil.ReadFile(filename)
	return strings.Split(string(bs), ","), error
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
