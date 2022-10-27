package routes

import (
	"github.com/g-fi/book-store/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(r *mux.Router) {
	r.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
	r.HandleFunc("/book/{id}", controllers.GetBookByID).Methods("GET")
	r.HandleFunc("/book/", controllers.AddBook).Methods("POST")
	r.HandleFunc("/book/", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/book/", controllers.DeleteBook).Methods("DELETE")
}
