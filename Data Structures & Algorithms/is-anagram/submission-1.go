func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    charMapS := make(map[rune]int)
    for _, v := range s {
        charMapS[v]++
    }
    charMapT := make(map[rune]int)
    for _, v := range t {
        charMapT[v]++
    }
    if len(charMapS) != len(charMapT) {
        return false
    }
    for k, v := range charMapS {
        if charMapT[k] != v {
            return false
        }
    }
    return true
}
