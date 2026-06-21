func twoSum(numbers []int, target int) []int {
    first := 0
	last := len(numbers) - 1
	result := make([]int, 2)
	for first < last {
		if numbers[first] + numbers[last] == target {
            return []int{first + 1, last + 1}
        } else if numbers[first] + numbers[last] < target {
            first++
        } else {
            last--
        }
	}
	return result
}
