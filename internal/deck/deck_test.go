package deck

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/subeenregmi/hawk-eye-assessment/internal/card"
)

func TestNewDeck(t *testing.T) {
	d := NewDeck()

	assert.Len(t, d.Cards, 52)
}

func TestShuffle(t *testing.T) {
	d := NewDeck()

	d.Shuffle()

	// This test can fail, but the odds are EXTREMELY low 1/(52!)
	assert.NotEqual(t, NewDeck(), d)
}

func TestDrawRandomCard(t *testing.T) {
	d := NewDeck()

	card := d.DrawRandomCard()

	assert.Len(t, d.Cards, 51)
	assert.NotContains(t, d.Cards, card)
}

func TestDrawingAll(t *testing.T) {
	d := NewDeck()

	var cards []card.Card

	for range 52 {
		cards = append(cards, d.DrawRandomCard())
	}

	assert.ElementsMatch(t, cards, NewDeck().Cards)
}
