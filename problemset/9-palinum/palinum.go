package palinum

import (
	"strconv"
)

func isPalindrome(x int) bool {
	target := strconv.FormatInt(int64(x), 10)
	for i := 0; i < len(target)/2; i++ {
		if target[i] != target[len(target)-1-i] {
			return false
		}
	}
	return true
}
