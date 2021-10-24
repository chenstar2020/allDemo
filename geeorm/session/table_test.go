package session

import "testing"

type Student struct{
	Id int `geeorm:"PRIMARY KEY"`
	Name string
	Score uint
}

func TestSession_CreateTable(t *testing.T) {

}
