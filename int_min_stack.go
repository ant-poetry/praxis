package praxis

type MinStack struct {
	s Stack
	m Stack
}

/** initialize your data structure here. */
//func Constructor() MinStack {
//	return MinStack{}
//}

func (this *MinStack) Push(x int) {
	this.s.Push(x)
	if this.m.Peek() == nil {
		this.m.Push(x)
	} else {
		if x <= this.GetMin() {
			this.m.Push(x)
		}
	}
}

func (this *MinStack) Pop() {
	peek := this.s.Peek()
	this.s.Pop()
	if peek != nil {
		if peek == this.m.Peek() {
			this.m.Pop()
		}
	}
}

func (this *MinStack) Top() int {
	in := this.s.Peek()
	if in == nil {
		return 0
	}
	return in.(int)
}

func (this *MinStack) GetMin() int {
	in := this.m.Peek()
	if in == nil {
		return 0
	}
	return in.(int)
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
