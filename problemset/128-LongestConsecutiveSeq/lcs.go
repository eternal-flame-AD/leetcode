package lcs

import (
	"sort"
)

func longestConsecutive(nums []int) int {
	if 0 == len(nums) {
		return 0
	}
	sort.Ints(nums)
	maxc := 1
	c := 1
	lastn := nums[0]
	for i := 1; i < len(nums); i++ {
		switch nums[i] - lastn {
		case 0:
		case 1:
			c++
		default:
			c = 1
		}
		if c > maxc {
			maxc = c
		}
		lastn = nums[i]
	}
	return maxc
}
