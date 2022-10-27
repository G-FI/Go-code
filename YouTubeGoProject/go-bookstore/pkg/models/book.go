package models

import (
	"fmt"
	"github.com/g-fi/book-store/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	//使用Model中的createTime, updateTime,此时删除只是标记为DELETED,而不是真正从数据库删除
	//gorm.Model
	ID   int64  `gorm:"column:id;primerKey" json:"id"`
	Name string `json:"name"`
}

//建立数据库连接
func init() {
	config.Connect()
	db = config.GetDB()
}

func NewBook(id int64, name string) *Book {
	return &Book{ID: id, Name: name}
}

//TODO: 添加错误处理，如果添加不成功
func AddBook(b *Book) {
	db.Create(b)
}
func DeleteBook(b *Book) error {
	result := db.Delete(b)
	if result.RowsAffected == 0 {
		return fmt.Errorf("No such record")
	}
	return result.Error
}

func GetBooks() ([]Book, error) {
	books := make([]Book, 0)
	result := db.Find(&books)
	return books, result.Error
}

func GetBookById(id int64) (*Book, error) {
	book := &Book{}
	result := db.Where("id = ?", id).First(book)
	return book, result.Error
}
func UpdateBook(old *Book, new *Book) error {
	result := db.Model(old).Updates(new)
	return result.Error
}
