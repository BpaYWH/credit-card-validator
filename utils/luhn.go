package utils

import (
	"fmt"
)

func IsLuhnValid(num int64) bool {
	if num > 9999999999999999 {
		return false
	}

	lastDigit := int(num % 10)
	num /= 10
	sum := 0
	digit := 1
	var curr int

	for num > 1 {
		curr = int(num % 10)
		num /= 10

		curr *= digit + 1
		digit ^= 1
		if curr > 9 {
			sum += curr % 10
			curr /= 10
			sum += curr
		} else {
			sum += curr
		}
	}

	return 10-sum%10 == lastDigit
}

func LuhnTest() {
	// test cases for my Luhn Algoorithm
	// True
	fmt.Println(IsLuhnValid(79927398713))
	// False
	fmt.Println(IsLuhnValid(79927398714))
}
