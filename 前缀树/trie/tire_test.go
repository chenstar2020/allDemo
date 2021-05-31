package trie

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T){
	root := &node{
		'#',
		make([]*node, 0),
	}

	root.insert("abcdefg", 0)
	root.insert("abcgn", 0)
	root.insert("abcgk", 0)
	root.insert("higu", 0)

	fmt.Println(root.find("abtg", 0))
}
