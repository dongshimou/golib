package list

import (
	"sync"
)

type BList struct {
	List
	sync.Mutex
	cond sync.Cond
}

func NewBList() *BList {
	bl := &BList{
		List: *New(),
	}
	bl.cond = sync.Cond{L: bl}
	return bl
}

func (this *BList) PushFront(v interface{}) {
	this.Lock()
	defer this.Unlock()
	this.push_front(v)
	this.cond.Signal()
}
func (this *BList) PushBack(v interface{}) {
	this.Lock()
	defer this.Unlock()
	this.push_back(v)
	this.cond.Signal()
}
func (this *BList) PopFront() {
	this.Lock()
	defer this.Unlock()
	for this.empty() {
		this.cond.Wait()
	}
	this.pop_front()
}
func (this *BList) PopBack() {
	this.Lock()
	defer this.Unlock()
	for this.empty() {
		this.cond.Wait()
	}
	this.pop_back()
}
func (this *BList) PopAndFront() interface{} {
	this.Lock()
	defer this.Unlock()
	for this.empty() {
		this.cond.Wait()
	}
	return this.PopAndFront()
}
func (this *BList) PopAndBack() interface{} {
	this.Lock()
	defer this.Unlock()
	for this.empty() {
		this.cond.Wait()
	}
	return this.PopAndBack()
}

func (this *BList) Front() interface{} {
	this.Lock()
	defer this.Unlock()
	for this.empty() {
		this.cond.Wait()
	}
	return this.front()
}
func (this *BList) Back() interface{} {
	this.Lock()
	defer this.Unlock()
	for this.empty() {
		this.cond.Wait()
	}
	return this.back()
}
func (this *BList) Size() int {
	this.Lock()
	defer this.Unlock()
	return this.size()
}
func (this *BList) Empty() bool {
	this.Lock()
	defer this.Unlock()
	return this.empty()
}
