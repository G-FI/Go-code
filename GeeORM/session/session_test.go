package session

import (
	"database/sql"
	"geeorm/dialect"
	"github.com/go-playground/assert/v2"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"testing"
)

var session *Session

func init() {
	db, err := sql.Open("sqlite3", "gee.db")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	dial, ok := dialect.GetDialect("sqlite3")
	if !ok {
		log.Fatal("dailect 获取失败")
	}
	session = New(db, dial)
}

type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

func (u *User) BeforeQuery(s *Session) {
	log.Println("执行BeforeQuery Hook")
}
func (u *User) AfterQuery(s *Session) {
	log.Println("执行After Hook")
	u.Age += 1000
}

func TestSession_Exec(t *testing.T) {
	session.Raw("Drop table if exists user").Exec()
	session.Raw("create table user(name text, age integer)").Exec()
	session.Raw("insert into user(name, age) values(?,?)", "abc", 0).Exec()
	session.Raw("insert into user(name, age) values(?,?), (?,?)", "hpy", 20, "hd", 19).Exec()
}

func TestSession_Query(t *testing.T) {
	session.Raw("Drop table if exists user").Exec()
	session.Raw("create table user(name text, age integer)").Exec()
	session.Raw("insert into user(name, age) values(?,?)", "abc", 0).Exec()
	session.Raw("insert into user(name, age) values(?,?), (?,?)", "hpy", 20, "hd", 19).Exec()

	row := session.Raw("select * from user limit 1").QueryRow()
	var user User
	row.Scan(&user.Name, &user.Age)
	assert.Equal(t, user.Name, "abc")
	assert.Equal(t, user.Age, 0)

	rows, err := session.Raw("select * from user").QueryRows()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		if err = rows.Scan(&user.Name, &user.Age); err != nil {
			log.Fatal(err)
		}
		t.Log(user)
	}
}
