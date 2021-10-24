package clause

import (
	"testing"
)

func TestSelect(t *testing.T){
	t.Log(_select("User", []string{"Id", "Name", "Age"}))
}

func TestInsert(t *testing.T){
	t.Log(_insert("User", []string{"Id", "Name", "Age"}))
}

func TestLimit(t *testing.T){
	t.Log(_limit(3))
}

func TestWhere(t *testing.T){
	t.Log(_where("Name = ?", "Tom"))
}

func TestValues(t *testing.T){
	t.Log(_values([]interface{}{1, "aa", 4.5}, []interface{}{'a'}))
}