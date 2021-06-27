package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Person struct{
	name string
	age int
}

func (p *Person)GetName(a, b, c int)(string, error){
	return p.name, nil
}

func (p *Person)SetAge(age int){
	p.age = age
}

func main(){
	var person Person
	typ := reflect.TypeOf(&person)
	for i := 0; i < typ.NumMethod();i++{
		method := typ.Method(i)
		argv := make([]string, 0, method.Type.NumIn())
		returns := make([]string, 0, method.Type.NumOut())

		for j := 1; j < method.Type.NumIn(); j++ {
			argv = append(argv, method.Type.In(j).Name())
		}
		fmt.Println("aaaaaaaa", method.Type.In(0))
		for j := 0; j < method.Type.NumOut(); j++ {
			returns = append(returns, method.Type.Out(j).Name())
		}
		fmt.Printf("func (w *%s) %s(%s) %s\n",
			typ.Elem().Name(), method.Name, strings.Join(argv, ","), strings.Join(returns, ","))
	}

	fmt.Println(reflect.TypeOf((*error)(nil)).Elem())
}