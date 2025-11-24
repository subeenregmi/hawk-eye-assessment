package deck

import (
	"math/rand"

	"github.com/subeenregmi/hawk-eye-assessment/internal/card"
	"github.com/subeenregmi/hawk-eye-assessment/internal/rank"
	"github.com/subeenregmi/hawk-eye-assessment/internal/suit"
)

func NewDeck() deck {
	cards := make([]card.Card, 0, 52)

	for _, suit := range suit.Suits {
		for _, rank := range rank.Ranks {
			cards = append(cards, card.Card{
				Rank: rank,
				Suit: suit,
			})
		}
	}

	return deck{
		Cards: cards,
	}
}

func (d *deck) Shuffle() {
	for i := range d.Cards {
		j := rand.Intn(len(d.Cards))

		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}
}

func (d *deck) DrawRandomCard() card.Card {
	i := rand.Intn(len(d.Cards))
	card := d.Cards[i]

	d.Cards = append(d.Cards[:i], d.Cards[i+1:]...)

	return card
}

func (d *deck) AddJoker() {
	d.Cards = append(d.Cards, card.Card{Rank: rank.Joker})
}
