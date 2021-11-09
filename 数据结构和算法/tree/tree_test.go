package tree

import "testing"

func InitTree() *TreeNode{
	return CreateBinaryTree([]interface{}{"A", "B", "D", "#", "#", "E", "#", "#", "C", "F", "H", "#", "#", "#", "G", "#", "#"})
}

func InitTreeInt() *TreeNode{
	return CreateBinaryTree([]interface{}{5, 4, 11, 7, "#", "#", 2, "#", "#", "#", 8, 13, "#", "#", 4, "#", 1, "#", "#"})
}

func TestCreateBinaryTree(t *testing.T) {
	InitTree()
}

func TestTreeNode_PreOrder(t *testing.T) {
	tree := InitTree()
	tree.PreOrder()
}

func TestTreeNode_PreOrder1(t *testing.T) {
	tree := InitTree()
	tree.PreOrder1()
}

func TestTreeNode_MidOrder(t *testing.T) {
	tree := InitTree()
	tree.MidOrder()
}

func TestTreeNode_MidOrder1(t *testing.T) {
	tree := InitTree()
	tree.MidOrder1()
}

func TestTreeNode_PostOrder(t *testing.T) {
	tree := InitTree()
	tree.PostOrder()
}

func TestTreeNode_PostOrder1(t *testing.T) {
	tree := InitTree()
	tree.PostOrder1()
}

func TestTreeNode_LayerOrder(t *testing.T) {
	tree := InitTree()
	tree.LayerOrder()
}

func TestTreeNode_MaxDepth(t *testing.T) {
	tree := InitTree()
	t.Log(tree.MaxDepth())
}

func TestTreeNode_MinDepth(t *testing.T) {
	tree := InitTree()
	t.Log(tree.MinDepth())
}

func TestTreeNode_Width(t *testing.T) {
	tree := InitTree()
	t.Log(tree.Width())
}

func TestTreeNode_PreOrderSum(t *testing.T) {
	tree := InitTreeInt()
	sum := 0
	result := 18
	tree.PreOrderSum(sum, result)
}


func TestTreeNode_Insert(t *testing.T) {
	tree := NewTree()
	tree.Data = 1
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(5)
	tree.Insert(4)
	tree.Insert(7)
	tree.PreOrder()
}