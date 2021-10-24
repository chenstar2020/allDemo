package list

import (
	"fmt"
)

// Element 结点
type Element struct {
	Data interface{}
	next *Element //后继指针
	pre  *Element //前置指针
}

type List struct {
	head   *Element
	length int
}

func (l *List) Init() *List {
	l.head = nil
	l.length = 0
	return l
}

func (l *List) Len() int {
	return l.length
}

func (l *List) Front() *Element {
	return l.head
}

func (l *List) Back() *Element {
	ptr := l.head
	for;ptr.next != nil;{
		ptr = ptr.next
	}
	return ptr
}

func (l *List) PushFront(v interface{}) *Element {
	l.length++

	if l.head.Data == nil{
		l.head.Data = v
		return l.head
	}

	newElem := &Element{
		Data: v,
		next: nil,
		pre:  nil,
	}
	head := l.head
	newElem.next = head
	head.pre = newElem
	l.head = newElem

	return newElem
}

func (l *List) PushFrontList(other *List) {
	if other.Len() == 0{
		return
	}

	head := other.head

	ptr := l.head
	for ;ptr != nil;{
		other.PushBack(ptr.Data)
		ptr = ptr.next
	}

	l.head = head
}

func (l *List) PushBack(v interface{}) *Element {
	l.length++

	if l.head.Data == nil{
		l.head.Data = v
		return l.head
	}

	newElem := &Element{
		Data: v,
		next: nil,
		pre:  nil,
	}
	ptr := l.head
	for ;ptr.next != nil; {
		ptr = ptr.next
	}

	ptr.next = newElem
	newElem.pre = ptr

	return newElem
}

func (l *List) PushBackList(other *List) {
	if other.Len() == 0{
		return
	}

	ptr := other.head
	for ;ptr != nil;{
		l.PushBack(ptr.Data)
		ptr = ptr.next
	}
}

func (l *List) findEle(e *Element)bool{
	ptr := l.head
	for;ptr != e && ptr != nil;{
		ptr = ptr.next
	}
	return ptr != nil
}

func (l *List) InsertBefore(v interface{}, mark *Element) *Element {
	if !l.findEle(mark){
		return nil
	}

	newElem := &Element{
		Data: v,
		next: nil,
		pre:  nil,
	}

	ptr := mark

	if ptr.pre == nil{   		//head node
		if ptr.Data == nil{
			ptr.Data = v
		}else{
			ptr.pre = newElem
			newElem.next = ptr
			l.head = newElem
		}
	}else{
		ptr.pre.next = newElem
		newElem.pre = ptr.pre

		newElem.next = ptr
		ptr.pre = newElem
	}

	l.length++

	return newElem
}

func (l *List) InsertAfter(v interface{}, mark *Element) *Element {
	if !l.findEle(mark){
		return nil
	}

	newElem := &Element{
		Data: v,
		next: nil,
		pre:  nil,
	}

	ptr := mark

	if ptr.next == nil{   		//tail node
		if ptr.Data == nil{
			ptr.Data = v
		}else{
			ptr.next = newElem
			newElem.pre = ptr
		}
	}else{
		ptr.next.pre = newElem
		newElem.next = ptr.next

		newElem.pre = ptr
		ptr.next = newElem
	}

	l.length++

	return newElem
}

func (l *List) MoveToFront(e *Element) {
	if !l.findEle(e){
		return
	}

	value := l.Remove(e)

	l.PushFront(value)
}

func (l *List) MoveToBack(e *Element) {
	if !l.findEle(e){
		return
	}

	value := l.Remove(e)

	l.PushBack(value)
}

func (l *List) MoveBefore(e, mark *Element) {
	if !l.findEle(e){
		return
	}
	if !l.findEle(mark){
		return
	}
	if e == mark{
		return
	}

	value := l.Remove(e)

	l.InsertBefore(value, mark)
}

func (l *List) MoveAfter(e, mark *Element) {
	if !l.findEle(e){
		return
	}
	if !l.findEle(mark){
		return
	}
	if e == mark{
		return
	}

	value := l.Remove(e)

	l.InsertAfter(value, mark)
}

func (l *List) Remove(e *Element) interface{} {
	head := l.head
	for ;head != nil;{
		if head == e{
			if head.pre == nil{ 		//head node
				head.next.pre = head.pre
				l.head = head.next
			}else if head.next == nil{  //tail node
				head.pre.next = head.next
			}else{
				head.pre.next = head.next
				head.next.pre = head.pre
			}
			break
		}
		head = head.next
	}

	l.length--

	return head.Data
}

func (l *List) Print(){
	fmt.Println("**************list start**************")
	head := l.head
	for ;head != nil;{
		fmt.Println(head.Data)
		head = head.next
	}
	fmt.Println("**************list end****************")
}

func New() *List{
	return &List{
		head:   &Element{nil, nil, nil},
		length: 0,
	}
}

type ListMethod interface {
	Init() *List            			//清空链表
	Len() 	int							//返回长度
	Front() *Element					//返回链表第一个元素
	Back()  *Element           //返回链表最后一个元素
	PushFront(v interface{}) *Element //前插
	//创建链表other拷贝，并将拷贝的最后一个位置连接到链表l的第一个位置
	PushFrontList(other *List)
	PushBack(v interface{}) *Element  //尾插
	PushBackList(other *List)
	InsertBefore(v interface{}, mark *Element) *Element //mark之前插入结点
	InsertAfter(v interface{}, mark *Element) *Element
	MoveToFront(e *Element)  //将元素e移动到链表的第一个位置
	MoveToBack(e *Element)   //将元素e移动到链表的最后一个位置
	MoveBefore(e, mark *Element)
	MoveAfter(e, mark *Element)
	Remove(e *Element) interface{}
}


var _ ListMethod = (*List)(nil)