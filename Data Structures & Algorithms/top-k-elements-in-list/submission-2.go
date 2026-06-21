func topKFrequent(nums []int, k int) []int {
    result := []int{}
	countMap := map[int]int{}
	for _, v := range nums {
		countMap[v]++
	}
	for i := 1; i <= k; i++ {
		number := 0
		biggest := 0
		for key, val := range countMap {
			if val > biggest {
				biggest = val
				number = key
			}
		}
		result = append(result, number)
		delete(countMap, number)
	}
	return result
}
