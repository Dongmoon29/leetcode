package twoSum

func twoSum(nums []int, target int) []int {
	// 빈 맵을 만들어서 숫자를 저장하고, 타겟에서 숫자를 뺀 값이 맵에 있는지 확인한다맵
	numMaps := make(map[int]int)
	for index, num := range nums {
		// 타겟에서 숫자를 뺀 값을 구한다
		complement := target - num
		// 만약 맵에 값이 있다면, 그 값과 현재 인덱스를 리턴한다
		if answer, ok := numMaps[complement]; ok {
			return []int{answer, index}
		}
		// 맵에 없다면 현재 숫자와 인덱스를 저장한다
		numMaps[num] = index
	}

	return []int{}
}
