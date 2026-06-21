func isPalindrome(s string) bool {
	optS := removeNonAlphanumeric(strings.ToLower(s))
	s2 := reverse(optS)
	return s2 == optS
}

func removeNonAlphanumeric(s string) string {
	result := []rune{}
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			result = append(result, r)
		}
	}
	return string(result)
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
