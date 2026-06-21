func groupAnagrams(strs []string) [][]string {
    if len(strs) == 1 {
        return [][]string{strs}
    }
    result := [][]string{}
    wordMap := map[[26]int][]string{}
    for _, v := range strs {
        temp := [26]int{}
        for _, c := range v {
            temp[c-'a']++
        }
        wordMap[temp] = append(wordMap[temp], v)
    }
    for _, v := range wordMap {
        result = append(result, v)
    }
    return result
}

