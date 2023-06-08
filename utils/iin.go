package utils

import (
	"strconv"
)

var cardRanges = map[string][][]int{
	"American Express": {{34}, {37}},
	"Diners Club":      {{36}, {300, 305}},
	"Discover":         {{6011}, {644, 649}, {65}, {622126, 622925}},
	"JCB":              {{3528, 3589}},
	"Maestro":          {{50}, {56, 69}},
	"Master":           {{2221, 2720}, {51, 55}},
	// "UnionPay":         {{620000, 628899}},
	"UnionPay": {},
	"Visa":     {{4}},
	// "Visa Electron":    {{4026}, {417500}, {4508}, {4844}, {4913}, {4917}},
	"Visa Electron": {},
}

var cardDigits = map[int][]string{
	12: {"Maestro"},
	13: {"Maestro", "Visa"},
	14: {"Diners Club", "Maestro", "Visa"},
	15: {"American Express", "Maestro", "Visa"},
	16: {"American Express", "Discover", "JCB", "Maestro", "Master", "UnionPay", "Visa", "Visa Electron"},
	17: {"Maestro", "Visa"},
	18: {"Maestro", "Visa"},
	19: {"Maestro", "Visa"},
}

func CheckIIN(cardNumber int64) string {
	cardNumberS := strconv.FormatInt(cardNumber, 10)
	cardLength := len(cardNumberS)
	typeList := cardDigits[cardLength]

	// loop through all card types in typeList
	for _, cardType := range typeList {
		// loop through all ranges in cardRanges
		for _, cardRange := range cardRanges[cardType] {
			// match cardRange with the first len(cardRange) digits of the cardNumber
			digits := len(strconv.Itoa(cardRange[0]))
			iinValue, _ := strconv.Atoi(cardNumberS[:digits])

			if iinValue >= cardRange[0] && iinValue <= cardRange[len(cardRange)-1] {
				return cardType
			}
		}
	}

	return "Unknown Card Type"
}
