package main

func main() {
	cards := newDeck()
	// hand, remainingCards := cards.deal(5)
	// cards.shuffle()
	cards.saveToFile("my-cards")

	cards = loadFromFile("my-cards")
	cards.print()
}
