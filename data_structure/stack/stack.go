package stack

import (
	"github.com/dongshimou/golib/data_structure/list"
)

type Stack struct {
	data *list.List
}
func NewStack()*Stack{
	s:=&Stack{
		data: list.New(),
	}
	return s
}
func (this*Stack)Push(v interface{}){
	this.data.PushBack(v)
}
func (this*Stack)Pop(){
	this.data.PopBack()
}
func (this*Stack)PopTop()interface{}{
	res:=this.Top()
	this.Pop()
	return res
}
func(this*Stack)Top()interface{}{
	return this.data.Back()
}
func(this *Stack)Size()int{
	return this.data.Size()
}
func(this*Stack)Empty()bool{
	return this.data.Empty()
}