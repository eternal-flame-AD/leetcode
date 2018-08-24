package twosum

func inArray(nums *[]int, target int, startindex int) int {
	for index := startindex; index < len(*nums); index++ {
		if (*nums)[index] == target {
			return index
		}
	}
	return 0
}

func twoSum(nums []int, target int) []int {
	for index, t := range nums {
		if index2 := inArray(&nums, target-t, index+1); index2 != 0 {
			return []int{index, index2}
		}
	}
	return []int{}
}
