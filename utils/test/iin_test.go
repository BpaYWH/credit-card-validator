package utils_test

import (
	"credit-card-validator/utils"
	"testing"
)

func TestCheckIIN(t *testing.T) {
	testCases := []struct {
		cardNumber int64
		expected   string
	}{
		// Test cases for known card types
		{376864516055733, "American Express"},
		{379538547393720, "American Express"},
		{36447536450217, "Diners Club"},
		{36021094231150, "Diners Club"},
		{6573661517310143, "Discover"},
		{3566029174275502, "JCB"},
		{5062419955268600710, "Maestro"},
		{508145597284, "Maestro"},
		{5325018270012940, "Master"},
		{5488484848964413, "Master"},
		{5265383082736454, "Master"},
		{4494213742019510, "Visa"},

		// Test cases for unknown card type
		{1234567890123456, "Unknown Card Type"},
		{1111000011110000, "Unknown Card Type"},
		{9876543210987654, "Unknown Card Type"},
	}

	for _, testCase := range testCases {
		result := utils.CheckIIN(testCase.cardNumber)
		if result != testCase.expected {
			t.Errorf("Expected %s for card number %d, but got %s", testCase.expected, testCase.cardNumber, result)
		}
	}
}
