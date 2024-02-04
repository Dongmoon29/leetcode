package palindromenumber

import "testing"

func TestIsPalindrome(t *testing.T) {
	num := 121
	result := isPalindrome(num)
	if result != true {
		t.Errorf("Expected true, but got %v", result)
	}
}

func TestIsPalindrome2(t *testing.T) {
	num := 10
	result := isPalindrome(num)
	if result == true {
		t.Errorf("Expected true, but got %v", result)
	}
}
