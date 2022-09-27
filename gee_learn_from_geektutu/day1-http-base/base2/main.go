package main

import (
	"fmt"
	"log"
	"net/http"
)

type Engine struct{}

// 此处算是实现了一个新的路由函数
// 可以向http.ResponseWriter中写入请求的响应
// http.Resquest对象中包含了HTTP请求的所有信息，比如请求地址，Header和Body等信息
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

func main() {
	engine := new(Engine)
	// 此处使用自己定义路由器，可以自己处理日志，异常等
	// 这里的路由器是指实现Handler接口的类型，那么在该接口的ServeHTTP函数中，去处理HTTP请求
	log.Fatal(http.ListenAndServe(":9999", engine))
}
