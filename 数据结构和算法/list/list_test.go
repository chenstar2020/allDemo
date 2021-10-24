package list

import (
	"testing"
)

func NewList()*List{
	l := New()
	l.PushFront("a")
	l.PushFront("b")
	l.PushFront("c")
	return l
}

func TestList_PushFront(t *testing.T) {
	l := NewList()
	l.PushFront("d")
	l.Print()
}

func TestList_PushBack(t *testing.T) {
	l := NewList()
	l.PushBack("d")
	l.Print()
}

func TestList_Front(t *testing.T) {
	l := NewList()
	t.Log(l.Front().Data.(string))
}

func TestList_Back(t *testing.T) {
	l := NewList()
	t.Log(l.Back().Data.(string))
}

func TestList_Init(t *testing.T) {
	l := NewList()
	l.Print()
	l.Init()
	l.Print()
}

func TestList_Len(t *testing.T) {
	l := NewList()
	l.PushFront("aaa")
	t.Log(l.Len())
}

func TestList_InsertBefore(t *testing.T) {
	t.Run("fun1", func(t *testing.T) {
		l := New()
		a := l.PushFront("a")
		l.PushFront("b")
		c := l.PushFront("c")
		l.Print()
		l.InsertBefore("cbefore", c)
		l.InsertBefore("abefore", a)
		l.Print()
	})

	t.Run("fun2", func(t *testing.T) {
		l := New()
		head := l.Front()
		l.InsertBefore("head", head)
		l.Print()
	})
}

func TestList_InsertAfter(t *testing.T) {
	t.Run("fun1", func(t *testing.T) {
		l := New()
		a := l.PushFront("a")
		l.PushFront("b")
		c := l.PushFront("c")
		l.Print()
		l.InsertAfter("aafter", a)
		l.InsertAfter("cafter", c)
		l.Print()
	})

	t.Run("func2", func(t *testing.T) {
		l := New()
		head := l.Front()
		l.InsertAfter("head", head)
		l.Print()
	})
}

func TestList_MoveToFront(t *testing.T) {
	l := NewList()
	move := l.PushBack("move")
	l.Print()
	l.MoveToFront(move)
	l.Print()
}

func TestList_MoveToBack(t *testing.T) {
	l := NewList()
	move := l.PushFront("move")
	l.Print()
	l.MoveToBack(move)
	l.Print()
}

func TestList_MoveBefore(t *testing.T) {
	l := NewList()
	move := l.PushFront("move")
	tag := l.PushFront("tag")
	l.Print()
	l.MoveBefore(move, tag)
	l.Print()
}

func TestList_MoveAfter(t *testing.T) {
	l := NewList()
	move := l.PushBack("move")
	//tag := l.PushBack("tag")
	l.Print()
	l.MoveAfter(move, move)
	l.Print()
}

func TestList_PushBackList(t *testing.T) {
	l1 := New()
	l1.PushBack("a")
	l1.PushBack("b")
	l1.PushBack("c")
	l2 := New()
	l2.PushBack("d")
	l2.PushBack("e")
	l2.PushBack("f")
	l1.PushBackList(l2)
	l1.Print()
}

func TestList_PushFrontList(t *testing.T) {
	l1 := New()
	l1.PushBack("a")
	l1.PushBack("b")
	l1.PushBack("c")
	l2 := New()
	l2.PushBack("d")
	l2.PushBack("e")
	l2.PushBack("f")
	l1.PushFrontList(l2)
	l1.Print()
	t.Log(l1.Len())
}