package main

import (
	"fmt"
	"task1/carddraw"
	"task1/cardgame"
)

func main() {

	deck := cardgame.NewDeck()
	fmt.Println(deck)

	deck.Shuffle()
	fmt.Println(deck)

	allCardsFromDeck:=carddraw.DrawAllCards(&deck)
	fmt.Println(deck)

	fmt.Printf("All you take from deck is: %v",allCardsFromDeck)
}
