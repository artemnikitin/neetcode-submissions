func threeSum(nums []int) [][]int {
	result := [][]int{}
	sort.Ints(nums)

	countMap := map[int]int{}
	for _, v := range nums {
		countMap[v]++
	}

	for i := 0; i < len(nums); i++ {
		countMap[nums[i]]--
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			countMap[nums[j]]--
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			target := -(nums[i] + nums[j])
			if countMap[target] > 0 {
				result = append(result, []int{nums[i], nums[j], target})
			}
		}
		for j := i + 1; j < len(nums); j++ {
			countMap[nums[j]]++
		}
	}

	return result
}

