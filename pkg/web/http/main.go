package main

import (
	"net/http"
)

func main() {
	web()
}

func web() {
	http.HandleFunc("/abc/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("ok"))
	})

	http.ListenAndServe(":8888", nil) // 适合无服务web. 如果需要中间件类功能需要封装 server
}
