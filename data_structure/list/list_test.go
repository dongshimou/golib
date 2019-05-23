package list

import (
	"testing"
)

func TestAll(t *testing.T) {

	{
		list := New()

		list.PushBack(1)
		list.PopFront()

		want := 0

		if want != list.Size() {
			t.Error("want", want, "got", list.Size())
		} else {
			t.Log("want", want, "got", list.Size())
		}
	}

	{
		list := New()

		list.PushBack(1)
		list.PushBack(2)
		if list.Back() != 2 || list.Front() != 1 {
			t.Error(list.Back(), list.Front())
		} else {
			t.Log(list.Back(), list.Front())
		}
	}
	{
		list := New()

		list.PushBack(1)
		list.PushBack(2)
		list.PushFront(3)

		if list.Back() != 2 || list.Front() != 3 {
			t.Error(list.Back(), list.Front())
		} else {
			t.Log(list.Back(), list.Front())
		}
	}
}
