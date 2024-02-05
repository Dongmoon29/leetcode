package longestcommonprefix

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	input := []string{"ab", "a"}
	expectedOutput := "fl"

	result := longestCommonPrefix(input)
	if result != expectedOutput {
		t.Errorf("expected %v but got %v", expectedOutput, result)
	}
}
