package main

import (
	"fmt"
	"log"
	"task1/carddraw"
	"task1/cardgame"
)

func main() {

	deck := cardgame.NewDeck()
	fmt.Println(deck)

	deck.Shuffle()
	fmt.Println(deck)

	allCardsFromDeck, err := carddraw.DrawAllCards(&deck)

	if err != nil {
		fmt.Println("I quit due to error: ", err)
		log.Fatal(err)
	}

	fmt.Println(deck)
	fmt.Printf("All you take from deck is: %v", allCardsFromDeck)
}
