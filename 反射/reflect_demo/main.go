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

	//获取结构体的值
	type T struct{
		A int
		B string
	}
	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("get struct value %d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
	var x float64 = 3.4
	fmt.Println("atype:", reflect.TypeOf(x))
	fmt.Println("avalue:", reflect.ValueOf(x))
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)  //Kind()方法描述的是基础类型，而不是静态类型
	fmt.Println("value:", v.Float())

	fmt.Println("***********************")
	y := v.Interface().(float64)    //将Value类型恢复其接口类型的值  Interface会把type和value信息打包并填充到接口变量中
	fmt.Println(y, v.Interface())


}
