package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func init() {
	log.SetPrefix("\x1b[92m[GO-MOVIES-SERVER]\x1b[0m")
	log.SetFlags(log.Ldate | log.Ltime)
}

type Movie struct {
	Id       string    `json:"id"`
	Name     string    `json:"name"`
	Director *Director `json:"director"`
}
type Director struct {
	Name string `json:"name"`
}

var movies map[string]Movie

func getMoviesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "Application/json")
	json.NewEncoder(w).Encode(movies)
}
func getMovieHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	requestedMovie := movies[vars["id"]]
	w.Header().Add("content-type", "Application/json")
	json.NewEncoder(w).Encode(requestedMovie)
}
func deleteMovieHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Add("content-type", "text/html")
	if _, err := movies[vars["id"]]; err == false {
		w.WriteHeader(http.StatusBadRequest)
		msg := "movie not found"
		w.Write([]byte(msg))
	} else {
		delete(movies, vars["id"])
		msg := "delete success"
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(msg))
	}
}
func createMovieHandler(w http.ResponseWriter, r *http.Request) {
	//客户端post传入一个json对象
	//先对json对象进行反序列化
	//1. 查看id是否已存在
	//	1.1若存在返回错误信息
	//	1.2不存在时添加新的movie，返回正确信息
	w.Header().Add("context-type", "text/html")
	m := Movie{}
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error in decoding json object"))
		return
	}

	if _, err := movies[m.Id]; err == true {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("this movie id has already exist"))
		return
	}
	movies[m.Id] = m
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("create successful"))
}

func updateMovieHandler(w http.ResponseWriter, r *http.Request) {
	//r.Body是json对象
	//将json对象进行反序列化，获取到对应的电影
	//查找id对应的movie是否存在
	//	1. 不存在就返回错误信息
	//	2. 存在就修改
	w.Header().Add("context-type", "text/html")
	m := Movie{}
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error in decoding request body"))
		return
	}
	//movie 不存在
	if _, err := movies[m.Id]; err == false {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("movie id not exist"))
		return
	}
	movies[m.Id] = m
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("update success"))
}
func main() {
	movies = make(map[string]Movie)
	movies["1"] = Movie{Id: "1", Name: "星际穿越", Director: &Director{Name: "诺兰"}}
	movies["2"] = Movie{Id: "2", Name: "盗梦空间", Director: &Director{Name: "诺兰"}}
	r := mux.NewRouter()
	r.HandleFunc("/movies", getMoviesHandler).Methods("GET")
	r.HandleFunc("/movie/{id}", getMovieHandler).Methods("GET")
	r.HandleFunc("/movie/{id}", deleteMovieHandler).Methods("DELETE")
	r.HandleFunc("/movie/update", updateMovieHandler).Methods("PUT")
	r.HandleFunc("/movie/create", createMovieHandler).Methods("POST")
	//m := Movie{Id: "123",
	//	Name: "xingji",
	//}
	//json.NewEncoder(os.Stdout).Encode(m)
	log.Println("Starting server at port 8000")
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatalf("error: %v\n", err)
	}
}
