package palindromenumber

import (
	"strconv"
	"strings"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var result strings.Builder
	convertedNumber := strconv.Itoa(x)
	for i := len(convertedNumber) - 1; i >= 0; i-- {
		result.WriteString(string(convertedNumber[i]))
	}
	return convertedNumber == result.String()
}

func isPalindrome2(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	divisor := 1
	for temp := x; temp >= 10; temp /= 10 {
		divisor *= 10
	}

	for x > 0 {
		firstDigit := x / divisor
		lastDigit := x % 10

		if firstDigit != lastDigit {
			return false
		}

		x = (x % divisor) / 10
		divisor /= 100
	}
	return true
}
