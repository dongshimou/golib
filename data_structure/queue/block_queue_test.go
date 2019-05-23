package queue

import (
	"context"
	"log"
	"testing"
	"time"
)

func TestAll(t *testing.T) {

	q:= NewBQueue()

	go func() {
		for i:=0;i<100;i++{
			go func() {
				for j:=0;j<50;j++{
					q.Pop()
				}
			}()
		}
	}()


	go func() {

		for i:=0;i<100;i++{

			go func() {
				for j:=0;j<100;j++{
					q.Push(i*100+j)
				}
			}()
		}

	}()



	time.Sleep(time.Second*5)

	want:=100*100/2
	if q.Size()!=want{
		t.Error("want size=",want,"get size=",q.Size())
	}else{
		t.Log("want size=",want,"get size=",q.Size())
	}
}

func TestBlockPop(t *testing.T){

	q:= NewBQueue()

	ctx,cancel:=context.WithCancel(context.Background())
	go func() {
		for{
			if len(ctx.Done())!=0{
				return
			}
			log.Println(q.PopFront())
		}
	}()

	for i:=0;i<10;i++{
		time.Sleep(time.Second)
		q.Push(i)
	}
	cancel()

}
