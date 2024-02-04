package twoSum

import "testing"

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	if result[0] != 0 || result[1] != 1 {
		t.Errorf("Expected [0, 1], but got %v", result)
	}
}
