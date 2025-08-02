package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, World!\n")
	}

	postArticleHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Posting Article\n")
	}

	fetchArticleListHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Article List\n")
	}

	fetchArticleHandler := func(w http.ResponseWriter, req *http.Request) {
		articleID := 1
		resStiring := fmt.Sprintf("Article No.%d\n", articleID)
		io.WriteString(w, resStiring)
	}

	postNiceHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Posting Nice")
	}

	postCommentHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Posting Comment\n")
	}

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/article", postArticleHandler)
	http.HandleFunc("/article/list", fetchArticleListHandler)
	http.HandleFunc("/article/1", fetchArticleHandler)
	http.HandleFunc("/article/nice", postNiceHandler)
	http.HandleFunc("/comment", postCommentHandler)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
