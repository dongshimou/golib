package queue

import (
	"sync"
)

type Pipe chan interface{}

const (
	DefaultPipeSize = 1
)

type SafeQueue struct {
	sync.Mutex
	Data     []interface{}
	pipePool sync.Pool
}

func New() *SafeQueue {
	sq := SafeQueue{
		Data: make([]interface{}, 0, DefaultPipeSize),
		pipePool: sync.Pool{New: func() interface{} {
			return make(Pipe, DefaultPipeSize)
		}},
	}
	return &sq
}

func (this *SafeQueue) Push(v interface{}) {
	this.push(v)
}
func (this *SafeQueue) AsyncPush(v interface{}) {
	pipe := this.pipePool.Get().(Pipe)
	pipe <- v
	go func() {
		defer this.pipePool.Put(pipe)
		this.push(<-pipe)
	}()
}
func (this *SafeQueue) AsyncPop() Pipe {
	pipe := make(Pipe, DefaultPipeSize)
	go func() {
		pipe <- this.pop
	}()
	return pipe
}

func (this *SafeQueue) Pop() interface{} {
	res := this.front()
	this.pop()
	return res
}
func (this *SafeQueue) Front() interface{} {
	return this.front()
}

func (this *SafeQueue) Size() int {
	return this.size()
}
func (this *SafeQueue) Empty() bool {
	return this.empty()
}

func (this *SafeQueue) push(v interface{}) {
	this.Lock()
	defer this.Unlock()
	this.Data = append(this.Data, v)
}
func (this *SafeQueue) pop() {
	this.Lock()
	defer this.Unlock()
	this.Data = this.Data[1:]
}
func (this *SafeQueue) front() interface{} {
	this.Lock()
	defer this.Unlock()
	return this.Data[len(this.Data)-1]
}
func (this *SafeQueue) size() int {
	this.Lock()
	defer this.Unlock()
	return len(this.Data)
}
func (this *SafeQueue) empty() bool {
	this.Lock()
	defer this.Unlock()
	return len(this.Data) == 0
}
