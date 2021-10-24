/*
 * mysql和golang语言类型的转换
 */

package dialect

import (
	"fmt"
	"reflect"
	"time"
)

type mysql struct{}

func (m *mysql) DataTypeOf(typ reflect.Value) string {
	switch typ.Kind(){
	case reflect.Bool:
		return "bool"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32:
		return "integer"
	case reflect.Int64, reflect.Uint64:
		return "bigint"
	case reflect.String:
		return "text"
	case reflect.Array, reflect.Slice:
		return "blob"
	case reflect.Struct:
		if _, ok := typ.Interface().(time.Time); ok{
			return "datetime"
		}
	}
	panic(fmt.Sprintf("invalid sql type %s (%s)", typ.Type().Name(), typ.Kind()))
}

func (m *mysql) TableExistSQL(tableName string) (string, []interface{}) {
	args := []interface{}{tableName}
	return "SELECT COUNT(*) AS NUM FROM information_schema.TABLES where TABLE_NAME = ?", args
}

var _ Dialect = (*mysql)(nil)

func init(){
	RegisterDialect("mysql", &mysql{})
}