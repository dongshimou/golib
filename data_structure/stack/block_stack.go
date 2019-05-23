package stack

import (
	"sync"
)

type BStack struct {
	Stack
	sync.Mutex
	cond sync.Cond
}

func NewBStack() *BStack {
	bs := &BStack{
		Stack: *NewStack(),
	}
	bs.cond = sync.Cond{L: bs}
	return bs
}
func (this *BStack) Push(v interface{}) {
	this.cond.L.Lock()
	defer this.cond.L.Unlock()
	this.Stack.Push(v)
	this.cond.Signal()
}
func (this *BStack) Pop() {
	this.cond.L.Lock()
	defer this.cond.L.Unlock()
	for this.Stack.Empty() {
		this.cond.Wait()
	}
	this.Stack.Pop()
}
func (this *BStack) AsyncPush(v interface{}) {
	go func() {
		this.Push(v)
	}()
}
func (this *BStack) AsyncPop() {
	go func() {
		this.Pop()
	}()
}
func (this *BStack) PopAndTop() interface{} {
	this.cond.L.Lock()
	defer this.cond.L.Unlock()
	for this.Stack.Empty() {
		this.cond.Wait()
	}
	res := this.Stack.PopAndTop()
	return res
}
func (this *BStack) Top() interface{} {
	this.cond.L.Lock()
	defer this.cond.L.Unlock()
	for this.Stack.Empty() {
		this.cond.Wait()
	}
	return this.Stack.Top()
}

func (this *BStack) Size() int {
	this.cond.L.Lock()
	defer this.cond.L.Unlock()
	return this.Stack.Size()
}

func (this *BStack) Empty() bool {
	this.cond.L.Lock()
	defer this.cond.L.Unlock()
	return this.Stack.Empty()
}
