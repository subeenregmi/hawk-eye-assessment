package card

import (
	"github.com/subeenregmi/hawk-eye-assessment/internal/rank"
	"github.com/subeenregmi/hawk-eye-assessment/internal/suit"
)

type Card struct {
	Rank rank.Rank
	Suit suit.Suit
}
