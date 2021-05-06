package offer

type CQueue struct {
	primary, second []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.primary = append(this.primary, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.second) != 0 {
		val := this.second[len(this.second)-1]
		this.second = this.second[0 : len(this.second)-1]
		return val
	}

	if len(this.primary) != 0 {
		for len(this.primary) != 0 {
			this.second = append(this.second, this.primary[len(this.primary)-1])
			this.primary = this.primary[0 : len(this.primary)-1]
		}
		val := this.second[len(this.second)-1]
		this.second = this.second[0 : len(this.second)-1]
		return val
	}
	return -1
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
