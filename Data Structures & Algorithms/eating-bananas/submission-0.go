func minEatingSpeed(piles []int, h int) int {
    first := 1
    last := 0
    for _, v := range piles {
        if v > last {
            last = v
        }
    }

    result := 0
    for first <= last {
        mid := first + (last - first)/2
        total := 0
        for _, p := range piles {
            total += int(math.Ceil(float64(p)/float64(mid)))
        }
        if total <= h {
            result = mid
            last = mid - 1
        } else {
            first = mid + 1
        }
    }
    return result
}
