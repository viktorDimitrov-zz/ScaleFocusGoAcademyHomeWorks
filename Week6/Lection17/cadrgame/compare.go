package cadrgame

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

func CompareTwoCards(card1 card, card2 card) int {
	if card1.value > card2.value || ((card1.value == card2.value) && (card1.suit > card2.suit)) {
		return 1
	} else if card1.value < card2.value || ((card1.value == card2.value) && (card1.suit < card2.suit)) {
		return -1
		// card1Val==Card2Val
	} else {
		return 0
	}
}






