package LongestCommonPrefix

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	length := 0
	for {
		ok := true
		for index := range strs {
			if len(strs[index]) == length || strs[index][length] != strs[0][length] {
				ok = false
				break
			}
		}
		if !ok {
			break
		}
		length++
	}
	return strs[0][:length]
}
