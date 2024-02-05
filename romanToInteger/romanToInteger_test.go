package romantointeger

import "testing"

func TestRomanToInt(t *testing.T) {
	input := "MCMXCIV"
	result := romanToInt(input)

	expectedValue := 1994

	if result != expectedValue {
		t.Errorf("Expected %v, but got %v", expectedValue, result)
	}
}
