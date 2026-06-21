func reverseBits(n int) int {
	if n == 0 {
		return 0
	}
	result := 0
	for i := 0; i < 32; i++ {
		result = (result << 1) | (n & 1)
		n = n >> 1
	}
	return result
}
