package utils_test

import (
	"credit-card-validator/utils"
	"testing"
)

func TestIsLuhnValid(t *testing.T) {
	testCases := []struct {
		num      int64
		expected bool
	}{
		// Test cases for valid Luhn numbers
		{45320151128336, true},
		{6011111111111117, true},
		{79927398713, true},
		{5105105105105100, true},

		// Test cases for invalid Luhn numbers
		{1234567890123456, false},
		{4242424242424241, false},
		{9876543210987654, false},
	}

	for _, testCase := range testCases {
		result := utils.IsLuhnValid(testCase.num)
		if result != testCase.expected {
			t.Errorf("Expected %v for number %d, but got %v", testCase.expected, testCase.num, result)
		}
	}
}
