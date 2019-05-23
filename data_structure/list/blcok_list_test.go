package list

import (
	"testing"
	"time"
)

func TestBL(t *testing.T) {

	l := NewBList()

	go func() {
		for i := 0; i < 100; i++ {
			go func() {
				for j := 0; j < 50; j++ {
					if j&1 == 1 {
						l.PopBack()
					} else {
						l.PopFront()
					}
				}
			}()
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			go func() {
				for j := 0; j < 100; j++ {
					if j&1 == 1 {
						l.PushBack(i*100 + j)
					} else {
						l.PushFront(i*100 + j)
					}
				}
			}()
		}
	}()

	time.Sleep(time.Second * 5)

	want := 100 * 100 / 2
	if l.Size() != want {
		t.Error("want size=", want, "get size=", l.Size())
	} else {
		t.Log("want size=", want, "get size=", l.Size())
	}

}
