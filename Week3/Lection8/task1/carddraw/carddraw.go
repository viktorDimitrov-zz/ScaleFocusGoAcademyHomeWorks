package carddraw

import (
	"fmt"
	"task1/cardgame"
)

type dealer interface {
	Deal() *cardgame.Card
}

func DrawAllCards(deal dealer) []cardgame.Card {
	// call the dealer's Draw() method, until you reach a nil Card
	dealedCards := []cardgame.Card{}
	for {
		if c:=deal.Deal(); c!= nil {
			fmt.Printf("You deal %d\n",c)
			dealedCards = append(dealedCards, *c)
		} else {
			return dealedCards
		}
	}
}
