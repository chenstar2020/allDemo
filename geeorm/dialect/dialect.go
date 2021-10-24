/*
 * 数据库字段和golang数据类型的转换接口
 */

package dialect

import "reflect"

type Dialect interface {
	DataTypeOf(typ reflect.Value) string           			  //将go语言类型转换为数据库的数据类型
	TableExistSQL(tableName string)(string, []interface{})    //返回某个表是否存在的sql语句
}

var dialectsMap = map[string]Dialect{}

func RegisterDialect(name string, dialect Dialect){
	dialectsMap[name] = dialect
}

func GetDialect(name string)(dialect Dialect, ok bool){
	dialect, ok = dialectsMap[name]
	return
}