package lolongestsubstring

func lengthOfLongestSubstring(s string) int {
	max := 0
	for startindex := range s {
		existingcharacters := make(map[byte]bool)
		var endindex int
		for endindex = startindex; endindex < len(s); endindex++ {
			if _, ok := existingcharacters[s[endindex]]; !ok {
				existingcharacters[s[endindex]] = true
			} else {
				break
			}
		}
		if endindex-startindex > max {
			max = endindex - startindex
		}
	}
	return max
}
