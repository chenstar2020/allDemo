package skip_table

import (
	"bytes"
	"math"
	"math/rand"
	"time"
)

const (
	maxLevel int = 18   //跳跃表最大层数
	probability float64 = 1 / math.E
)


type handleEle func(e *Element) bool

type (
	//定义跳跃表用到的数据结构

	//跳跃表的节点  这个节点是element切片
	//切片的大小代表着一个节点有多少个索引
 	//头节点有18级
	Node struct{
		next []*Element
	}

	//存储数据的元素
	//每个元素有key和value 以及下一个节点的指针
	Element struct {
		Node        //指向下一个节点
		key []byte
		value interface{}
	}


	SkipList struct {
		Node   //Node指向跳跃表的开头
		maxLevel int
		Len		 int
		randSource rand.Source
		probability float64
		probTable   []float64
		prevNodesCache []*Node
	}
)

//创建跳跃表
func NewSkipList() *SkipList{
	return &SkipList{
		Node: Node{
			next: make([]*Element, maxLevel),
		},
		prevNodesCache: make([]*Node, maxLevel),
		maxLevel: maxLevel,
		randSource: rand.New(rand.NewSource(time.Now().UnixNano())),
		probability: probability,
		probTable: probabilityTable(probability, maxLevel),
	}
}

//获得元素的key
func (e *Element) Key() []byte{
	return e.key
}

//获得元素的value
func (e *Element) Value() interface{}{
	return e.value
}

//设置元素的value
func (e *Element)SetValue(value interface{}){
	e.value = value
}

//当前节点的下一个节点(第一层)
func (e *Element) Next() *Element{
	return e.next[0]
}

//跳跃表头节点
func (t *SkipList) Front() *Element{
	return t.next[0]
}

func (t *SkipList) Get(key []byte) *Element{
	var prev = &t.Node
	var next *Element

	//从最高层开始查找
	for i := t.maxLevel - 1; i >= 0; i--{
		next = prev.next[i]

		for next != nil && bytes.Compare(key, next.key) > 0 {  //目标key比下一个key大，向右移动
			prev = &next.Node
			next = next.next[i]
		}
	}

	if next != nil && bytes.Compare(key, next.key) <= 0{
		return next
	}

	return nil
}


func (t *SkipList) Exist(key []byte) bool{
	return t.Get(key) != nil
}

func (t *SkipList) backNodes(key []byte) []*Node{
	var prev = &t.Node
	var next *Element

	//记录查找节点过程中所经过的节点
	prevs := t.prevNodesCache

	for i := t.maxLevel - 1; i >=0; i--{
		next = prev.next[i]

		for next != nil && bytes.Compare(key, next.key) > 0 {  //目标key比下一个key大，向右移动
			prev = &next.Node
			next = next.next[i]
		}

		prevs[i] = prev
	}
	return prevs
}

func (t *SkipList) Remove(key []byte) *Element{
	//记录查找过程中每层经过的最后一个节点
	prev := t.backNodes(key)

	//第一层的下一个节点不为空，且key相等 说明找到了要删除的节点
	if element := prev[0].next[0]; element != nil && bytes.Compare(element.key, key) <= 0{
		for k, v := range element.next{
			//把要删除节点的下一个节点赋值给要删除节点的前一个节点
			prev[k].next[k] = v
		}

		//跳跃表长度减一
		t.Len--
		return element
	}
	return nil
}

func (t *SkipList) Put(key []byte, value interface{}) *Element{
	var element *Element

	prev := t.backNodes(key)

	//key已经存在的情况
	if element = prev[0].next[0]; element != nil && bytes.Compare(key, element.key) <= 0{
		element.value = value
		return element
	}

	element = &Element{
		key: key,
		value: value,
		Node:  Node{
			//创建一个随机层的切片
			next: make([]*Element, t.randomLevel()),
		},
	}

	for i:= range element.next{
		element.next[i] = prev[i].next[i]
		prev[i].next[i] = element
	}

	t.Len++
	return element
}

//随机生成索引层数
func (t *SkipList) randomLevel()(level int){
	//生成0-1之间的随机数
	r := float64(t.randSource.Int63()) / (1 << 63)

	//第一层一定生成
	level = 1

	for level < t.maxLevel && r < t.probTable[level]{
		level++
	}

	return
}

func probabilityTable(probability float64, maxLevel int)(table []float64){
	for i := 1; i <= maxLevel; i++ {
		prob := math.Pow(probability, float64(i - 1))
		table = append(table, prob)
	}
	return table
}