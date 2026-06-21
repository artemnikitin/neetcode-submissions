func topKFrequent(nums []int, k int) []int {
	countMap := map[int]int{}
	for _, v := range nums {
		countMap[v]++
	}
    frequency := make([][]int, len(nums) + 1)
	for key, val := range countMap {
		frequency[val] = append(frequency[val], key)
	}
	result := []int{}
    for i := len(frequency) - 1; i > 0; i-- {
        for _, v := range frequency[i] {
            result = append(result, v)
            if len(result) == k {
                return result
            }
        }
    }
	return result
}
