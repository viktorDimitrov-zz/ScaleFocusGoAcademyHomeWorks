package main

import (
	"fmt"
	"main/hw/cardgame"
)

func main() {

	//compare two cards
	var card1 = cardgame.card{0, cardgame.club}
	var card2 = cardgame.card{int(cardgame.J), cardgame.spade}
	var result int
	if cardgame.validateCard(card1) && cardgame.validateCard(card2) {
		cardgame.printTwoCards(card1, card2)
		result = cardgame.CompareTwoCards(card1, card2)
		cardgame.printWinnerCard(result)
	}

	card1.value = cardgame.K
	card1.suit = cardgame.club
	card2.value = cardgame.K
	card2.suit = cardgame.club
	if cardgame.validateCard(card1) && cardgame.validateCard(card2) {
		cardgame.printTwoCards(card1, card2)
		result = cardgame.CompareTwoCards(card1, card2)
		cardgame.printWinnerCard(result)
	}

	//finding maximal card in slice of cards
	// the deck to find maximal card
	deck := []cardgame.card{{cardgame.K, cardgame.club}, {2, cardgame.club}, {cardgame.A, cardgame.spade}, {cardgame.A, cardgame.diamond}, {10, cardgame.heart}, {10, cardgame.spade}}

	fmt.Printf("%v---%T\n", deck, deck)
	max := cardgame.FindMaxCard(deck)
	fmt.Printf("Maximal card in deck is %v-%v", cardgame.returnRealVolume(max.value), cardgame.returnRealSuit(max.suit))
}
