package utils

func IsLuhnValid(num int64) bool {
	sum := 0
	digit := 0
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

	return sum%10 == 0
}
