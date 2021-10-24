/*
 * 解析表结构
 */

package schema

import (
	"geeorm/dialect"
	"go/ast"
	"reflect"
	"strings"
)

type Field struct {
	Name string         //字段名称
	Type string         //字段类型
	Tag  string         //字段Tag(约束条件)
}

type Schema struct {
	Model interface{}         //数据库对象
	Name string               //表名
	Fields []*Field           //字段（名称、类型、Tag)
	FieldNames []string       //字段名
	fieldMap map[string]*Field
}


func (schema *Schema) GetField(name string) *Field{
	return schema.fieldMap[name]
}

func (schema *Schema) RecordValues(dest interface{}) []interface{} {
	destValue := reflect.Indirect(reflect.ValueOf(dest))
	var fieldValues []interface{}
	for _, field := range schema.Fields{
		fieldValues = append(fieldValues, destValue.FieldByName(field.Name).Interface())
	}
	return fieldValues
}

// Parse 解析表
func Parse(dest interface{}, d dialect.Dialect)*Schema{
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()
	schema := &Schema{
		Model: dest,
		Name:  strings.ToLower(modelType.Name()),
		fieldMap: make(map[string]*Field),
	}

	//解析字段
	for i := 0; i < modelType.NumField(); i++ {
		p := modelType.Field(i)
		if !p.Anonymous && ast.IsExported(p.Name){
			field := &Field{
				Name: p.Name,
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(p.Type))),
			}
			if v, ok := p.Tag.Lookup("geeorm"); ok{
				field.Tag = v
			}
			schema.Fields = append(schema.Fields, field)
			schema.FieldNames = append(schema.FieldNames, p.Name)
			schema.fieldMap[p.Name] = field
		}
	}
	return schema
}
