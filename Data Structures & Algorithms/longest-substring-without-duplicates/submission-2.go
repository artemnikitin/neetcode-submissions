func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	set := make(map[byte]bool)
	maxL := 0
	l := 0
	for r := 0; r < len(s); r++ {
		for set[s[r]] {
			delete(set, s[l])
			l++
		}
		set[s[r]] = true
		if r-l+1 > maxL {
			maxL = r - l + 1
		}
	}
	return maxL
}
