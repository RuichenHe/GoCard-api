package deck

type Suit uint8

const (
	Spade Suit = iota //Go's iota identifier is used in const declarations to simplify definitions of incrementing numbers
	Diamond
	Club
	Heart
	Joker
)

type Rank uint8

const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

type Card struct {
	Suit
	Rank
}
