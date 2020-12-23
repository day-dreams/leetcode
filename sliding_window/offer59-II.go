package sliding_window

type MaxQueue struct {
	max   int
	nodes []int
}

func Constructor() MaxQueue {
	return MaxQueue{
		max:   -1,
		nodes: nil,
	}

}

func (this *MaxQueue) Max_value() int {
	if this.max < 0 {
		return -1
	}
	return this.nodes[this.max]
}

func (this *MaxQueue) Push_back(value int) {
	if this.max == -1 {
		this.max = len(this.nodes)
	} else if this.nodes[this.max] < value {
		this.max = len(this.nodes)
	}
	this.nodes = append(this.nodes, value)

}

func (this *MaxQueue) Pop_front() int {
	if len(this.nodes) == 1 {
		this.max = -1
		rv := this.nodes[0]
		this.nodes = nil
		return rv
	} else if this.max < 0 {
		return -1
	} else if this.max == 0 {
		// 重新找max
		rv := this.nodes[0]
		this.nodes = this.nodes[1:]

		max := -1000000
		for i, node := range this.nodes {
			if node > max {
				this.max = i
				max = node
			}
		}

		return rv
	} else {
		this.max--
		rv := this.nodes[0]
		this.nodes = this.nodes[1:]
		return rv
	}

}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
