func twoSum(nums []int, target int) []int {
    numsMap := map[int]int{}
    for i, v := range nums {
        numsMap[v] = i
    }
    for i, v := range nums {
        j, ok := numsMap[target-v]
        if ok && j != i {
            return []int{i, j}
        }
    }
    return []int{}
}
