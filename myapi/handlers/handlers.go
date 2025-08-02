package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, World!\n")
}

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Article\n")
}

func FetchArticleListHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Article List\n")
}

func FetchArticleHandler(w http.ResponseWriter, req *http.Request) {
	articleID := 1
	resStiring := fmt.Sprintf("Article No.%d\n", articleID)
	io.WriteString(w, resStiring)
}

func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Nice")
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Comment\n")
}
