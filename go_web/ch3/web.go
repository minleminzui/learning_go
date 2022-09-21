package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// Request 是用户请求的信息,用来解析用户的请求信息,包括 post, get, cookie, url等信息
// Response 服务器需要反馈给客户端的信息
// Conn 用户的每次请求连接
// Handlder 处理请求和生成返回信息的处理逻辑

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // 解析参数, 默认不解析参数
	fmt.Println(r.Form) // 这些信息被输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("val: ", strings.Join(v, ""))
	}
	fmt.Fprint(w, "Hello yitong!") // 这个写到 w 的是输出到客户端的
}

func main() {
	// 原生支持高并发,用户的每一次请求都是在一个新的goroutinue服务
	http.HandleFunc("/", sayhelloName)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
