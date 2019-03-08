package main

import "fmt"

func main() {
	cards := newDeck()
	cards.saveToFile("foo.out")

	//cards := newDeckFromFile("foo.out")
	//cards.print()

	cards.shuffle()

	fmt.Println("\n New deck!")
	cards.print()
}
