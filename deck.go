package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type of 'deck' which is a slice of strings

type deck []string

func newDeck() deck {
	// create and return a list of playing cards
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// create a new function that belongs to the deck type:
// when called, print each card in a deck.
// We do this via a Go reciever function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)

}

func (d deck) shuffle() {
	// create random index number to swap, using current time as int64 to provide random seed
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// now use r (of Type rand, with custom seed, to randomize deck)
	for i := range d {
		newPos := r.Intn(len(d) - 1)

		// now use this number to swap current pos of slice with the newPos
		d[i], d[newPos] = d[newPos], d[i]
	}
}
