package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "queen", "jack", "king", "ten":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	// Get the values of the player's cards and the dealer's card
	cardValue1 := ParseCard(card1)
	cardValue2 := ParseCard(card2)
	dealerCardValue := ParseCard(dealerCard)

	// Calculate the total value of the player's two cards
	total := cardValue1 + cardValue2

	switch {
	// Rule 1: If both cards are aces, split (P)
	case card1 == "ace" && card2 == "ace":
		return "P"
	// Rule 2: If it's a Blackjack (21) and the dealer does not show an ace, face card, or 10, win (W)
	case total == 21 && dealerCardValue < 10:
		return "W"
	// Rule 3: If the total is between 17 and 20, stand (S)
	case total >= 17 && total <= 20:
		return "S"
	// Rule 4: If the total is between 12 and 16, stand unless the dealer has 7 or higher, in which case hit
	case total >= 12 && total <= 16:
		if dealerCardValue >= 7 {
			return "H"
		}
		return "S"
	// Rule 5: If the total is 11 or lower, always hit (H)
	case total <= 11:
		return "H"
	default:
		return "S"
	}
}
