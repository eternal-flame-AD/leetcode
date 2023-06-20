package mergesortedarray

import "testing"

type TestCase struct {
	nums1 []int
	nums2 []int
	ret   []int
}

func TestMergeSortedArray(t *testing.T) {
	testCases := []TestCase{
		{
			nums1: []int{1, 2, 3, 0, 0, 0},
			nums2: []int{2, 5, 6},
			ret:   []int{1, 2, 2, 3, 5, 6},
		},
		{
			nums1: []int{1},
			nums2: []int{},
			ret:   []int{1},
		},
		{
			nums1: []int{0},
			nums2: []int{1},
			ret:   []int{1},
		},
	}

	for _, tc := range testCases {
		merge(tc.nums1, len(tc.nums1)-len(tc.nums2), tc.nums2, len(tc.nums2))
		for i := 0; i < len(tc.nums1); i++ {
			if tc.nums1[i] != tc.ret[i] {
				t.Errorf("merge(%v, %v, %v, %v) return %v, want %v", tc.nums1, len(tc.nums1)-len(tc.nums2), tc.nums2, len(tc.nums2), tc.nums1, tc.ret)
			}
		}
	}
}
