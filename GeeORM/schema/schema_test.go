package schema

import (
	"geeorm/dialect"
	"testing"
)

type User struct {
	Name string
	Age  int
	Id   int `geeorm:"PRIMARY KEY"`
}

var dial dialect.Dialect

func init() {
	d, ok := dialect.GetDialect("sqlite3")
	dial = d
	if !ok {
		panic("dialect unsupported")
	}
}

func TestParse(t *testing.T) {
	schema := Parse(&User{}, dial)
	t.Logf("%v", schema)
}

func TestSchema_RecordValues(t *testing.T) {
	schema := Parse(&User{}, dial)
	u := &User{Name: "hpy", Age: 12, Id: 1}
	fields := schema.RecordValues(u)
	t.Log(fields)
}