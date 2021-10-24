package queue

import "sync"

type Queue struct {
	array []int
	size int
	lock sync.Mutex
}

func NewQueue()*Queue{
	return &Queue{
		array: make([]int, 0),
		size: 0,
		lock: sync.Mutex{},
	}
}

func (q *Queue) Add(v int) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.array = append(q.array, v)
	q.size++
}

func (q *Queue) Remove() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.size == 0{
		return 0
	}

	v := q.array[q.size]

	for i := 1; i < q.size;i++{
		q.array[i-1] = q.array[i]
	}
	q.array = q.array[0:q.size -1]

	return v
}

var _ QueueMethod = (*Queue)(nil)

type QueueMethod interface {
	Add(v int)
	Remove()int
}


//环形队列
type MyCircularQueue struct {
	array []int   //使用数组存储队列
	front int     //队首指针下标
	rear int      //队尾指针下标
	size int      //队列长度
}


func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		array: make([]int, k + 1),
		front: 0,
		rear: 0,
		size: k + 1,
	}
}


func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull(){
		return false
	}

	this.array[this.front] = value
	this.front = (this.front + 1) % this.size

	return true
}


func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty(){
		return false
	}

	this.rear = (this.rear + 1) % this.size
	return true
}


func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty(){
		return -1
	}

	return this.array[(this.size + this.front - 1) % this.size]
}


func (this *MyCircularQueue) Front() int {
	if this.IsEmpty(){
		return -1
	}
	return this.array[this.rear]
}


func (this *MyCircularQueue) IsEmpty() bool {
	return this.front == this.rear
}


func (this *MyCircularQueue) IsFull() bool {
	return (this.front + 1) % this.size == this.rear
}





