package main

import (
	"github.com/g-fi/book-store/pkg/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	//注册路由选项
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	log.Fatal(http.ListenAndServe(":8001", r))
}
