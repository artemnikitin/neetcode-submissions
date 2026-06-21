func maxSubArray(nums []int) int {
    maxSum := nums[0]
	currentSum := 0
	for _, v := range nums {
		currentSum = max(currentSum+v, v)
		maxSum = max(maxSum, currentSum)
	}
	return maxSum
}
