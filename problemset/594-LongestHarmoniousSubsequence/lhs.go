package lhs

import (
	"sort"
)

type numRecord struct {
	index          int
	occurencecount int
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func findLHS(nums []int) int {
	sorted := make([]int, len(nums))
	numminplace := make(map[int]*numRecord, 0)
	nummaxplace := make(map[int]*numRecord, 0)
	copy(sorted[:], nums)
	sort.Ints(sorted)
	for index, val := range nums {
		if _, ok := numminplace[val]; !ok {
			numminplace[val] = new(numRecord)
			numminplace[val].index = index
			numminplace[val].occurencecount = 1
			nummaxplace[val] = new(numRecord)
			nummaxplace[val].index = index
			nummaxplace[val].occurencecount = 1
			continue
		}
		nummaxplace[val] = &numRecord{
			index:          index,
			occurencecount: nummaxplace[val].occurencecount + 1,
		}
	}
	maxlength := 0
	for i := 0; i < len(sorted)-1; i++ {
		if sorted[i+1]-1 == sorted[i] {
			maxplace := nummaxplace[sorted[i]].occurencecount + nummaxplace[sorted[i+1]].occurencecount
			maxlength = maxInt(maxlength, maxplace)
		}
	}
	return maxlength
}
