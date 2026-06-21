type MinStack struct {
	array []int
	mins  []int
}

func Constructor() MinStack {
	return MinStack{
		array: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	if len(this.array) == 0 {
		this.mins = append(this.mins, val)
	} else {
		if val <= this.mins[len(this.mins)-1] {
			this.mins = append(this.mins, val)
		}
	}
	this.array = append(this.array, val)
}

func (this *MinStack) Pop() {
	top := this.array[len(this.array)-1]
	this.array = this.array[:len(this.array)-1]
	if top == this.mins[len(this.mins)-1] {
		this.mins = this.mins[:len(this.mins)-1]
	}
}

func (this *MinStack) Top() int {
	return this.array[len(this.array)-1]
}

func (this *MinStack) GetMin() int {
	return this.mins[len(this.mins)-1]
}
