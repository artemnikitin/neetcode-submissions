func pivotIndex(nums []int) int {
	prefixSum := make([]int, len(nums))
	total := 0
	for i, num := range nums {
		total += num
		prefixSum[i] = total
	}

	for i := 0; i < len(nums); i++ {
		leftSum := 0
		if i > 0 {
			leftSum = prefixSum[i-1]
		}
		rightSum := 0
		if i < len(nums)-1 {
			rightSum = prefixSum[len(nums)-1] - prefixSum[i]
		}
		if leftSum == rightSum {
			return i
		}
	}

	return -1
}
