package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	suits := []string{"Clubs", "Diamonds", "Hearts", "Spades"}
	values := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPos := r.Intn(len(d) - 1)

		d[i], d[newPos] = d[newPos], d[i]
	}
}

func deal(d deck, handSize int) (deck, deck) {
	if handSize > len(d) {
		return d, nil
	}
	return d[:handSize], d[handSize:]
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bd, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	return strings.Split(string(bd), ",")
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
