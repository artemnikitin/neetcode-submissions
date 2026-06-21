func groupAnagrams(strs []string) [][]string {
    if len(strs) == 1 {
        return [][]string{strs}
    }
    result := [][]string{}
    wordMap := map[string][]string{}
    for _, v := range strs {
        wordMap[sortString(v)] = append(wordMap[sortString(v)], v)
    }
    for _, v := range wordMap {
        result = append(result, v)
    }
    return result
}

func sortString(s string) string {
    chars := strings.Split(s, "")
    sort.Strings(chars)
    return strings.Join(chars, "")
}
