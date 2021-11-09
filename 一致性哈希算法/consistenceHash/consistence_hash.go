package consistenceHash

import (
	"container/heap"
	"hash/crc32"
	"sort"
	"strconv"
)

type hashFunc func(data []byte)uint32

type ConsistenceHash struct {
	replicas int    //复制次数
	keys   []int    //保存所有的key-》hash
	hashMap map[int]string
	hashFunc hashFunc
}

func New(replicas int, hashFunc2 hashFunc)*ConsistenceHash{
	if hashFunc2 == nil {
		hashFunc2 = crc32.ChecksumIEEE
	}
	return &ConsistenceHash{
		replicas: replicas,
		hashMap: map[int]string{},
		hashFunc: hashFunc2,
	}
}

func (d *ConsistenceHash)Set(keys ...string){
	for _, key := range keys {
		for i := 1; i <= d.replicas; i++ {
			hash := int(d.hashFunc([]byte(strconv.Itoa(i) + key)))
			d.keys = append(d.keys, hash)
			d.hashMap[hash] = key
		}
	}

	sort.Ints(d.keys)
}


func (d *ConsistenceHash)Del(key string){
	for i := 0; i < d.replicas; i++ {
		hash := int(d.hashFunc([]byte(strconv.Itoa(i) + key)))

		for j := 0; j < len(d.keys); j++ {
			if d.keys[j] == hash{
				d.keys = append(d.keys[:j], d.keys[j+1:]...)
				break
			}
		}

		delete(d.hashMap, hash)
	}
}


func (d *ConsistenceHash)Get(key string)string{
	if len(d.keys) == 0 {
		return ""
	}

	hash := int(d.hashFunc([]byte(key)))
	for i := 0; i < len(d.keys); i++ {
		if d.keys[i] >= hash {
			if match, ok := d.hashMap[d.keys[i]]; ok {
				return match
			}
		}
	}

	return d.hashMap[d.keys[0]]
}

func (d *ConsistenceHash)Test(){
	heap.Init()
}