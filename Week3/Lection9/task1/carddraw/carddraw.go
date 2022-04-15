package carddraw

import (
	"fmt"
	"task1/cardgame"
)

type dealer interface {
	Draw() (*cardgame.Card, error)
	Done() bool
}

func DrawAllCards(deal dealer) ([]cardgame.Card, error) {
	// call the dealer's Draw() method, until you reach a nil Card
	dealedCards := []cardgame.Card{}
	for {
		c, er := deal.Draw()
		if er == cardgame.ErrNoCardsInDeck && !deal.Done() {
			if deal.Done() {
				return dealedCards, nil
			} else {
				return nil, fmt.Errorf("erro I cannot draw a card:%w ", er)
			}
		} else {
			dealedCards = append(dealedCards, *c)
		}
		fmt.Println(c)
	}
}
