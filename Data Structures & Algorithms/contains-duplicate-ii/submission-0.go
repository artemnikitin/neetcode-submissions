func containsNearbyDuplicate(nums []int, k int) bool {
	seen := map[int]int{}
	l := 0
	for r := 0; r < len(nums); r++ {
		if _, ok := seen[nums[r]]; ok {
			return true
		}
		if r-l+1 > k {
			delete(seen, nums[l])
			l++
		}
		seen[nums[r]] = r
	}
	return false
}
