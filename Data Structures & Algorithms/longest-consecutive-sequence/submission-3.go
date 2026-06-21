func longestConsecutive(nums []int) int {
    if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
    count := map[int]bool{}
	for _, v := range nums {
		count[v] = true
	}
	longest := 0
	for i := 0; i < len(nums); i++ {
		_, exists := count[nums[i]-1]
		if !exists {
			temp := nums[i] + 1
			length := 1
			_, exists = count[temp]
			for exists {
				temp++
				length++
				_, exists = count[temp]
			}
			if length > longest {
				longest = length
			}
		}
	}
	return longest
}
