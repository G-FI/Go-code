package dialect

import "reflect"

//variable 保存数据库名与实例的映射
var dialectsMap = map[string]Dialect{}

type Dialect interface {
	//返回传入字段在数据相应数据库的类型
	DataTypeOf(reflect.Value) string
	//返回某个表是否存在的dialect(sql, vars)
	TableExist(string) (string, []interface{})
}

//注册数据库
func RegisterDatabase(name string, dialect Dialect) {
	dialectsMap[name] = dialect
}

func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectsMap[name]
	return
}
