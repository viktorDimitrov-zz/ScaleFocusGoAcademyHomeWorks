package cadrgame

import (
	"fmt"
	"strconv"
)

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
