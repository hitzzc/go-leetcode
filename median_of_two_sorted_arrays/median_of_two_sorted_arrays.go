package median_of_two_sorted_arrays

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	total := l1 + l2
	if total%2 == 1 {
		return float64(findKthNum(nums1, nums2, total/2+1))
	} else {
		return 0.5 * float64(findKthNum(nums1, nums2, total/2)+findKthNum(nums1, nums2, total/2+1))
	}
}

func findKthNum(nums1 []int, nums2 []int, k int) (val int) {
	l1, l2 := len(nums1), len(nums2)
	if l1 > l2 {
		return findKthNum(nums2, nums1, k)
	}
	if l1 == 0 {
		return nums2[k-1]
	}
	if k == 1 {
		if nums1[0] < nums2[0] {
			return nums1[0]
		} else {
			return nums2[0]
		}
	}
	var pa, pb int
	if l1 < k/2 {
		pa = l1
	} else {
		pa = k / 2
	}
	pb = k - pa
	if nums1[pa-1] < nums2[pb-1] {
		return findKthNum(nums1[pa:], nums2, k-pa)
	} else if nums1[pa-1] > nums2[pb-1] {
		return findKthNum(nums1, nums2[pb:], k-pb)
	} else {
		return nums1[pa-1]
	}
}
