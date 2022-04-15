package cadrgame

func FindMaxCard(cards []card) card {
	// use compareCards here to find the maximum ...
	var maxCard card = cards[0]

	for _, v := range cards {
		if CompareTwoCards(v, maxCard) == 1 {
			maxCard = v
		}
	}
	return maxCard
}