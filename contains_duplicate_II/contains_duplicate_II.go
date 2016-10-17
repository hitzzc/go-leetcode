package contains_duplicate_II

func containsNearbyDuplicate(nums []int, k int) bool {
	m := map[int]int{}
	for i := range nums {
		if j, ok := m[nums[i]]; ok {
			if i-j <= k {
				return true
			}
		}
		m[nums[i]] = i
	}
	return false
}
