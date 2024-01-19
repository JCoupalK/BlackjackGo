package main

type gameState struct {
	PlayerHand      []Card
	DealerHand      []Card
	Deck            []Card
	GameOver        bool
	PlayerWins      bool
	Tie             bool
	DealerRevealing bool
}

func (gs *gameState) DealInitialCards() {
	for i := 0; i < 2; i++ {
		card, deck := dealCard(gs.Deck)
		gs.Deck = deck
		gs.PlayerHand = append(gs.PlayerHand, card)

		card, deck = dealCard(gs.Deck)
		gs.Deck = deck
		gs.DealerHand = append(gs.DealerHand, card)
	}
}

func (gs *gameState) PlayerHit() {
	// Check if the current score is already 21, prevent hitting
	if gs.CalculateScore(gs.PlayerHand) == 21 {
		return
	}

	card, deck := dealCard(gs.Deck)
	gs.Deck = deck
	gs.PlayerHand = append(gs.PlayerHand, card)

	if gs.CalculateScore(gs.PlayerHand) > 21 {
		gs.GameOver = true
		gs.PlayerWins = false
	}
}

func (gs *gameState) DealerPlay() {
	for gs.CalculateScore(gs.DealerHand) < 17 {
		card, deck := dealCard(gs.Deck)
		gs.Deck = deck
		gs.DealerHand = append(gs.DealerHand, card)
	}
	gs.DealerRevealing = true
}

func (gs *gameState) CalculateScore(hand []Card) int {
	score := 0
	aces := 0
	for _, card := range hand {
		switch {
		case card.Rank == "A":
			aces++
			score += 11
		case card.Rank == "K", card.Rank == "Q", card.Rank == "J":
			score += 10
		case card.Rank == "10":
			score += 10
		default:
			score += int(card.Rank[0] - '0') // Safe for 2-9
		}
	}

	// Adjust for Aces if score exceeds 21
	for score > 21 && aces > 0 {
		score -= 10
		aces--
	}
	return score
}

func (gs *gameState) DetermineWinner() {
	playerScore := gs.CalculateScore(gs.PlayerHand)
	dealerScore := gs.CalculateScore(gs.DealerHand)

	if playerScore > 21 {
		gs.GameOver = true
		gs.PlayerWins = false
	} else if dealerScore > 21 || playerScore > dealerScore {
		gs.GameOver = true
		gs.PlayerWins = true
	} else if dealerScore > playerScore {
		gs.GameOver = true
		gs.PlayerWins = false
	} else {
		gs.GameOver = true
		gs.Tie = true
	}
}
