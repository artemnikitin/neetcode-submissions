func subsets(nums []int) [][]int {
	result := [][]int{{}}
	for _, v := range nums {
		length := len(result)
		for i := 0; i < length; i++ {
			newSubset := append([]int{}, result[i]...)
			newSubset = append(newSubset, v)
			result = append(result, newSubset)
		}
	}
	return result
}