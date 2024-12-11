package main

func main() {

	cards := newDeck()

	cards.saveToFile("my_cards")

	deck2 := newDeckFromFile("my_cards")
	deck2.shuffle()
	deck2.print()

}
