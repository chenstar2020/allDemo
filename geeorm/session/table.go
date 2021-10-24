/*
 * 表相关操作
 */

package session

import (
	"fmt"
	"geeorm/log"
	"geeorm/schema"
	"reflect"
	"strings"
)

// Model 解析结构体为表结构
func (s *Session) Model(value interface{}) *Session{
	if s.refTable == nil || reflect.TypeOf(value) != reflect.TypeOf(s.refTable.Model){
		s.refTable = schema.Parse(value, s.dialect)
	}
	return s
}

// RefTable 得到解析后的结果
func (s *Session) RefTable() *schema.Schema{
	if s.refTable == nil{
		log.ERROR("Model is not set")
		return nil
	}
	return s.refTable
}

func (s *Session) CreateTable() error{
	table := s.RefTable()
	var columns []string
	for _, field := range table.Fields{
		columns = append(columns, fmt.Sprintf("%s %s %s", field.Name, field.Type, field.Tag))
	}

	//拼接列名
	desc := strings.Join(columns, ",")

	_, err := s.Raw(fmt.Sprintf("CREATE TABLE %s (%s)", table.Name, desc)).Exec()
	return err
}

func (s *Session) DropTable()error{
	table := s.RefTable()
	_, err := s.Raw(fmt.Sprintf("DROP TABLE %s", table.Name)).Exec()
	return err
}

func (s *Session) HasTable() bool{
	sql, values := s.dialect.TableExistSQL(s.RefTable().Name)

	rows, err := s.Raw(sql, values...).QueryRows()
	if err != nil{
		return false
	}

	var num int
	for rows.Next(){
		err = rows.Scan(&num)
		if err != nil{
			return false
		}
		if num > 0{
			return true
		}
	}
	return false
}

