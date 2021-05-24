package main

import "fmt"

type Getter interface {
	Get(key string) ([]byte, error)
}

type GetterFunc func(key string)([]byte, error)

func (f GetterFunc)Get(key string)([]byte, error){
	return f(key)
}

func main() {
	var f1 = GetterFunc(func(key string) ([]byte, error) {
		fmt.Println("Getter", key)
		return nil, nil
	})

	f1("key1")
}

