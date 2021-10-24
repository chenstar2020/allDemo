package schema

import (
	"geeorm/dialect"
	"testing"
)

type User struct{
	Name string `geeorm:"PRIMARY KEY"`
	Age int `json:"age"`
}

var dial, _ = dialect.GetDialect("sqlite3")

func TestParse(t *testing.T) {
	schema := Parse(&User{}, dial)
	t.Log(schema.Name)
	for name, field := range schema.fieldMap{
		t.Log(name, field.Name, field.Type, field.Tag)
	}
}
