package main

import (
	"fmt"
	"strings"

	"github.com/muesli/termenv"
)

// termenv for color styling
var p = termenv.ColorProfile()

// renderCardSymbol returns a string symbol representation of a card with color
func renderCardSymbol(card Card) string {
	var suitSymbol string
	var suitColor termenv.Color

	switch card.Suit {
	case "Hearts", "Diamonds":
		suitSymbol = "♥"               // Use "♦" for Diamonds
		suitColor = p.Color("#FF0000") // Red color
	case "Clubs", "Spades":
		suitSymbol = "♣"               // Use "♠" for Spades
		suitColor = p.Color("#FFFFFF") // Black color
	}

	coloredSymbol := termenv.String(suitSymbol).Foreground(suitColor).String()
	return fmt.Sprintf("%s%s", card.Rank, coloredSymbol)
}

// renderHandSymbol returns a string representation of a hand of cards using symbols
func renderHandSymbol(hand []Card) string {
	var cards []string
	for _, card := range hand {
		cards = append(cards, renderCardSymbol(card))
	}
	return strings.Join(cards, " ")
}

// renderGameState returns a string representation of the current game state with a table-like layout
func renderGameState(gs gameState, width int) string {
	var sb strings.Builder

	sb.WriteString("\n🃏 Blackjack Table 🃏\n\n")
	sb.WriteString("Dealer's Hand:\n")
	if gs.GameOver {
		sb.WriteString("  " + renderHandSymbol(gs.DealerHand) + "\n")
	} else {
		sb.WriteString("  [🂠] " + renderCardSymbol(gs.DealerHand[1]) + "\n") // Hidden card
	}

	sb.WriteString("\nYour Hand:\n")
	sb.WriteString("  " + renderHandSymbol(gs.PlayerHand) + "\n")

	playerScore, isSoftPlayer := gs.CalculateScore(gs.PlayerHand)
	scoreText := fmt.Sprintf("%d", playerScore)
	if isSoftPlayer {
		scoreText = fmt.Sprintf("Soft %s", scoreText)
	}
	sb.WriteString(fmt.Sprintf("\nYour Score: %s\n", scoreText))

	if gs.GameOver {
		dealerScore, isSoftDealer := gs.CalculateScore(gs.DealerHand)
		dealerScoreText := fmt.Sprintf("%d", dealerScore)
		if isSoftDealer {
			dealerScoreText = fmt.Sprintf("Soft %s", dealerScoreText)
		}
		sb.WriteString(fmt.Sprintf("Dealer's Score: %s\n", dealerScoreText))
		sb.WriteString("\n" + gameOutcomeMessage(gs) + "\n")
	} else if playerScore == 21 {
		sb.WriteString("\n21! Press 'S'.\n")
	} else {
		sb.WriteString("\nPress 'H' to hit or 'S' to stand.\n")
	}

	return sb.String()
}

// gameOutcomeMessage returns a string message about the game outcome
func gameOutcomeMessage(gs gameState) string {
	var sb strings.Builder

	if gs.Tie {
		sb.WriteString("🤝 It's a tie!\n\nPress 'R' to restart.")
	} else if gs.PlayerWins {
		sb.WriteString("🎉 You win!\n\nPress 'R' to restart.")
	} else {
		sb.WriteString("💔 Dealer wins!\n\nPress 'R' to restart.")
	}

	return sb.String()
}
