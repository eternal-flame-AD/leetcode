package container

func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	max := 0
	for l < r {
		if height[l] < height[r] {
			if height[l]*(r-l) > max {
				max = height[l] * (r - l)
			}
			l++
		} else {
			if height[r]*(r-l) > max {
				max = height[r] * (r - l)
			}
			r--
		}
	}
	return max
}
