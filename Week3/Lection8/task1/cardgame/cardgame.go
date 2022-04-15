package cardgame

import (
	"fmt"
	"math/rand"
	"time"
)

type Card struct {
	value int
	suit  int
}

type Deck struct {
	cards []Card
}

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

func (d *Deck) Deal() *Card {
	if len(d.cards) == 0 {
		return nil
	}
	firstCard := d.cards[0]
	d.cards = d.cards[1:]
	
	return &firstCard
}
