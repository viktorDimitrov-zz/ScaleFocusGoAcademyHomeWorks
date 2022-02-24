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

func main() {

	var c1Value, c1Suit int = 3, club
	var c2Value, c2Suit int = J, spade

	var result int = CompareCardsI(c1Value, c1Suit, c2Value, c2Suit)
	PrintWinner(result)

	c1Value, c1Suit = A, club
	c2Value, c2Suit = J, spade
	result = CompareCardsI(c1Value, c1Suit, c2Value, c2Suit)
	PrintWinner(result)

	c1Value, c1Suit = 2, club
	c2Value, c2Suit = 2, spade
	result = CompareCardsI(c1Value, c1Suit, c2Value, c2Suit)
	PrintWinner(result)

	c1Value, c1Suit = K, club
	c2Value, c2Suit = K, club
	result = CompareCardsI(c1Value, c1Suit, c2Value, c2Suit)
	PrintWinner(result)

}

func CompareCardsI(cardOneVal int, cardOneSuit int, cardTwoVal int, cardTwoSuit int) int {
	PrintTwoCards(cardOneVal, cardOneSuit, cardTwoVal, cardTwoSuit)

	if cardOneVal > cardTwoVal {
		return 1
	} else if cardOneVal < cardTwoVal {
		return -1
		// card1Val==Card2Val
	} else {
		if cardOneSuit > cardTwoSuit {
			return 1
		} else if cardOneSuit < cardTwoSuit {
			return -1
		} else {
			return 0
		}
	}

}
func PrintTwoCards(c1v int, c1s int, c2v int, c2s int) {

	fmt.Printf("1 card--> %s %s\n", ReturnRealVolume(c1v), ReturnRealSuit(c1s))
	fmt.Printf("2 card--> %s %s\n", ReturnRealVolume(c2v), ReturnRealSuit(c2s))
}

func ReturnRealSuit(cardInt int) string {
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

func ReturnRealVolume(cardVol int) string {
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

func PrintWinner(res int) {
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

func validateCards(value int, suit int) bool {
	var result bool = true
	if !(value >= 2 && value <= 14) {
		fmt.Printf("Value of card is %d shoud be in range 2-14!", value)
		result = false
	}
	if !(suit >= 0 && suit <= 3) {
		fmt.Println("Not valid suit - %d!", suit)
		result = false
	}
	return result
}
