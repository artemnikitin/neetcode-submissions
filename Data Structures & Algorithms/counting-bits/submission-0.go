func countBits(n int) []int {
	if n == 0 {
		return []int{0}
	}

	result := []int{}
	result = append(result, 0)
	for i := 1; i <= n; i++ {
		ones := 0
		n := i
		for n != 0 {
			if n&1 == 1 {
				ones++
			}
			n = n >> 1
		}
		result = append(result, ones)
	}

	return result
}
