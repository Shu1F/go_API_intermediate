package main

import (
	"log"
	"net/http"

	"github.com/Shu1F/go_API_intermediate/handlers"
)

func main() {

	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/article", handlers.PostArticleHandler)
	http.HandleFunc("/article/list", handlers.FetchArticleListHandler)
	http.HandleFunc("/article/1", handlers.FetchArticleHandler)
	http.HandleFunc("/article/nice", handlers.PostNiceHandler)
	http.HandleFunc("/comment", handlers.PostCommentHandler)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
