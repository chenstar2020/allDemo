package stack

import "testing"

func NewStack()*Stack{
	s := NewStack()
	return s
}

func TestStack_Push(t *testing.T) {
	s := NewStack()
	s.Push("abc")
	s.Push("def")
	t.Log(s.Size())
}

func TestStack_Pop(t *testing.T) {
	s := NewStack()
	s.Push("abc")
	s.Push("def")
	t.Log(s.Pop())
	t.Log(s.Pop())
	t.Log(s.Pop())
}

func TestStack_Peek(t *testing.T) {
	s := NewStack()
	s.Push("abc")
	s.Push("def")
	t.Log(s.Peek())
}