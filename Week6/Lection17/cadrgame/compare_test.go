package cadrgame

import "testing"

func TestCompareTwoCards(t *testing.T) {
	smaller := card{value: 2, suit: club}
	bigger := card{value: 3, suit: club}
	answer:=CompareTwoCards(smaller,bigger)

	if answer!=-1{
		t.Errorf("Expected -1, got %d",answer)
	}

	answer=CompareTwoCards(bigger,smaller)
	if answer!=1{
		t.Errorf("Expected 1, got %d",answer)
	}

	answer=CompareTwoCards(bigger,bigger)
	if answer!=0{
		t.Errorf("Expected 1, got %d",answer)
	}
}
