func characterReplacement(s string, k int) int {
    count := map[byte]int{}
    result := 0
    maxf := 0
    l := 0
    for r := 0; r < len(s); r++ {
        count[s[r]]++
        if count[s[r]] > maxf {
            maxf = count[s[r]]
        }
        if (r - l + 1) - maxf > k {
            count[s[l]]--
            l++
        }
        if r - l + 1 > result {
            result = r - l + 1
        }
    }
    return result
}
