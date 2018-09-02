package elem

func removeElement(nums []int, val int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		if nums[l] == val {
			nums[l], nums[r] = nums[r], nums[l]
			r--
		} else {
			l++
		}
	}
	return r + 1
}
