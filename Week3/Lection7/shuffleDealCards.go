package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"
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

type Card struct {
	value int
	suit  int
}

type Deck struct {
	cards []Card
}

func main() {

	deck := NewDeck()
	fmt.Println(deck)

	deck.Shuffle()
	fmt.Println(deck)

	c, er := deck.Deal()
	if er != nil {
		fmt.Println("Error in deling a card: ", er)
	} else {
		fmt.Printf("You take card: %s-%s\n", returnRealValue(c.value), returnRealSuit(c.suit))
	}

	c, er = deck.Deal()
	if er != nil {
		fmt.Println("Error in deling a card: ", er)
	} else {
		fmt.Printf("You take card: %s-%s\n", returnRealValue(c.value), returnRealSuit(c.suit))
	}

	c, er = deck.Deal()
	if er != nil {
		fmt.Println("Error in deling a card: ", er)
	} else {
		fmt.Printf("You take card: %s-%s\n", returnRealValue(c.value), returnRealSuit(c.suit))
	}

	c, er = deck.Deal()
	if er != nil {
		fmt.Println("Error in deling a card: ", er)
	} else {
		fmt.Printf("You take card: %s-%s\n", returnRealValue(c.value), returnRealSuit(c.suit))
	}
	fmt.Println(deck)
}

//Lection 7 Task1

func NewDeck() (deck Deck) {

	values := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, J, D, K, A}
	suits := []int{club, diamond, heart, spade}
	for val := 0; val < len(values); val++ {
		currentCard := Card{}
		for su := 0; su < len(suits); su++ {
			currentCard = Card{value: values[val], suit: suits[su]}
			deck.cards = append(deck.cards, currentCard)
		}
	}
	return deck
}

func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	shuffleNumbers := rand.Perm(len(d.cards)) //generate permutation in given interval
	shuffleDeck := Deck{}
	currentCard := Card{}

	for _, num := range shuffleNumbers {
		currentCard = d.cards[num]
		shuffleDeck.cards = append(shuffleDeck.cards, currentCard)
	}
	fmt.Println("You have just shuffled the deck!")
	//fmt.Println(shuffleDeck)
	*d = shuffleDeck
}

func (d *Deck) Deal() (*Card, error) {
	if len(d.cards) == 0 {
		return nil, errors.New("Ther are no more cards in the deck!")
	} else {
		firstCard := d.cards[0]
		d.cards = d.cards[1:]
		return &firstCard, nil
	}
}

//Lection 5 Task1

func compareTwoCards(card1 Card, card2 Card) int {
	if card1.value > card2.value || ((card1.value == card2.value) && (card1.suit > card2.suit)) {
		return 1
	} else if card1.value < card2.value || ((card1.value == card2.value) && (card1.suit < card2.suit)) {
		return -1
		// card1Val==Card2Val
	} else {
		return 0
	}
}

//Lection 5 task2

func findMaxCard(cards []Card) Card {
	// use compareCards here to find the maximum ...
	var maxCard Card = cards[0]

	for _, v := range cards {
		if compareTwoCards(v, maxCard) == 1 {
			maxCard = v
		}
	}
	return maxCard
}

func printTwoCards(c1v Card, c2v Card) {
	fmt.Printf("1 card--> %s %s\n", returnRealValue(c1v.value), returnRealSuit(c1v.suit))
	fmt.Printf("2 card--> %s %s\n", returnRealValue(c2v.value), returnRealSuit(c2v.suit))
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

func returnRealValue(cardVol int) string {
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

func validateCard(cardToValidate Card) bool {
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
