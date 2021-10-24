package main

import "fmt"

type student struct{
	id int
	name string
}

type Sumifier interface {
	Add(a, b int32) int32
}

type Sumer struct {
	id int32
}

func (math Sumer)Add(a, b int32) int32{
	return a + b
}

func main(){
	adder := Sumer{id:12345}
	m := Sumifier(adder)
	fmt.Println(m.Add(1, 2))
}

func foo(){

}
