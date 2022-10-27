package session

import (
	"github.com/go-playground/assert/v2"
	"reflect"
	"testing"
)

func TestSession_Insert(t *testing.T) {
	u := &User{Name: "hpy", Age: 12}
	session.Model(&User{}).DropTable()
	session.Model(&User{}).CreateTable()
	rwoAffected, _ := session.Insert(u)
	assert.Equal(t, rwoAffected, int64(1))
}

func TestSession_Find(t *testing.T) {
	//添加三个user{hpy, 20}{hd, 19}{jy, 18}
	users := make([]User, 0, 3)
	session.Model(&User{}).Find(&users)
	t.Log(users)
}
func TestSession_Count(t *testing.T) {
	session.Model(&User{}).DropTable()
	session.Model(&User{}).CreateTable()
	session.Insert(&User{Name: "hpy", Age: 12}, &User{Name: "hd", Age: 11})
	count, _ := session.Model(&User{}).Count()
	assert.Equal(t, count, int64(2))
}

func TestSession_Update(t *testing.T) {
	session.Model(&User{}).DropTable()
	session.Model(&User{}).CreateTable()
	session.Model(&User{}).Insert(&User{Name: "hpy", Age: 12}, &User{Name: "hd", Age: 11})
	session.Model(&User{}).Update(map[string]interface{}{"name": "wj", "age": 45})
	var user []User
	session.Model(&User{}).Find(&user)
	t.Log(user)
}
func TestSession_Delete(t *testing.T) {
	session.Model(&User{}).DropTable()
	session.Model(&User{}).CreateTable()
	session.Model(&User{}).Insert(&User{Name: "hpy", Age: 12}, &User{Name: "hd", Age: 11})
	rows, _ := session.Delete()
	assert.Equal(t, rows, int64(2))
}

func TestChainOperation(t *testing.T) {
	users := []*User{
		&User{Name: "hpy", Age: 20},
		&User{Name: "hd", Age: 19},
		&User{Name: "jy", Age: 18},
		&User{Name: "zhb", Age: 17},
		&User{Name: "hhm", Age: 16},
	}
	session.Model(&User{}).DropTable()
	session.Model(&User{}).CreateTable()
	t.Log(session.Model(&User{}).Insert(users[0], users[1], users[2], users[3], users[4]))

	var userFound []User
	t.Log(session.Find(&userFound))
	t.Log(userFound)

	userFound = make([]User, 0)
	t.Log(session.Where("Age > ?", 18).Find(&userFound))
	t.Log(userFound)

	userFound = make([]User, 0)
	t.Log(session.Where("Age > ?", 17).OrderBy("Age ASC").Limit(2).Find(&userFound))
	t.Log(userFound)

	//userFound = make([]User, 0)
	t.Log(session.Where("Name like ?", "h%").Delete())
	t.Log(userFound)
}

func TestSession_First(t *testing.T) {
	session.Model(&User{}).DropTable()
	session.Model(&User{}).CreateTable()
	session.Model(&User{}).Insert(&User{Name: "hpy", Age: 12}, &User{Name: "hd", Age: 11})
	var u User
	session.First(&u)
	assert.Equal(t, reflect.DeepEqual(u, User{Name: "hpy", Age: 12}), true)
}
