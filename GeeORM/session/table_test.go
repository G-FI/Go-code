package session

import (
	"database/sql"
	"geeorm/dialect"
	"github.com/go-playground/assert/v2"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
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
	os.Exit(m.Run())
}

func TestSession_CreateTable(t *testing.T) {
	session.Raw("drop table if exists user").Exec()
	err := session.Model(&User{}).CreateTable()
	if err != nil {
		t.Fatal(err)
	}

	session.Raw("insert into user values(?,?)", "hpy", 20).Exec()
	row := session.Raw("select * from user limit 1").QueryRow()
	var user User
	if err = row.Scan(&user.Name, &user.Age); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, user.Name, "hpy")
	assert.Equal(t, user.Age, 20)
}
func TestSession_HasTable(t *testing.T) {
	session.Raw("Drop table if exists user").Exec()
	session.Model(&User{}).CreateTable()
	assert.Equal(t, session.Model(&User{}).HasTable(), true)
}
func TestSession_DropTable(t *testing.T) {
	session.Raw("Drop table if exists user").Exec()
	session.Model(&User{}).CreateTable()
	assert.Equal(t, session.Model(&User{}).HasTable(), true)
	session.Model(&User{}).DropTable()
	assert.Equal(t, session.Model(&User{}).HasTable(), false)
}
