package clause

import (
	"testing"
)

func testSelect(t *testing.T){
	var clause Clause
	clause.Set(SELECT, "User", []string{"*"})
	clause.Set(WHERE, "Name = ?", "Tom")
	clause.Set(ORDERBY, "Age ASC")
	clause.Set(LIMIT, 3)
	sql, vars := clause.Build(SELECT, WHERE, ORDERBY, LIMIT)
	t.Log(sql, vars)
}

func testInsert(t *testing.T){
	var clause Clause
	clause.Set(INSERT, "User", []string{"Id", "Name", "Age"})
	clause.Set(VALUES, []interface{}{4, "Chen", 26})
	sql, vars := clause.Build(INSERT, VALUES)
	t.Log(sql, vars)
}

func TestClause_Build(t *testing.T) {
	t.Run("select", testSelect)
	t.Run("insert", testInsert)
}
