package threesum

import (
	"sort"
)

func inIntArray(nums []int, target, start, stop int) bool {
	for i := start; i <= stop; i++ {
		if nums[i] == target {
			return true
		}
	}
	return false
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	zerol := -1
	zeror := -1
	for index, val := range nums {
		if val == 0 {
			zeror = index
			if zerol == -1 {
				zerol = index
			}
		}
		if val < 0 && index+1 != len(nums) && nums[index+1] > 0 {
			zerol = index + 1
			zeror = index
		}
		if val > 0 {
			break
		}
	}
	if zerol == -1 {
		return res
	}

	// 0 0 0
	if zeror-zerol >= 2 {
		res = append(res, []int{0, 0, 0})
	}

	// + - 0
	if zeror-zerol >= 0 {
		for p := 0; p < zerol; p++ {
			if inIntArray(nums, -nums[p], zeror+1, len(nums)-1) {
				res = append(res, []int{-nums[p], 0, nums[p]})
			}
			for p+1 < len(nums) && nums[p+1] == nums[p] {
				p++
			}
		}
	}

	// + + -
	for l1 := 0; l1 < zerol; l1++ {
		for l2 := l1 + 1; l2 < zerol; l2++ {
			if inIntArray(nums, 0-nums[l1]-nums[l2], zeror+1, len(nums)-1) {
				res = append(res, []int{0 - nums[l1] - nums[l2], nums[l2], nums[l1]})
			}
			for l2+1 < len(nums) && nums[l2+1] == nums[l2] {
				l2++
			}
		}
		for l1+1 < len(nums) && nums[l1+1] == nums[l1] {
			l1++
		}
	}

	//+ - -
	for r1 := zeror + 1; r1 < len(nums); r1++ {
		for r2 := r1 + 1; r2 < len(nums); r2++ {
			if inIntArray(nums, 0-nums[r1]-nums[r2], 0, zerol-1) {
				res = append(res, []int{nums[r2], nums[r1], 0 - nums[r2] - nums[r1]})
			}
			for r2+1 < len(nums) && nums[r2+1] == nums[r2] {
				r2++
			}
		}
		for r1+1 < len(nums) && nums[r1+1] == nums[r1] {
			r1++
		}
	}
	return res
}
