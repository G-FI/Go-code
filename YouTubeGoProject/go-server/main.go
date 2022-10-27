package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 page not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Not support method", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "hello")
}
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "formHandler: %v\n", err)
		return
	}
	name := r.FormValue("name")
	addr := r.FormValue("addr")
	fmt.Fprintf(w, "name: %v, addr: %v", name, addr)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("ListenAndServer: %v\n", err)
	}
}
