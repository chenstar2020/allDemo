package trie

type node struct {
	key uint8
	children []*node
}

var root *node

func new(){
	root = &node{
		'#',
		make([]*node, 0),
	}
}

func (n *node) patternChild(key uint8)*node{
	for _, child := range n.children{
		if child.key == key{
			return child
		}
	}
	return nil
}


func (n *node)insert(str string, num int){
	if num == len(str){
		return
	}

	child := n.patternChild(str[num])
	if child == nil {
		child = &node{
			str[num],
			make([]*node, 0),
		}
		n.children = append(n.children, child)
	}

	child.insert(str, num + 1)
}

func (n *node)find(str string, num int)bool{
	if num == len(str){
		return true
	}
	child := n.patternChild(str[num])
	if child == nil {
		return false
	}

	return child.find(str, num + 1)
}