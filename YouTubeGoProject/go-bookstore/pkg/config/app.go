package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//type User struct {
//	ID         int64
//	Username   string `gorm:"column:username"`
//	Password   string `gorm:"column:password"`
//	CreateTime int64  `gorm:"column:createtime"`
//}
//
//func (u User) TableName() string {
//	return "users"
//}
var (
	db *gorm.DB
)

func Connect() {
	username := "root"
	passwd := "qwer159357++"
	host := "127.0.0.1"
	port := 3306
	Dbname := "go_bookstore"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		username, passwd, host, port, Dbname)
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败，Error:" + err.Error())
	}
	db = d
	//u := User{
	//	Username:   "hpy",
	//	Password:   "20203021",
	//	CreateTime: time.Now().Unix(),
	//}
	//if err := db.Create(&u).Error; err != nil {
	//	log.Fatal("插入失败," + err.Error())
	//}
	//u = User{}
	//result := db.Where("username=?", "hpy").First(&u)
	//if errors.Is(result.Error, gorm.ErrRecordNotFound) {
	//	log.Fatalf("找不到记录")
	//}
	//fmt.Println("查找成功: ", u)
	//
	////更新
	//fmt.Println("更新record")
	//db.Model(&User{}).Where("username=?", "hpy").Update("password", "abcdefg")
	//u = User{}
	//result = db.Where("username=?", "hpy").First(&u)
	//if errors.Is(result.Error, gorm.ErrRecordNotFound) {
	//	log.Fatalf("找不到记录")
	//}
	//fmt.Println("查找成功: ", u)
	//删除所有record
	//db.Where("username=?", "hpy").Delete(&User{})
}
func GetDB() *gorm.DB {
	return db
}
