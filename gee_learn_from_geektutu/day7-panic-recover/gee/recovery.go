package gee

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
)

func Recovery() HandlerFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				message := fmt.Sprintf("%s", err)
				log.Printf("%s\n\n", trace(message))
				c.Fail(http.StatusInternalServerError, "Internal Server Error")
			}
		}()

		c.Next()
	}
}

// trace 函数用来获取panic时的堆栈信息
func trace(message string) string {
	var pcs [32]uintptr
	// Callers 是用来返回调用栈的程序计数器，
	// 第0个Caller是Callers本身，第1个是上一层trace
	// 第2个是再上一层defer func
	n := runtime.Callers(3, pcs[:]) // skip first 3 caller

	var str strings.Builder
	str.WriteString(message + "\nTraceback:")

	for _, pc := range pcs[:n] {
		// 通过FuncForPC获取对应函数
		fn := runtime.FuncForPC(pc)
		// 通过FileLine(pc)获取到调用该函数的文件名和行号
		file, line := fn.FileLine(pc)
		str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}

	return str.String()
}
