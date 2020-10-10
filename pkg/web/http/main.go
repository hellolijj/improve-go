package main

import "net/http"

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("ok"))
	})

	// 类型转换
	http.Handle("/h", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("h ok"))
	}))

	http.ListenAndServe(":8888", nil)  // 适合无服务web. 如果需要中间件类功能需要封装 server
}