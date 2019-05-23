package priority_queue

import (
	"sync"
)

type BPriorityQueue struct {
	PriorityQueue
	cond sync.Cond
	sync.Mutex
}

func NewBPriorityQueue() *BPriorityQueue {
	bpq := &BPriorityQueue{}
	bpq.cond = sync.Cond{L: bpq}
	return bpq
}

func (this *BPriorityQueue) Push(v interface{}) {
	this.Lock()
	defer this.Unlock()
	this.push(v)
	this.cond.Signal()
}

func (this *BPriorityQueue) Pop() interface{} {
	this.Lock()
	defer this.Unlock()
	for this.empty() {
		this.cond.Wait()
	}
	res := this.top()
	this.pop()
	return res
}
func (this *BPriorityQueue) Top() interface{} {
	this.Lock()
	defer this.Unlock()
	for this.empty() {
		this.cond.Wait()
	}
	return this.top()
}
func (this *BPriorityQueue) Size() int {
	this.Lock()
	defer this.Unlock()
	return this.size()
}
func (this *BPriorityQueue) Empty() bool {
	this.Lock()
	defer this.Unlock()
	return this.empty()
}
