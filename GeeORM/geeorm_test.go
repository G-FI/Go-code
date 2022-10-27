package geeorm

import (
	"errors"
	"geeorm/session"
	"github.com/go-playground/assert/v2"
	"testing"
)

type User struct {
	Name string "PRIMARY KEY"
	Age  int
}

func TestSession_Model(t *testing.T) {
	engine, err := NewEngine("sqlite3", "gee.db")
	if err != nil {
		t.Fatal(err)
	}

	session := engine.NewSession()
	session.Model(&User{})
	session.DB()
}
func TestSession_CreateTable(t *testing.T) {
	engine, err := NewEngine("sqlite3", "gee.db")
	if err != nil {
		t.Fatal(err)
	}
	session := engine.NewSession()
	session.Model(&User{}).CreateTable()
	assert.Equal(t, session.HasTable(), true)
}

func TestSession_DropTable(t *testing.T) {
	engine, err := NewEngine("sqlite3", "gee.db")
	if err != nil {
		t.Fatal(err)
	}
	session := engine.NewSession()
	session.Model(&User{}).DropTable()
}

func TestEngine_Transaction_Commit(t *testing.T) {
	engine, err := NewEngine("sqlite3", "gee.db")
	defer engine.Close()
	if err != nil {
		t.Fatal(err)
	}
	s := engine.NewSession()
	s.Model(&User{}).DropTable()
	engine.Transaction(func(s *session.Session) (result interface{}, err error) {
		_ = s.Model(&User{}).CreateTable()
		_, err = s.Insert(&User{"Tom", 18})
		return
	})

	var user User
	s.Model(&User{}).First(&user)
	assert.Equal(t, user, User{Name: "Tom", Age: 18})
}
func TestEngine_Transaction_Rollback(t *testing.T) {
	engine, err := NewEngine("sqlite3", "gee.db")
	defer engine.Close()
	if err != nil {
		t.Fatal(err)
	}
	s := engine.NewSession()
	s.Model(&User{}).DropTable()
	s.Model(&User{}).CreateTable()
	engine.Transaction(func(s *session.Session) (result interface{}, err error) {
		s.Model(&User{}).Insert(&User{Name: "Tom", Age: 18})
		//返回error, 将会执行rollback
		return nil, errors.New("Error")
	})
	var user User
	s.Model(&User{}).First(&user)
	assert.Equal(t, user, User{})
}
