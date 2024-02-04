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
