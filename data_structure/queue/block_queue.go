package queue

import (
	"sync"
)

type BQueue struct {
	Queue
	null     sync.Mutex
	cond     sync.Cond
	exist    bool
}
func NewBQueue() *BQueue {
	sq := BQueue{
		Queue:*NewQueue(),
	}
	sq.null=sync.Mutex{}
	sq.cond=sync.Cond{L:&sq.null}
	return &sq
}

func (this*BQueue)Push(v interface{}) {
	this.cond.L.Lock()
	defer this.cond.L.Unlock()
	this.Queue.Push(v)
	this.cond.Signal()
	this.exist = true
}
func (this *BQueue)Pop() {
	this.cond.L.Lock()
	defer this.cond.L.Unlock()
	for !this.exist {
		this.cond.Wait()
	}
	this.Queue.Pop()
	if this.Queue.Size() == 0 {
		this.exist = false
	}
}
func (this*BQueue)PopFront()interface{} {
	this.cond.L.Lock()
	defer this.cond.L.Unlock()
	for !this.exist {
		this.cond.Wait()
	}
	res := this.Queue.PopFront()
	if this.Queue.Empty(){
		this.exist=false
	}
	return res
}
func (this *BQueue)AsyncPush(v interface{}){
	go func() {
		this.Push(v)
	}()
}
func (this *BQueue)AsyncPop(){
	go func() {
		this.Pop()
	}()
}
func (this*BQueue)Front()interface{} {
	this.cond.L.Lock()
	defer this.cond.L.Unlock()
	for !this.exist {
		this.cond.Wait()
	}
	return this.Queue.Front()
}

func (this*BQueue)Size()int{
	this.cond.L.Lock()
	defer this.cond.L.Unlock()
	return this.Queue.Size()
}

func (this *BQueue)Empty()bool{
	this.cond.L.Lock()
	defer this.cond.L.Unlock()
	return this.Queue.Empty()
}
