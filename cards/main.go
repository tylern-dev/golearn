package main

func main() {
	// cards := newDeck()

	// cards.saveToFile("my_cards")

	// cardBytes, _ := newDeckFromFile("my_cards")

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
