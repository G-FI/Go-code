package controllers

import (
	"github.com/g-fi/book-store/pkg/models"
	"github.com/g-fi/book-store/pkg/utils"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	/*使用models.GetBooks方法返回books slice
	调用utils.Marshal 将books进行序列化，
	发送给客户端*/
	w.Header().Add("content-type", "application/json")
	books, _ := models.GetBooks()
	if data, err := utils.Marshal(books); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{}"))
		utils.Logf("GetBooks: %v", err)
	} else {

		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 0, 64)
	book, err := models.GetBookById(id)
	if err == gorm.ErrRecordNotFound {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{}"))
	} else {
		if data, err := utils.Marshal(book); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{}"))
			utils.Logf("GetBookByID: %v", err)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(data)
		}
	}

}
func AddBook(w http.ResponseWriter, r *http.Request) {
	b := &models.Book{}
	if err := utils.ParseBody(r, b); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("AddBook: parse body error, " + err.Error()))
		return
	}
	models.AddBook(b)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("add success"))
}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "text/html")
	book := &models.Book{}
	if err := utils.ParseBody(r, book); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad request"))
		utils.Logf("DeleteBook: %v", err)
	} else {
		if err = models.DeleteBook(book); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("delete success"))
		}
	}
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	/*前端发送json数组，其中有两个元素，第一个是old，第二个是new
	获取旧的book信息，和新的book信息，用两个结构体保存
	将两个book传给model，model中进行更新，并返回是否成功*/
	books := make([]models.Book, 2)
	if err := utils.ParseBody(r, &books); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad request"))
		utils.Logf("DeleteBook: %v", err)
	} else {
		w.WriteHeader(http.StatusOK)
		_ = models.UpdateBook(&books[0], &books[1])
		w.Write([]byte("update success"))
	}
}
