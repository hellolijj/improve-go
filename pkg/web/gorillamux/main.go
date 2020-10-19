package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

var (
	HomeHandler = func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("home"))
	}
	ProductsHandler = func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("product" + request.URL.Path))
	}
	ArticlesHandler = func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("article"))
	}
)

func main() {
	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/products/{key}", ProductsHandler)
	r.HandleFunc("/articles", ArticlesHandler)
	http.ListenAndServe(":8888", r)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
