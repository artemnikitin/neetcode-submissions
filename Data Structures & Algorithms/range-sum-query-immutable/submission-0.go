type NumArray struct {
	prefixSum []int
}

func Constructor(nums []int) NumArray {
	sums := make([]int, len(nums))
	total := 0
	for i, num := range nums {
		total += num
		sums[i] = total
	}
	return NumArray{
		prefixSum: sums,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	l := 0
	if left > 0 {
		l = this.prefixSum[left-1]
	}
	return this.prefixSum[right] - l
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */