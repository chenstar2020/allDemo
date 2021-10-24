package clause

import (
	"strings"
)

type Clause struct {
	sql map[Type]string
	sqlVars map[Type][]interface{}
}

type Type int
const (
	INSERT Type = iota
	VALUES
	SELECT
	LIMIT
	WHERE
	ORDERBY
)

// Set 设置子句的命令和参数
func (c *Clause) Set(name Type, vars ...interface{}){
	if c.sql == nil{
		c.sql = make(map[Type]string)
		c.sqlVars = make(map[Type][]interface{})
	}
	sql, vars := generators[name](vars...)
	c.sql[name] = sql
	c.sqlVars[name] = vars
}

// Build 组合各个子句，拼接成完整的sql
func (c *Clause)Build(orders ...Type)(string, []interface{}){
	var sqls []string
	var vars []interface{}
	//将各个子句拼接起来
	for _, order := range orders{
		if sql, ok := c.sql[order]; ok{
			sqls = append(sqls, sql)
			vars = append(vars, c.sqlVars[order]...)
		}
	}
	return strings.Join(sqls, " "), vars
}