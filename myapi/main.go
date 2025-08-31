package main

import (
	"log"
	"net/http"

	"github.com/Shu1F/go_API_intermediate/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// u, _ := url.Parse("http://localhost:8080?page=1&page=2&a=1&")
	// queryMap := u.Query()

	// fmt.Println(queryMap["page"])
	// fmt.Println(queryMap["a"])
	// fmt.Println(queryMap["b"])

	r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", handlers.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", handlers.FetchArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", handlers.FetchArticleHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", handlers.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", handlers.PostCommentHandler).Methods(http.MethodPost)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
