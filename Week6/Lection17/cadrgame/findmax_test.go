package cardgame

import "testing"

func TestFindMaxCard(t *testing.T) {
	deck := []card{{K, club}, {2, club}, {A, spade}, {A, diamond}, {10, heart}, {10, spade}}
	maxcard = card{A,spade}

	answer := FindMaxCard(deck)
	if maxcard != answer {
		t.Errorf("Expected %d, got %d", maxcard, answer)
	}
}
