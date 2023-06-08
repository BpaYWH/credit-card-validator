package utils

import (
	"fmt"
)

func IsLuhnValid(num int) bool {
	lastDigit := num % 10
	num /= 10
	sum := 0
	digit := 0x1
	curr := 0

	for num > 1 {
		curr = num % 10
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
