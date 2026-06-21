func search(nums []int, target int) int {
    first := 0
    last := len(nums) - 1
    for first <= last {
        mid := first + (last - first) / 2
        if nums[mid] == target {
            return mid
        } else if nums[mid] < target {
            first++
        } else {
            last--
        }
    }
    return -1
}
