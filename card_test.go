package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	//Output:
	//Ace of Hearts
}

func TestNew(t *testing.T) {
	cards := New()
	if len(cards) != 4*13 {
		t.Error("Wrong number of cards!")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	exp := Card{Rank: Ace, Suit: Spade}
	if cards[0] != exp {
		t.Error("Expect Ace of Spades to be the first one!recieved:", cards[0])
	}
}

func TestSort(t *testing.T) {
	cards := New(Sort(Less))
	exp := Card{Rank: Ace, Suit: Spade}
	if cards[0] != exp {
		t.Error("Expect Ace of Spades to be the first one!recieved:", cards[0])
	}
}
