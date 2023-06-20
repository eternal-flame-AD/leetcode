package mergesortedarray

func merge(nums1 []int, m int, nums2 []int, n int) {
	pSrc1 := m - 1
	pSrc2 := n - 1
	pDst := m + n - 1
	for pSrc1 >= 0 && pSrc2 >= 0 {
		if nums1[pSrc1] > nums2[pSrc2] {
			nums1[pDst] = nums1[pSrc1]
			pSrc1--
		} else {
			nums1[pDst] = nums2[pSrc2]
			pSrc2--
		}
		pDst--
	}
	for pSrc2 >= 0 {
		nums1[pDst] = nums2[pSrc2]
		pSrc2--
		pDst--
	}
	for pSrc1 >= 0 {
		nums1[pDst] = nums1[pSrc1]
		pSrc1--
		pDst--
	}
}
