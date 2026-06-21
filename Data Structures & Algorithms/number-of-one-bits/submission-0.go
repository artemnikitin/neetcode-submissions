func hammingWeight(n int) int {
	result := 0
	for n != 0 {
		if n&1 == 1 {
			result++
		}
		n = n >> 1
	}
	return result
}
