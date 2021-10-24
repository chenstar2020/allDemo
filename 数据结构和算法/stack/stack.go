package stack

import "sync"

type Stack struct{
	array []interface{}
	size int
	lock sync.RWMutex
}

func NewStack()*Stack{
	return &Stack{
		array: make([]interface{}, 0),
		size: 0,
		lock: sync.RWMutex{},
	}
}

func (s *Stack) Push(v interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.array = append(s.array, v)
	s.size++
}

func (s *Stack) Pop() (v interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.IsEmpty(){
		return nil
	}

	v = s.array[s.size-1]
	s.array = s.array[0:s.size-1]
	s.size--

	return
}

func (s *Stack) Peek() (v interface{}) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	if s.IsEmpty(){
		return nil
	}

	return s.array[s.size - 1]
}

// SingleIncreasePush 单调递增栈（栈底 -> 栈顶 ： 大 -> 小）
func (s *Stack) SingleIncreasePush(v int) int{
	count := 0

	if s.IsEmpty() || s.Peek().(int) >= v {
		s.Push(v)
	}else{
		for ;s.Peek() != nil && s.Peek().(int) < v;{
			count++
			s.Pop()
		}
		s.Push(v)
	}

	return count
}

// SingleDecreasePush 单调递减栈（栈顶 -> 栈底 ： 小 -> 大）
func (s *Stack) SingleDecreasePush(v int) int{
	count := 0

	if s.IsEmpty() || s.Peek().(int) <= v {
		s.Push(v)
	}else{
		for ;!s.IsEmpty() && s.Peek().(int) > v;{
			count++
			s.Pop()
		}
		s.Push(v)
	}

	return count
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

type StackMethod interface {
	Push(v interface{})
	Pop()(v interface{})
	Peek()(v interface{})
	Size() int
	IsEmpty()bool
}

var _ StackMethod = (*Stack)(nil)