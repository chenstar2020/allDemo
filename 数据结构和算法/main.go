package main

import (
	"./queue"
	"./stack"
	"fmt"
	"strings"
)

/**************************************************
					链表
***************************************************/

type ListNode struct{
	Val int
	Next *ListNode
}

// 两数相加 https://leetcode-cn.com/problems/add-two-numbers/
func addTwoNumbers(l1 *ListNode, l2 *ListNode)*ListNode{
	newList := &ListNode{}
	head := newList
	ptr1 := l1
	ptr2 := l2

	jinwei := 0
	for ;ptr1 != nil || ptr2 != nil;{
		num1 := 0
		if ptr1 != nil{
			num1 = ptr1.Val
			ptr1 = ptr1.Next
		}
		num2 := 0
		if ptr2 != nil{
			num2 = ptr2.Val
			ptr2 = ptr2.Next
		}

		sum := num1 + num2 + jinwei

		newNode := &ListNode{
			Val: sum % 10,
			Next: nil,
		}

		newList.Next = newNode

		jinwei = sum / 10

		newList = newList.Next
	}

	if jinwei > 0{
		newNode := &ListNode{
			Val: jinwei,
			Next: nil,
		}
		newList.Next = newNode
	}

	return head.Next
}

func testAddTwoNumbers(){
	head1 := &ListNode{
		2,
		nil,
	}
	ptr1 := head1
	node := &ListNode{
		4,
		nil,
	}
	ptr1.Next = node
	ptr1 = ptr1.Next

	node = &ListNode{
		3,
		nil,
	}
	ptr1.Next = node
	ptr1 = ptr1.Next

	node = &ListNode{
		9,
		nil,
	}
	ptr1.Next = node
	ptr1 = ptr1.Next

	node = &ListNode{
		8,
		nil,
	}
	ptr1.Next = node
	ptr1 = ptr1.Next

	for ptr := head1; ptr != nil; ptr = ptr.Next{
		fmt.Printf("%d", ptr.Val)
	}
	fmt.Println()

	head2 := &ListNode{
		5,
		nil,
	}
	ptr2 := head2
	node = &ListNode{
		6,
		nil,
	}
	ptr2.Next = node
	ptr2 = ptr2.Next

	node = &ListNode{
		8,
		nil,
	}
	ptr2.Next = node
	ptr2 = ptr2.Next

	for ptr := head2; ptr != nil; ptr = ptr.Next{
		fmt.Printf("%d", ptr.Val)
	}
	fmt.Println()

	rList := addTwoNumbers(head1, head2)

	for ptr := rList; ptr != nil; ptr = ptr.Next{
		fmt.Printf("%d", ptr.Val)
	}
}

//删除链表的倒数第N个结点 https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/
func removeNthFromEnd(head *ListNode, n int) *ListNode{
	//第一种方法 先遍历一遍，得到链表长度
/*	lLen := 0
	ptr := head
	for ;ptr != nil;ptr = ptr.Next{
		lLen++
	}

	if lLen == n {  //删除第一个元素
		head = head.Next
	}else{
		ptr = head
		for count := 0; count < lLen - n - 1; count++{
			ptr = ptr.Next
		}
		if ptr.Next != nil{
			ptr.Next = ptr.Next.Next
		}else{      //只有一个元素的情况
			head = nil
		}
	}

	return head*/

	//第二种方法 双指针，指针一先走n步
	ptr1 := head
	ptr2 := head
	for ;n != 0;n--{
		ptr1 = ptr1.Next
	}


	if ptr1 == nil {  //删除第一个元素
		head = head.Next
	}else{
		for ;ptr1.Next != nil;{
			ptr1 = ptr1.Next
			ptr2 = ptr2.Next
		}

		if ptr2.Next != nil{
			ptr2.Next = ptr2.Next.Next
		}else{        //只有一个元素的情况
			head = nil
		}
	}

	return head
}

func testRemoveNthFromEnd(){
	head1 := &ListNode{
		2,
		nil,
	}
/*	ptr1 := head1
	node := &ListNode{
		4,
		nil,
	}
	ptr1.Next = node
	ptr1 = ptr1.Next

	node = &ListNode{
		6,
		nil,
	}
	ptr1.Next = node
	ptr1 = ptr1.Next

	node = &ListNode{
		8,
		nil,
	}
	ptr1.Next = node
	ptr1 = ptr1.Next*/

	rList := removeNthFromEnd(head1, 1)

	for ptr := rList; ptr != nil; ptr = ptr.Next{
		fmt.Printf("%d", ptr.Val)
	}
}

