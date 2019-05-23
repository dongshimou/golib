package queue

import (
	"github.com/dongshimou/golib/data_structure/list"
)

type Queue struct {
	data *list.List
}
func NewQueue()*Queue{
	q:=&Queue{
		data: list.New(),
	}
	return q
}
func (this *Queue)Push(v interface{}) {
	this.data.PushBack(v)
}
func (this *Queue)Pop(){
	this.data.PopFront()
}
func (this *Queue)PopFront()interface{}{
	res:=this.Front()
	this.Pop()
	return res
}
func (this* Queue)Front()interface{}{
	return this.data.Front()
}
func(this*Queue)Size()int{
	return this.data.Size()
}
func (this*Queue)Empty()bool{
	return this.data.Empty()
}
