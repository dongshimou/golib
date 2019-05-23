package priority_queue

import (
	"log"
	"testing"
)

func TestAll(t *testing.T) {

	pq := NewPriorityQueue(func(a interface{}, b interface{}) bool {
		return a.(int) > b.(int)
	})
	pq.Push(1)
	pq.Push(3)
	pq.Push(2)
	pq.Push(10)
	log.Println(pq.Pop())
	log.Println(pq.Pop())
	log.Println(pq.Pop())
	pq.Push(3)
	log.Println(pq.Pop())
}