//两两交换链表中的结点 https://leetcode-cn.com/problems/swap-nodes-in-pairs/
func swapPairs(head *ListNode) *ListNode {
	if head == nil{
		return head
	}

	ptr := &ListNode{      //构造一个头结点
		Val: 0,
		Next: head,
	}

	if ptr.Next.Next != nil{  //保存转换后的头结点
		head = ptr.Next.Next
	}

	for ;ptr.Next != nil && ptr.Next.Next != nil;{
		tmp1 := ptr.Next.Next
		tmp2 := ptr.Next.Next.Next       //保存下次要交换的节点
		ptr.Next.Next.Next = ptr.Next
		ptr.Next.Next = tmp2

		ptr.Next = tmp1
		ptr = ptr.Next.Next
	}

	return head
}

func testSwapPairs(){
	head := newList()
	rList := swapPairs(head)

	for ptr := rList; ptr != nil; ptr = ptr.Next{
		fmt.Printf("%d", ptr.Val)
	}
}

func newList()*ListNode{
	head := &ListNode{
		2,
		nil,
	}
	ptr := head
	node := &ListNode{
		4,
		nil,
	}
	ptr.Next = node
	ptr = ptr.Next

	node = &ListNode{
		6,
		nil,
	}
	ptr.Next = node
	ptr = ptr.Next

	node = &ListNode{
		8,
		nil,
	}
	ptr.Next = node
	ptr = ptr.Next

	return head
}

/**************************************************
					栈
***************************************************/
//接雨水  https://leetcode-cn.com/problems/trapping-rain-water/
func trap(height []int) int {
	//左边最大值
	leftMax := make([]int, 0)
	max := 0
	for _, v := range height{
		leftMax = append(leftMax, max)
		if v > max{
			max = v
		}
	}
	//右边最大值
	max = 0
	tmp := make([]int, 0)
	for k, _ := range height{
		tmp = append(tmp, max)
		if height[len(height) - k - 1] > max{
			max = height[len(height) - k - 1]
		}
	}
	rightMax := make([]int, 0)
	for k, _ := range tmp{
		rightMax = append(rightMax, tmp[len(tmp) - k - 1])
	}

	total := 0
	for k, v := range height{
		total += Max(Min(leftMax[k], rightMax[k]) - v, 0)
	}
	return total
}

func Max(a, b int) int{
	if a > b{
		return a
	}
	return b
}

func Min(a, b int) int{
	if a > b{
		return b
	}
	return a
}

//简化路径 https://leetcode-cn.com/problems/simplify-path/
func simplifyPath(path string) string {
	s := stack.New()
	arr := strings.Split(path, "/")

	for _, v := range arr{
		if len(v) == 0{
			continue
		}

		if v == "."{
			continue
		}else if v == ".."{
			s.Pop()
		}else{
			s.Push(v)
		}
	}

	simplePath := ""
	for ;;{
		tmp := s.Pop()
		if tmp == nil{
			break
		}
		simplePath = fmt.Sprintf("/%s%s", tmp, simplePath)
	}
	if len(simplePath) == 0 {
		simplePath = "/"
	}
	return simplePath
}

//柱状图中最大的矩形 	https://leetcode-cn.com/problems/largest-rectangle-in-histogram/
func largestRectangleArea(heights []int) int {
	hLen := len(heights)
	maxArea := 0
	s := stack.NewStack()
	for k, v := range heights{
		right := k
		left := s.SingleDecreasePush(v)

		for ;right < hLen;{
			if heights[right] >= v{
				right++
			}else{
				break
			}
		}
		fmt.Println(left, right - k, v)
		area := (left + (right - k)) * v
		if area > maxArea{
			maxArea = area
		}
	}
	return maxArea
}

/**************************************************
					队列
***************************************************/
func josephCircle(m, n int)int{
	array := make([]int, m)
	for i := 0; i < m; i++{
		array[i] = i + 1
	}

	ptr := 0      //数组下标
	tmp := 0      //临时计数
	count := 0    //出局人数
	for;;{
		if array[ptr] != -1{
			tmp++
		}

		if tmp == n{      //出局
			count++
			if count >= m {
				return array[ptr]
			}
			array[ptr] = -1
			tmp = 0
		}

		ptr = (ptr + 1) % m
	}
}

func main(){
	//testAddTwoNumbers()
	//testRemoveNthFromEnd()
	//testSwapPairs()
	//fmt.Println(trap([]int{0,1,0,2,1,0,1,3,2,1,2,1}))
	//fmt.Println(simplifyPath("/../"))
	//fmt.Println(largestRectangleArea([]int{2,1,5,6,2,3, 11}))
	//fmt.Println(singleIncreaseStack([]int{10, 7, 4, 12, 10}))
	//fmt.Println(largestRectangleArea([]int{3,6,5,7,4,8,1,0}))
	//fmt.Println(josephCircle(10, 5))

	q := queue.Constructor(5)
	fmt.Println(q.IsFull())
	fmt.Println(q.IsEmpty())
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	q.EnQueue(5)
	q.DeQueue()
	fmt.Println(q.EnQueue(6))
	q.EnQueue(7)
	q.DeQueue()
	q.DeQueue()
	q.EnQueue(8)
	fmt.Println(q.Front())
	fmt.Println(q.Rear())
}
