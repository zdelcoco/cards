package main

func main() {
	// initialize a new deck
	cardDeck := newDeck()

	cardDeck.shuffleDeck()

	// deal 5 cards to 3 people
	playerDecks := cardDeck.deal(3, 5)

	// output player hands
	playerDecks.printDecks()

	// output remaining cards in deck
	cardDeck.printDeck()

	// fmt.Println(cardDeck.toString())

	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))

	// errMessage := cardDeck.saveToFile("card_deck")
	// if errMessage != nil {
	// 	fmt.Println(errMessage)
	// }

	// cardDeck := newDeckFromFile("card_deck")
	// cardDeck.print()

}
