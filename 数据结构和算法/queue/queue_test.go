package queue

import "testing"

func TestNewQueue(t *testing.T) {
	NewQueue()
}

func TestQueue_Add(t *testing.T) {
	q := NewQueue()
	q.Add("aaa")
	q.Add("bbb")
	t.Log(q.Len())
}

func TestQueue_Remove(t *testing.T) {
	q := NewQueue()
	q.Add("aaa")
	q.Add("bbb")
	t.Log(q.Len())

	q.Remove()
	t.Log(q.Len())
	q.Remove()
	t.Log(q.Len())
	q.Remove()
	t.Log(q.Len())
}