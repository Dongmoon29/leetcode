package longestcommonprefix

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	result := ""
	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != c {
				return result
			}
		}
		result = result + string(c)
	}
	return result
}
