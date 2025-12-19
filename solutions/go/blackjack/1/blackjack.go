package blackjack

import "slices"

var cardNames = []string{"ace", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "jack", "queen", "king"}

var cardsToNumber = map[string]int{"ace":11, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven":7, "eight": 8, "nine":9, "ten":10, "jack":10, "queen":10, "king":10}

const (
    optionStand = "S"
    optionHit = "H"
    optionSplit = "P"
    optionWin = "W"
    blackJack = 21
)

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    switch {
    case slices.Contains(cardNames, card):
        return cardsToNumber[card]
    default:
        return 0
    }
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	cardNumber1 := ParseCard(card1)
    cardNumber2 := ParseCard(card2)
    sum := cardNumber1 + cardNumber2
    dealerCardNumber := ParseCard(dealerCard)

	switch {
    case cardNumber1 == cardsToNumber["ace"] && cardNumber2 == cardsToNumber["ace"]:
        return optionSplit
    case sum == blackJack:
        dealerWins := []string{"jack", "queen", "king", "ace", "ten"}
        if !slices.Contains(dealerWins, dealerCard) {
            return optionWin
        } else {
            return optionStand
        }
    case sum >= 17 && sum <= 20:
        return optionStand
    case sum >= 12 && sum <= 16:
        if dealerCardNumber >= 7 {
            return optionHit
        } else {
            return optionStand
        }
    case sum <= 11:
        return optionHit
    default:
        return "unknown"
    }
}
