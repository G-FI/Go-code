package clause

import (
	"github.com/go-playground/assert/v2"
	"reflect"
	"testing"
)

func TestClause_Build(t *testing.T) {
	c := Clause{}
	c.Set(SELECT, "user", []string{"*"})
	c.Set(WHERE, "name=?", "hpy")
	c.Set(ORDERBY, "age DES")
	c.Set(LIMIT, 10)
	sql, vars := c.Build(SELECT, WHERE, ORDERBY, LIMIT)
	assert.Equal(t, sql, "SELECT * FROM user WHERE name=? ORDER BY age DES LIMIT ?")
	assert.Equal(t, reflect.DeepEqual(vars, []interface{}{"hpy", 10}), true)

	c2 := Clause{}
	c2.Set(INSERT, "user", []string{"name", "age"})
	c2.Set(VALUES, []interface{}{"hpy", 20}, []interface{}{"hd", 19})
	sql, vars = c2.Build(INSERT, VALUES)
	assert.Equal(t, sql, "INSERT INTO user(name, age) VALUES (?, ?), (?, ?)")
	assert.Equal(t, reflect.DeepEqual(vars, []interface{}{"hpy", 20, "hd", 19}), true)

	c3 := Clause{}
	c3.Set(SELECT, "user", []string{"*"})
	sql, vars = c3.Build(SELECT, WHERE, ORDERBY, LIMIT)
	t.Log(sql, vars)
}
