package clause

import (
	"github.com/go-playground/assert/v2"
	"reflect"
	"testing"
)

type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

func TestSelect(t *testing.T) {
	sql, _ := generators[SELECT]("user", []string{"name", "age"})
	assert.Equal(t, sql, "SELECT name, age FROM user")
}

func TestInsert(t *testing.T) {
	sql, _ := generators[INSERT]("user", []string{"name", "age"})
	assert.Equal(t, sql, "INSERT INTO user(name, age)")
}
func TestWhere(t *testing.T) {
	sql, vars := generators[WHERE]("name=?", "hpy")
	assert.Equal(t, sql, "WHERE name=?")
	assert.Equal(t, reflect.DeepEqual(vars, []interface{}{"hpy"}), true)
}

func TestOrderBy(t *testing.T) {
	sql, _ := generators[ORDERBY]("age DSC")
	assert.Equal(t, sql, "ORDER BY age DSC")
}
func TestValues(t *testing.T) {
	sql, vars := generators[VALUES]([]interface{}{"hpy", 12})
	assert.Equal(t, sql, "VALUES (?, ?)")
	assert.Equal(t, reflect.DeepEqual(vars, []interface{}{[]interface{}{"hpy", 12}}), true)

	sql, vars = generators[VALUES]([]interface{}{"hpy", 12}, []interface{}{"hd", 11})
	assert.Equal(t, sql, "VALUES (?, ?), (?, ?)")

	assert.Equal(t, reflect.DeepEqual(vars, []interface{}{[]interface{}{"hpy", 12}, []interface{}{"hd", 11}}), true)
}
func TestLimit(t *testing.T) {
	sql, vars := generators[LIMIT](10)
	assert.Equal(t, sql, "LIMIT ?")
	assert.Equal(t, reflect.DeepEqual(vars, []interface{}{10}), true)
}

func TestDelete(t *testing.T) {
	sql, vars := generators[DELETE]("user")
	assert.Equal(t, sql, "DELETE FROM user")
	assert.Equal(t, reflect.DeepEqual(vars, []interface{}{}), true)
}
func TestUpdate(t *testing.T) {
	sql, vars := generators[UPDATE]("user", map[string]interface{}{"name": "hpy", "age": 1})
	assert.Equal(t, sql, "UPDATE user SET name=?, age=?")
	assert.Equal(t, reflect.DeepEqual(vars, []interface{}{"hpy", 1}), true)
}
