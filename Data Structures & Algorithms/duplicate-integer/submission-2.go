func hasDuplicate(nums []int) bool {
    if len(nums) == 1 {
        return false
    }
    numsMap := make(map[int]bool)
    for _, v := range nums {
        numsMap[v] = true
    }
    return len(nums) != len(numsMap)
}
