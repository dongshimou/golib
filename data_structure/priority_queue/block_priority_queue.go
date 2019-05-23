package priority_queue

import (
	"sync"
)

type BPriorityQueue struct {

	PriorityQueue

	cond  sync.Cond
	mutex sync.Mutex
	exist bool
}

func NewBPriorityQueue()*BPriorityQueue{
	bpq:=&BPriorityQueue{}
	bpq.mutex=sync.Mutex{}
	bpq.cond=sync.Cond{L:&bpq.mutex}
	return bpq
}

func (this *BPriorityQueue)Push(v interface{}){
	this.cond.L.Lock()
	defer this.cond.L.Unlock()
	this.push(v)
	this.cond.Signal()
	this.exist=true
}

func (this *BPriorityQueue)Pop()interface{}{
	this.cond.L.Lock()
	defer this.cond.L.Unlock()
	for !this.exist{
		this.cond.Wait()
	}
	res:=this.top()
	this.pop()
	if this.empty(){
		this.exist=false
	}
	return res
}
func (this *BPriorityQueue)Top()interface{}{
	this.cond.L.Lock()
	defer this.cond.L.Unlock()
	for !this.exist{
		this.cond.Wait()
	}
	return this.top()
}
func (this *BPriorityQueue)Size()int{
	this.cond.L.Lock()
	defer this.cond.L.Unlock()
	return this.size()
}
func (this *BPriorityQueue)Empty()bool{
	this.cond.L.Lock()
	defer this.cond.L.Unlock()
	return this.empty()
}