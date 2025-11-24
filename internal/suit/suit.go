package suit

type Suit string

const (
	Spades   Suit = "Spades"
	Hearts   Suit = "Hearts"
	Diamonds Suit = "Diamonds"
	Clubs    Suit = "Clubs"
)

var Suits = [4]Suit{Spades, Hearts, Diamonds, Clubs}
