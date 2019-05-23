package stack

import (
	"sync"
)

type BStack struct {
	Stack
	null sync.Mutex
	cond sync.Cond
	exist bool
}

func NewBStack()*BStack{
	bs:=BStack{
		Stack:*NewStack(),
	}
	bs.null=sync.Mutex{}
	bs.cond=sync.Cond{L:&bs.null}
	return &bs
}
func (this*BStack)Push(v interface{}){
	this.cond.L.Lock()
	defer this.cond.L.Unlock()
	this.Stack.Push(v)
	this.cond.Signal()
	this.exist=true
}
func (this*BStack)Pop(){
	this.cond.L.Lock()
	defer this.cond.L.Unlock()
	for !this.exist {
		this.cond.Wait()
	}
	this.Stack.Pop()
	if this.Stack.Size() == 0 {
		this.exist = false
	}
}
func (this *BStack)AsyncPush(v interface{}){
	go func() {
		this.Push(v)
	}()
}
func (this *BStack)AsyncPop(){
	go func() {
		this.Pop()
	}()
}
func (this*BStack)PopTop()interface{} {
	this.cond.L.Lock()
	defer this.cond.L.Unlock()
	for !this.exist {
		this.cond.Wait()
	}
	res := this.Stack.PopTop()
	if this.Stack.Empty(){
		this.exist=false
	}
	return res
}
func (this*BStack)Top()interface{} {
	this.cond.L.Lock()
	defer this.cond.L.Unlock()
	for !this.exist {
		this.cond.Wait()
	}
	return this.Stack.Top()
}

func (this*BStack)Size()int{
	this.cond.L.Lock()
	defer this.cond.L.Unlock()
	return this.Stack.Size()
}

func (this *BStack)Empty()bool{
	this.cond.L.Lock()
	defer this.cond.L.Unlock()
	return this.Stack.Empty()
}
