func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return len(nums)
	}
	l := 1
	for r := 1; r < len(nums); r++ {
		if nums[r] != nums[r-1] {
			nums[l] = nums[r]
			l++
		}
	}
	return l
}
