package tree

import (
	"../queue"
	"../stack"
	"fmt"
	"math"
)

type TreeNode struct{
	Data 	interface{}
	Left 	*TreeNode
	Right 	*TreeNode
}

func NewTree() *TreeNode{
	return &TreeNode{
		Data: nil,
		Left: nil,
		Right: nil,
	}
}

//先序遍历创建二叉树
/*
原始数据：[A, B, D, #, #, E, #, #, C, F, #, #, G, #, #]
创建树：
	  A
	/   \
   B     C
  / \   / \
 D   E F   G
##   ####  ##
*/

func CreateBinaryTree(data []interface{}) *TreeNode{
	n := 0
	return CreateBinaryTreeByNum(data, &n)
}

func CreateBinaryTreeByNum(data []interface{}, n *int) *TreeNode{
	if *n >= len(data){
		return nil
	}
	if data[*n] == "#"{
		*n++
		return nil
	}

	newNode := NewTree()
	newNode.Data = data[*n]
	*n++

	newNode.Left = CreateBinaryTreeByNum(data, n)
	newNode.Right = CreateBinaryTreeByNum(data, n)

	return newNode
}

/*******************二叉搜索树************************/
//插入
func(tree *TreeNode) Insert(v interface{}){
	if v.(int) < tree.Data.(int){
		if tree.Left != nil{
			tree.Left.Insert(v)
		}else{
			tree.Left = &TreeNode{v, nil, nil}
		}
	}else{
		if tree.Right != nil{
			tree.Right.Insert(v)
		}else{
			tree.Right = &TreeNode{v, nil, nil}
		}
	}
}

func (tree *TreeNode)Search(v interface{}) bool{
	if v.(int) == tree.Data.(int){
		return true
	}else if v.(int) < tree.Data.(int){
		if tree.Left != nil{
			return tree.Left.Search(v)
		}else{
			return false
		}
	}else{
		if tree.Right != nil{
			return tree.Right.Search(v)
		}else{
			return false
		}
	}
	return false
}

func (tree *TreeNode)FindMax()int{
	if tree == nil{
		return -1
	}
	if tree.Right == nil{
		return tree.Data.(int)
	}
	return tree.Right.FindMax()
}


func (tree *TreeNode)Delete(){

}

/*******************二叉搜索树************************/

// PreOrder 先序遍历
func (tree *TreeNode)PreOrder(){
	if tree == nil{
		return
	}
	fmt.Println(tree.Data)
	tree.Left.PreOrder()
	tree.Right.PreOrder()
}

// PreOrder1 先序遍历非递归实现
func (tree *TreeNode)PreOrder1(){
	if tree == nil{
		return
	}
	s := stack.New()
	s.Push(tree)
	for;!s.IsEmpty();{
		top := s.Pop()
		if top.(*TreeNode) != nil{
			fmt.Println(top.(*TreeNode).Data)
			if top.(*TreeNode).Right != nil{
				s.Push(top.(*TreeNode).Right)
			}
			if top.(*TreeNode).Left != nil{
				s.Push(top.(*TreeNode).Left)
			}
		}
	}
}

// MidOrder 中序遍历
func (tree *TreeNode)MidOrder(){
	if tree == nil{
		return
	}
	tree.Left.MidOrder()
	fmt.Println(tree.Data)
	tree.Right.MidOrder()
}

// MidOrder1 中序遍历非递归实现
func (tree *TreeNode)MidOrder1(){
	if tree == nil{
		return
	}
	s := stack.New()
	s.Push(tree)
	node := tree
	for;!s.IsEmpty();{
		for;node != nil && node.Left != nil;{
			s.Push(node.Left)
			node = node.Left
		}
		node = s.Pop().(*TreeNode)
		fmt.Println(node.Data)
		node = node.Right
		if node != nil{
			s.Push(node)
		}
	}

}

// PostOrder 后序遍历
func (tree *TreeNode)PostOrder(){
	if tree == nil{
		return
	}
	tree.Left.PostOrder()
	tree.Right.PostOrder()
	fmt.Println(tree.Data)
}

// PostOrder1 后序遍历非递归实现
func (tree *TreeNode)PostOrder1(){
	if tree == nil{
		return
	}
	s := stack.New()
	s.Push(tree)
	pre := &TreeNode{}     //指向当前节点的前一个节点
	for;!s.IsEmpty();{
		node := s.Peek().(*TreeNode)
		if (node.Left == nil && node.Right == nil) ||
			(pre != nil &&(pre == node.Left || pre == node.Right)){   //说明左右子节点都已经访问过了
			fmt.Println(node.Data)
			s.Pop()
			pre = node
		}else{
			if node.Right != nil{
				s.Push(node.Right)
			}
			if node.Left != nil{
				s.Push(node.Left)
			}
		}
	}
}

// LayerOrder 层次遍历
func (tree *TreeNode)LayerOrder(){
	if tree == nil{
		return
	}

	q := queue.NewQueue()
	q.Add(tree)
	for;;{
		node := q.Remove()
		if node != nil{
			fmt.Println(node.(*TreeNode).Data)
			if node.(*TreeNode).Left != nil{
				q.Add(node.(*TreeNode).Left)
			}
			if node.(*TreeNode).Right != nil{
				q.Add(node.(*TreeNode).Right)
			}
		}
		if q.IsEmpty(){
			break
		}
	}
}

// MaxDepth 最大深度
func (tree *TreeNode)MaxDepth()int{
	if tree == nil{
		return 0
	}
	leftDepth := tree.Left.MaxDepth()
	rightDepth:= tree.Right.MaxDepth()
	return int(math.Max(float64(leftDepth), float64(rightDepth))) + 1
}

// MinDepth 最小深度
func (tree *TreeNode)MinDepth()int{
	if tree == nil{
		return 0
	}
	leftDepth := tree.Left.MinDepth()
	rightDepth:= tree.Right.MinDepth()
	return int(math.Min(float64(leftDepth), float64(rightDepth))) + 1
}

// Width 宽度
func (tree *TreeNode)Width()int{
	if tree == nil{
		return 0
	}

	maxWidth := 0
	q := queue.NewQueue()
	q.Add(tree)

	for ;!q.IsEmpty();{
		width := q.Len()
		if width > maxWidth{
			maxWidth = width
		}

		for i := 0; i < width;i++{
			node := q.Remove()
			if node.(*TreeNode).Left != nil{
				q.Add(node.(*TreeNode).Left)
			}
			if node.(*TreeNode).Right != nil{
				q.Add(node.(*TreeNode).Right)
			}
		}
	}

	return maxWidth
}

func (tree *TreeNode)PreOrderSum(sum int, result int){
	if tree == nil{
		return
	}
	//fmt.Println(tree.Data)
	sum += tree.Data.(int)

	if tree.Left == nil && tree.Right == nil && sum == result{
		fmt.Println("find it", sum, result)
	}
	tree.Left.PreOrderSum(sum, result)
	tree.Right.PreOrderSum(sum, result)
}