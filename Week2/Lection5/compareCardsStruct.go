package main

import (
	"fmt"
	"strconv"
)

const (
	club = iota
	diamond
	heart
	spade
)

const (
	J = iota + 11
	D
	K
	A
)

type card struct {
	value int
	suit  int
}

func main() {
	//compare two cards
	var card1 = card{0, club}
	var card2 = card{int(J), spade}
	var result int
	if validateCard(card1) && validateCard(card2) {
		printTwoCards(card1, card2)
		result = compareTwoCards(card1, card2)
		printWinnerCard(result)
	}

	card1 = card{int(A), club}
	card2 = card{int(J), spade}
	if validateCard(card1) && validateCard(card2) {
		printTwoCards(card1, card2)
		result = compareTwoCards(card1, card2)
		printWinnerCard(result)
	}

	card1 = card{2, club}
	card2 = card{2, spade}
	if validateCard(card1) && validateCard(card2) {
		printTwoCards(card1, card2)
		result = compareTwoCards(card1, card2)
		printWinnerCard(result)
	}

	card1.value = K
	card1.suit = club
	card2.value = K
	card2.suit = club
	if validateCard(card1) && validateCard(card2) {
		printTwoCards(card1, card2)
		result = compareTwoCards(card1, card2)
		printWinnerCard(result)
	}

	//finding maximal card in slice of cards
	// the deck to find maximal card
	deck := []card{{K, club}, {2, club}, {A, spade}, {A, diamond}, {10, heart}, {10, spade}}

	fmt.Printf("%v---%T\n", deck, deck)
	max := findMaxCard(deck)
	fmt.Printf("Maximal card in deck is %v-%v", returnRealVolume(max.value), returnRealSuit(max.suit))
}

func compareTwoCards(card1 card, card2 card) int {
	if card1.value > card2.value || ((card1.value == card2.value) && (card1.suit > card2.suit)) {
		return 1
	} else if card1.value < card2.value || ((card1.value == card2.value) && (card1.suit < card2.suit)) {
		return -1
		// card1Val==Card2Val
	} else {
		return 0
	}
}

func findMaxCard(cards []card) card {
	// use compareCards here to find the maximum ...
	var maxCard card = cards[0]

	for _, v := range cards {
		if compareTwoCards(v, maxCard) == 1 {
			maxCard = v
		}
	}
	return maxCard
}

func printTwoCards(c1v card, c2v card) {
	fmt.Printf("1 card--> %s %s\n", returnRealVolume(c1v.value), returnRealSuit(c1v.suit))
	fmt.Printf("2 card--> %s %s\n", returnRealVolume(c2v.value), returnRealSuit(c2v.suit))
}

func returnRealSuit(cardInt int) string {
	var resultSuit string
	switch cardInt {
	case 0:
		resultSuit = "club"
	case 1:
		resultSuit = "diamond"
	case 2:
		resultSuit = "heart"
	case 3:
		resultSuit = "spade"
	default:
		resultSuit = "Mistake suit"
	}
	return resultSuit
}

func returnRealVolume(cardVol int) string {
	var realVol string
	switch cardVol {

	case 2, 3, 4, 5, 6, 7, 8, 9, 10:
		realVol = strconv.Itoa(cardVol)
	case 11:
		realVol = "J"
	case 12:
		realVol = "Q"
	case 13:
		realVol = "K"
	case 14:
		realVol = "A"

	}
	return realVol
}

func printWinnerCard(res int) {
	switch res {
	case 1:
		fmt.Println("First card")
	case 0:
		fmt.Println("Equual")
	case -1:
		fmt.Println("Second card")
	default:
		fmt.Println("Mistake!")
	}
}

func validateCard(cardToValidate card) bool {
	var result bool = true
	if !(cardToValidate.value >= 2 && cardToValidate.value <= 14) {
		fmt.Printf("Value of card is %d shoud be in range 2-14!\n", cardToValidate.value)
		result = false
	}
	if !(cardToValidate.suit >= 0 && cardToValidate.suit <= 3) {
		fmt.Printf("Not valid suit - %d!\n", cardToValidate.suit)
		result = false
	}
	return result
}
