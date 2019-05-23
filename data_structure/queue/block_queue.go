package queue

import (
	"sync"
)

type BQueue struct {
	Queue
	sync.Mutex
	cond sync.Cond
}

func NewBQueue() *BQueue {
	sq := BQueue{
		Queue: *NewQueue(),
	}
	sq.cond = sync.Cond{L: &sq}
	return &sq
}

func (this *BQueue) Push(v interface{}) {

	this.Lock()
	defer this.Unlock()
	this.Queue.Push(v)
	this.cond.Signal()
}
func (this *BQueue) Pop() {
	this.Lock()
	defer this.Unlock()
	for this.Queue.Empty() {
		this.cond.Wait()
	}
	this.Queue.Pop()
}
func (this *BQueue) PopAndFront() interface{} {
	this.Lock()
	defer this.Unlock()
	for this.Queue.Empty() {
		this.cond.Wait()
	}
	res := this.Queue.PopAndFront()
	return res
}
func (this *BQueue) AsyncPush(v interface{}) {
	go func() {
		this.Push(v)
	}()
}
func (this *BQueue) AsyncPop() {
	go func() {
		this.Pop()
	}()
}
func (this *BQueue) Front() interface{} {
	this.Lock()
	defer this.Unlock()
	for this.Queue.Empty() {
		this.cond.Wait()
	}
	return this.Queue.Front()
}

func (this *BQueue) Size() int {
	this.Lock()
	defer this.Unlock()
	return this.Queue.Size()
}

func (this *BQueue) Empty() bool {
	this.Lock()
	defer this.Unlock()
	return this.Queue.Empty()
}
