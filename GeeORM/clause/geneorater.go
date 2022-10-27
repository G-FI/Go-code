package clause

import (
	"fmt"
	"strings"
)

//生成对应关键词对应的sql子句，及参数

type generator func(values ...interface{}) (string, []interface{})

var generators map[Type]generator

func init() {
	generators = make(map[Type]generator)
	generators[SELECT] = _select
	generators[INSERT] = _insert
	generators[VALUES] = _values
	generators[WHERE] = _where
	generators[ORDERBY] = _orderby
	generators[LIMIT] = _limit
	generators[DELETE] = _delete
	generators[UPDATE] = _update
	generators[COUNT] = _count
}

//写生成语句的辅助函数，返回标准库接受的sql语句，以及参数

//DELETE FROM $tableName
func _delete(values ...interface{}) (string, []interface{}) {
	return fmt.Sprintf("DELETE FROM %s", values[0].(string)), []interface{}{}
}

//UPDATE $tableName SET ($columnName = ?)*
//values两个元素，tableName, map[]interface{}
func _update(values ...interface{}) (string, []interface{}) {
	tableName := values[0].(string)
	kv := values[1].(map[string]interface{})
	var column []string
	var vars []interface{}
	for k, v := range kv {
		column = append(column, k+"=?")
		vars = append(vars, v)
	}
	return fmt.Sprintf("UPDATE %s SET %s", tableName, strings.Join(column, ", ")), vars
}

//1: tableName
func _count(values ...interface{}) (string, []interface{}) {
	return _select(values[0], []string{"COUNT(*)"})
}

//接受两个参数，一个tableName, 一个fields: []string
func _insert(values ...interface{}) (string, []interface{}) {
	//INSERT INTO $tableName(fields...)
	tableName := values[0]
	fields := strings.Join(values[1].([]string), ", ")
	return fmt.Sprintf("INSERT INTO %s(%s)", tableName, fields), []interface{}{}
}

//参数[tableName, [fields]]
func _select(values ...interface{}) (string, []interface{}) {
	//SELECT $fields FROM $tableName
	tableName := values[0]
	//间values[1]转化为[]string
	fields := strings.Join(values[1].([]string), ", ")
	return fmt.Sprintf("SELECT %s FROM %s", fields, tableName), []interface{}{}
}
func getBindVars(num int) string {
	var builder []string
	for i := 0; i < num; i++ {
		builder = append(builder, "?")
	}
	return fmt.Sprintf("(%s)", strings.Join(builder, ", "))
}
func _values(values ...interface{}) (string, []interface{}) {
	//VALUES (?, ?)*  vars = values
	//传入的values切片中每一个元素都是一个interface切片，表示一组value, *传出是要将传入的二维切片flatten*
	sql := strings.Builder{}
	var vars []interface{}
	sql.WriteString("VALUES ")
	for i, val := range values {
		//确定参数个数,确定写入几个？
		vars = append(vars, val.([]interface{})...)
		num := len(val.([]interface{}))
		sql.WriteString(getBindVars(num))
		//判断是不是最后一参数，如果不是则追加空格继续循环，是则不用
		if i < len(values)-1 {
			sql.WriteString(", ")
		}
	}
	return sql.String(), vars
}

func _where(values ...interface{}) (string, []interface{}) {
	predict := "WHERE " + values[0].(string)
	vars := values[1:]
	return predict, vars
}
func _orderby(values ...interface{}) (string, []interface{}) {
	//ORDER BY $desc
	return fmt.Sprintf("ORDER BY %s", values[0]), []interface{}{}
}
func _limit(values ...interface{}) (string, []interface{}) {
	//LIMIT $num
	return "LIMIT ?", values
}
