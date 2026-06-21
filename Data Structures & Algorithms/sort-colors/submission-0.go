func sortColors(nums []int) {
    if len(nums) == 1 {
		return
	}

	count0 := 0
	count1 := 0
	count2 := 0

	for _, num := range nums {
		switch num {
		case 0:
			count0++
		case 1:
			count1++
		case 2:
			count2++
		}
	}

	for i := 0; i < count0; i++ {
		nums[i] = 0
	}
	for i := count0; i < count0+count1; i++ {
		nums[i] = 1
	}
	for i := count0 + count1; i < len(nums); i++ {
		nums[i] = 2
	}
}
