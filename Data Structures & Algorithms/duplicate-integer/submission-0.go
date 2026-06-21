func hasDuplicate(nums []int) bool {
    for i , v := range nums {
        for i2 , v2 := range nums {
            if v == v2 && i != i2 {
                return true
            }
        }
    }
    return false
}
