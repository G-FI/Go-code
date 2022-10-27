package cmd_test

import (
	"database/sql"
	"geeorm"
	"github.com/go-playground/assert/v2"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"reflect"
	"testing"
)

func TestSqlite(t *testing.T) {
	db, _ := sql.Open("sqlite3", "gee.db")
	defer func() { db.Close() }()
	db.Exec("Drop TABLE  If exists user")
	db.Exec("create table user(name text, age integer );")

	//tx, _ := db.Begin()
	res, err := db.Exec("insert into user(name, age) values(?,?), (?,?)", "tom", 1, "jerry", 2)
	if err != nil {
		t.Log(res.RowsAffected())
	}
	//tx.Rollback()

	row := db.QueryRow("select * from user where name=?", "jerry")
	type user struct {
		name string
		age  int
	}
	var u user
	row.Scan(&u.name, &u.age)
	t.Log(u)
}
func TestType(t *testing.T) {
	var s *string
	typeS := reflect.TypeOf(s)
	t.Logf("%T", typeS.Kind())
}
func TestFirstDay(t *testing.T) {
	e, _ := geeorm.NewEngine("sqlite3", "gee.db")
	session := e.NewSession()
	session.Raw("Drop table if exists user").Exec()
	session.Raw("create table user(name text, age integer)").Exec()
	session.Raw("insert into user(name, age) values($1, $2)", "hpy", 12).Exec()
	row := session.Raw("select age from user limit 1").QueryRow()
	//var name string
	//row.Scan(&name)
	//assert.Equal(t, name, "hpy")
	var age int
	row.Scan(&age)
	assert.Equal(t, age, 12)
}

type User struct {
	gorm.Model
	Name string
	Age  int
}

func TestGorm(t *testing.T) {

	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	u := User{Name: "hpy", Age: 20}
	u1 := User{Name: "hd", Age: 19}
	result := db.Create([]*User{&u, &u1})
	t.Log(u.ID)
	t.Log(result.RowsAffected)
	var u2 User
	db.First(&u2)
	t.Log(u2)
}
