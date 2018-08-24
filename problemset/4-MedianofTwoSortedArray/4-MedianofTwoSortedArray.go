package medoftwosortedarray

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	newarray := make([]int, 0)
	lp1 := 0
	lp2 := 0
	for lp1 != len(nums1) || lp2 != len(nums2) {
		switch {
		case lp1 < len(nums1) && (lp2 == len(nums2) || nums1[lp1] < nums2[lp2]):
			newarray = append(newarray, nums1[lp1])
			lp1++
		case lp2 < len(nums2) && (lp1 == len(nums1) || nums1[lp1] >= nums2[lp2]):
			newarray = append(newarray, nums2[lp2])
			lp2++
		}
	}
	med1 := len(newarray) / 2
	var med2 int
	if len(newarray)%2 == 1 {
		med2 = med1
	} else {
		med2 = med1 - 1
	}
	return (float64(newarray[med1]) + float64(newarray[med2])) / 2
}
