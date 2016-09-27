package merge_sorted_array

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := m + n - 1; m > 0 && n > 0 && i >= 0; i-- {
		if nums1[m-1] > nums2[n-1] {
			nums1[i] = nums1[m-1]
			m--
		} else {
			nums1[i] = nums2[n-1]
			n--
		}
	}
	if m == 0 {
		for ; n > 0; n-- {
			nums1[n-1] = nums2[n-1]
		}
	}
	return
}
