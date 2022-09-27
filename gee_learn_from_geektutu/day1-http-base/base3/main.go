package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	// 参考gin，使用New()创建gee实例，使用GET()方法添加路由，最后使用Run()启动Web服务。
	// 此处的路由只是静态路由，不支持/hello/:name

	r := gee.New()
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})

	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	// 这里的Run内部是调用了net/http的ListenAndServe
	r.Run(":9999")
}
