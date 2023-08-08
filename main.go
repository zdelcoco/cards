package main

func main() {
	// initialize a new deck
	cardDeck := newDeck()

	// deal 5 cards to 3 people
	playerDecks := cardDeck.deal(3, 5)

	// output player hands
	playerDecks.printDecks()

	// output remaining cards in deck
	cardDeck.print()

}
