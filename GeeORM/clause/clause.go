package clause

import "strings"

type Type int

const (
	INSERT Type = iota
	VALUES
	SELECT
	LIMIT
	WHERE
	ORDERBY
	UPDATE
	DELETE
	COUNT
)

//Clause 一个完整Sql查询语句，
//一个子句对应一个Type，对应一个sql 语句string表示，以及一个[]interface{}参数
type Clause struct {
	sql     map[Type]string
	sqlVars map[Type][]interface{}
}

//设置一条sql语句的各子句以及对应参数
func (c *Clause) Set(name Type, vars ...interface{}) {
	if c.sql == nil {
		c.sql = make(map[Type]string)
		c.sqlVars = make(map[Type][]interface{})
	}
	c.sql[name], c.sqlVars[name] = generators[name](vars...)
}

//根据传入子句顺序构造一条sql语句及总的参数
func (c *Clause) Build(types ...Type) (string, []interface{}) {
	var sql []string       //各个独立子句
	var vars []interface{} //各个独立子句参数集合
	for _, tp := range types {
		if v, ok := c.sql[tp]; ok {
			sql = append(sql, v)
			vars = append(vars, c.sqlVars[tp]...)
		}
	}
	return strings.Join(sql, " "), vars
}
