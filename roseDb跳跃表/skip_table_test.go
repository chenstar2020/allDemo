package skip_table

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestProbabilityTable(t *testing.T) {
	table := probabilityTable(1 / math.E, 18)
	fmt.Println(table)
}


func TestNewSkipList(t *testing.T) {
	skipList := NewSkipList()
	assert.Equal(t, skipList, nil)
}

func TestRandomLevel(t *testing.T){
	skipList := NewSkipList()
	fmt.Println(skipList.randomLevel())
}

func TestSkipList_Put(t *testing.T) {
	skipList := NewSkipList()
	skipList.Put([]byte("key1"), []byte("binary"))
	skipList.Put([]byte("key2"), 19)
	skipList.Put([]byte("key3"), "bbb")
	skipList.Put([]byte("key4"), 30)
	skipList.Put([]byte("key5"), false)

	ele := skipList.Remove([]byte("key5"))
	if ele == nil{
		fmt.Println("delete ok")
	}
}