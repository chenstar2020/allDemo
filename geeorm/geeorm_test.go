package geeorm

import "testing"

type Student struct{
	Id int `geeorm:"PRIMARY KEY"`
	Name string
	Score uint
}

type Test struct{

}
func TestNewEngine(t *testing.T) {
	_, err := NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/bw_conductor?parseTime=true&loc=Local")
	if err != nil{
		t.Errorf("new engine fail:%v", err)
	}
}


func TestEngine_NewSession(t *testing.T) {
	engine, _ := NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/bw_conductor?parseTime=true&loc=Local")
	defer engine.Close()

	s := engine.NewSession()
	s.Raw("CREATE TABLE User(Name text);").Exec()
	s.Raw("DROP TABLE User").Exec()
}

func TestEngine_CreateTable(t *testing.T) {
	engine, _ := NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/bw_conductor?parseTime=true&loc=Local")
	defer engine.Close()

	s := engine.NewSession()
	s.Model(&Student{})
	s.CreateTable()
}

func TestEngine_DropTable(t *testing.T) {
	engine, _ := NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/bw_conductor?parseTime=true&loc=Local")
	defer engine.Close()

	s := engine.NewSession()
	s.Model(&Student{})
	s.DropTable()
}

func TestEngine_HasTable(t *testing.T) {
	engine, _ := NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/bw_conductor?parseTime=true&loc=Local")
	defer engine.Close()

	s := engine.NewSession()
	s.Model(&Student{})
	t.Log("table exist:", s.HasTable())
}

func TestEngine_Insert(t *testing.T){
	engine, _ := NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/bw_conductor?parseTime=true&loc=Local")
	defer engine.Close()
}