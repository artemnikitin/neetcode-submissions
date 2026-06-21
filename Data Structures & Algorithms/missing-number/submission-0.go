func missingNumber(nums []int) int {
	referenceSum := 0
	for i := 0; i <= len(nums); i++ {
		referenceSum += i
	}
	actualSum := 0
	for _, num := range nums {
		actualSum += num
	}
	return referenceSum - actualSum
}
