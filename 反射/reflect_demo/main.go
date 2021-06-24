package main

//参考文档  http://c.biancheng.net/view/4407.html

import (
	"fmt"
	"reflect"
)

type student struct{}
func main(){
	//var a int
	var stu student
	typeOfA := reflect.TypeOf(stu)
	fmt.Println(typeOfA.Name(), typeOfA.Kind())  //student  struct


	type cat struct{
		Name string
		Type int `json:"type" id:"100"`
	}

	//对指针获取反射对象
	ins := &cat{}
	typeOfCat := reflect.TypeOf(ins)
	fmt.Printf("name:'%v' kind:'%v'\n", typeOfCat.Name(), typeOfCat.Kind())   // ''  ptr
	typeOfCatElem := typeOfCat.Elem()  //等效于对指针变量*操作
	fmt.Printf("element name:'%v' kind:'%v'\n", typeOfCatElem.Name(), typeOfCatElem.Kind()) //cat struct

	//获取结构体的字段
	ins2 := cat{
		Name:"mimi",
		Type: 1,
	}
	typeOfCat2 := reflect.TypeOf(ins2)
	for i := 0; i < typeOfCat2.NumField(); i++ {  //获取结构体成员字段数量
		filedType := typeOfCat2.Field(i)       //返回索引对应结构体字段的信息
		fmt.Printf("name:%v tag: '%v'\n", filedType.Name, filedType.Tag)  //打印字段名称和Tag
	}
	if catType, ok := typeOfCat2.FieldByName("Type");ok{
		fmt.Println(catType.Tag.Get("json"), catType.Tag.Get("id"))
	}
}
