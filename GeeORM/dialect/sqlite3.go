package dialect

import (
	"fmt"
	"reflect"
	"time"
)

//sqlite3 dialect
type sqlite3 struct {
}

//注册sqlite3 方言
func init() {
	RegisterDatabase("sqlite3", &sqlite3{})
}

func (s sqlite3) DataTypeOf(v reflect.Value) string {
	//类型匹配
	switch v.Kind() {
	case reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uintptr:
		return "integer"
	case reflect.Uint64,
		reflect.Int64:
		return "bigint"
	case reflect.Float32,
		reflect.Float64:
		return "real"
	case reflect.Bool:
		return "bool"
	case reflect.String:
		return "text"
	case reflect.Slice,
		reflect.Array:
		return "blob"
	case reflect.Struct:
		//表中支持的golang struct只有time
		if _, ok := v.Interface().(time.Time); ok {
			return "datetime"
		} else {
			panic(fmt.Sprintf("unsupported sql type: %s (%v)", v.Type().Name(), v.Type().Kind()))
		}
	}
	panic(fmt.Sprintf("unsupported sql type: %s (%v)", v.Type().Name(), v.Type().Kind()))
}

func (s *sqlite3) TableExist(tableName string) (sql string, vars []interface{}) {
	vars = []interface{}{tableName}
	sql = "select name from sqlite_master where type='table' and name=?"
	return
}
