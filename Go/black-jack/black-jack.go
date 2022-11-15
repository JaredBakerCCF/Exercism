package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    switch {
    case card == "ace":
    		return 11
    case card == "king":
    		return 10
    case card == "queen":
    		return 10
    case card == "jack":
    		return 10
    case card == "ten":
    		return 10
    case card == "nine":
    		return 9
    case card == "eight":
    		return 8
    case card == "seven":
    		return 7
    case card == "six":
    		return 6
    case card == "five":
    		return 5
    case card == "four":
    		return 4
    case card == "three":
    		return 3
    case card == "two":
    		return 2 
    case card == "joker":
    		return 0
    }
	panic("Please implement the ParseCard function")
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    Card1 := ParseCard(card1)
    Card2 := ParseCard(card2)
    DealerCard := ParseCard(dealerCard)
    var Sum int = Card1 + Card2
    switch {
        case Sum == 21 && DealerCard != 10 && DealerCard != 11: //Black Jack with non figure card
    		return "W"
        case Sum == 21 && DealerCard == 11 : //Black Jack with ten figure dealer
    		return "S"
        case Sum == 21 && DealerCard == 10 : //Black Jack with ten figure dealer
    		return "S"
        case Card1 == 11 && Card2 == 11 : //Pocket Aces
    		return "P"
        case Sum >= 17 && Sum <= 20 : //Greater than 17 but less than 20
    		return "S"
        case Sum >= 12 && Sum <= 16 && DealerCard < 7 : //Greater than 12 but less than 16 and dealer has less than 7
    		return "S"
        case Sum >= 12 && Sum <= 16 && DealerCard >= 7 : //Greater than 12 but less than 16 and dealer has greater than 7
    		return "H"
        case Sum <= 11 :
    		return "H"
    }
	panic("Please implement the FirstTurn function")
}
