package testmain

import (
	"fmt"
	"os"
	"testing"
)

func setup(){
	fmt.Println("Before all tests")
}

func teardown(){
	fmt.Println("After all tests")
}

func Test1(t *testing.T){
	fmt.Println("i'm test1")
}

func Test2(t *testing.T){
	fmt.Println("i'm test2")
	t.Fatalf("test2 fail")
}

func TestMain(m *testing.M){
	setup()
	code := m.Run()    //这里会触发所有测试用例的执行
	teardown()
	os.Exit(code)
}
