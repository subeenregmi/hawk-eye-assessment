package game

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/subeenregmi/hawk-eye-assessment/internal/deck"
	"github.com/subeenregmi/hawk-eye-assessment/internal/rank"
)

func (g *Game) Start() error {
	d := deck.NewDeck()

	// add jokers before starting
	for i := 0; i < int(g.Jokers); i++ {
		d.AddJoker()
	}

	d.Shuffle()

	displayRules()

	currentCard := d.DrawRandomCard()

	// ensure we dont start on the joker card
	for currentCard.Rank == rank.Joker {
		currentCard = d.DrawRandomCard()
	}

	fmt.Printf("The starting card is: %v\n", currentCard)

	for i := 0; i < int(g.WinThreshold); i++ {
		userPrediction, err := getUserPrediction()
		if err != nil {
			return err
		}

		nextCard := d.DrawRandomCard()

		if nextCard.Rank == rank.Joker {
			fmt.Printf("JOKER %v! You have been forced to resign\n", nextCard)
			return nil
		}

		if nextCard.Equal(currentCard) {
			fmt.Printf("LUCKY! The card was the same: %v\n", nextCard)
			continue
		}

		if nextCard.GreaterThan(currentCard) != userPrediction {
			fmt.Printf("FAIL! The card was: %v\n", nextCard)
			return nil
		}

		fmt.Printf("CORRECT! The next card is: %v\n", nextCard)
		currentCard = nextCard
	}

	fmt.Println("CONGRATULATIONS! You have successfully completed the game!")

	return nil
}

func displayRules() {
	fmt.Printf(`Welcome to Higher ⬆️/ Lower ⬇️! 

The game is very simple, guess if the next card will be higher or lower than the previous.

On each turn enter a 'H' or 'L' to indicate what you think! Good luck :)

`)
}

func getUserPrediction() (bool, error) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			return false, err
		}

		// get rid of the new line character at the end
		input = strings.TrimSpace(input)

		if input != "H" && input != "L" {
			fmt.Println("> input needs to be either H (higher) or L (lower)")
			continue
		}

		return input == "H", nil
	}
}
