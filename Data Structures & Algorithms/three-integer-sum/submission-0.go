func threeSum(nums []int) [][]int {
	result := [][]int{}
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		a := nums[i]
		if a > 0 {
			break
		}
		if i > 0 && a == nums[i-1] {
			continue
		}
		b := i + 1
		c := len(nums) - 1
		for b < c {
			sum := a + nums[b] + nums[c]
			if sum == 0 {
				result = append(result, []int{a, nums[b], nums[c]})
				b++
				c--
				for b < c && nums[b] == nums[b-1] {
					b++
				}
				for b < c && nums[c] == nums[c+1] {
					c--
				}
			} else if sum < 0 {
				b++
			} else {
				c--
			}
		}
	}

	return result
}

