package card

import (
	"strconv"

	"github.com/subeenregmi/hawk-eye-assessment/internal/rank"
	"github.com/subeenregmi/hawk-eye-assessment/internal/suit"
)

func (c Card) String() string {
	var symbol string

	switch c.Rank {
	case rank.Ace:
		symbol = "A"
	case rank.Jack:
		symbol = "J"
	case rank.Queen:
		symbol = "Q"
	case rank.King:
		symbol = "K"
	default:
		symbol = strconv.Itoa(int(c.Rank))
	}

	symbol += " "

	switch c.Suit {
	case suit.Clubs:
		symbol += "♣"
	case suit.Diamonds:
		symbol += "♦"
	case suit.Hearts:
		symbol += "♥"
	case suit.Spades:
		symbol += "♠"
	}

	return symbol
}

func (c *Card) GreaterThan(b Card) bool { return c.Rank > b.Rank }

func (c *Card) LessThan(b Card) bool { return c.Rank < b.Rank }

func (c *Card) Equal(b Card) bool { return c.Rank == b.Rank }

func (c *Card) GreaterThanEqual(b Card) bool { return c.Rank >= b.Rank }

func (c *Card) LessThanEqual(b Card) bool { return c.Rank <= b.Rank }
