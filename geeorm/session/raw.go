/*
 * 执行sql命令
 */

package session

import (
	"database/sql"
	"geeorm/clause"
	"geeorm/dialect"
	"geeorm/log"
	"geeorm/schema"
	"strings"
)

type Session struct {
	db *sql.DB
	dialect dialect.Dialect   //golang数据类型和db数据转换
	refTable *schema.Schema   //存储表结构的解析结果
	clause clause.Clause
	sql strings.Builder
	sqlVars []interface{}
}

func New(db *sql.DB, dialect dialect.Dialect) *Session{
	return &Session{db: db, dialect: dialect}
}

func (s *Session) Clear(){
	s.sql.Reset()
	s.sqlVars = nil
	s.clause = clause.Clause{}
}

func (s *Session) DB() *sql.DB{
	return s.db
}

// Raw 接收sql和参数
func (s *Session) Raw(sql string, values ...interface{})*Session{
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	s.sqlVars = append(s.sqlVars, values...)
	return s
}

func (s *Session) Exec() (result sql.Result, err error){
	defer s.Clear()
	log.INFO(s.sql.String(), s.sqlVars)
	if result, err = s.DB().Exec(s.sql.String(), s.sqlVars...); err != nil{
		log.ERROR(err)
	}
	return
}

func (s *Session)QueryRow() *sql.Row{
	defer s.Clear()
	log.INFO(s.sql.String(), s.sqlVars)
	return s.DB().QueryRow(s.sql.String(), s.sqlVars...)
}

func (s *Session)QueryRows() (*sql.Rows, error) {
	defer s.Clear()
	log.INFO(s.sql.String(), s.sqlVars)
	return s.DB().Query(s.sql.String(), s.sqlVars...)
}
