package praxis

type Element struct {
	Val  interface{}
	next *Element
}

type Stack struct {
	ele *Element
}

func (this *Stack) Pop() {
	if this.ele != nil {
		ele := this.ele
		this.ele = ele.next
		ele.next = nil
	}
}

func (this *Stack) Push(val interface{}) {
	ele := &Element{val, nil}
	if this.ele == nil {
		this.ele = ele
	} else {
		ele.next = this.ele
		this.ele = ele
	}
}

func (this *Stack) Peek() interface{} {
	if this.ele == nil {
		return nil
	}
	return this.ele.Val
}

func (this *Stack) clear() {
	for this.ele != nil {
		ele := this.ele
		ele.next = nil
		this.ele = this.ele.next
	}
}
