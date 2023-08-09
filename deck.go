package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// Create a new type of 'deckSlice'
// which is a slice of decks
type deckSlice []deck

func removeIndex(d deck, index int) deck {
	return append(d[:index], d[index+1:]...)
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) printDeck() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (ds deckSlice) printDecks() {
	for i, player := range ds {
		fmt.Println("Player ", i+1)
		player.printDeck()
		fmt.Println("-----------")
	}
}

func (d *deck) dealCard() string {
	returnCard := (*d)[0]
	*d = removeIndex(*d, 0)
	return returnCard
}

func (d deck) shuffleDeck() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func (d *deck) deal(numDecks int, numCardsToDealEachDeck int) deckSlice {
	dealIteration := 0
	var returnDecks []deck = make([]deck, numDecks)

	for dealIteration < numCardsToDealEachDeck {
		for i := 0; i < numDecks; i++ {
			returnDecks[i] = append(returnDecks[i], (*d).dealCard())
		}
		dealIteration++
	}

	return returnDecks

}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0777)
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}
