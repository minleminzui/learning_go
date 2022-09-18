package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/rpc"
)

type Result struct {
	Num, Ans int
}

type Cal int

// func (cal *Cal) Square(num int) *Result {
// 	return &Result{
// 		Num: num,
// 		Ans: num * num,
// 	}
// }

func (cal *Cal) Square(num int, result *Result) error {
	result.Num = num
	result.Ans = num * num
	return nil
}

type MyType int

func (cal *MyType) Minus_one(num int, result *MyType) error {
	// 由此可以看出client的初值没有用
	temp := *result
	*result = temp + MyType(num) - 1
	return nil
}

// 对于不满足rpc规范的函数，不会被注册
func (cal *Cal) Plus_one(n int) int {
	return n + 1
}

// func (t *T) MethodName(argType T1, replyType *T2) error
// 1.方法类型 (T) 是导出的(首字母大写)
// 2.方法名(MethodName)是导出的
// 3.方法有2个参数(argType T1, replyType *T2), 均为导出/内置类型
// 4.方法的第2个参数一个指针(replyType *T2)
// 5.方法的返回值类型是error
// net/rpc对参数个数的限制比较严格,仅能有2个,第一个参数是调用提供的请求参数,
// 第二个参数是返回给调用者的响应参数,也就是说,服务端计算结果写在第二个参数
// 如果调用过程中发生错误,会返回error给调用者
func main() {
	// 使用rpc.Register 发布Cal中满足RPC注册条件的方法(Cal.Square)
	// rpc.Register(new(Cal))
	// rpc.Register(new(MyType))
	// // 使用rpc.HandleHTTP注册处理RPC消息的HTTP Handler
	// rpc.HandleHTTP()

	// log.Printf("Serving RPC server on port %d", 1234)
	// // 使用http.ListenAndServe监听1234端口，等待RPC请求
	// if err := http.ListenAndServe(":1234", nil); err != nil {
	// 	log.Fatal("Error serving", err)
	// }

	// 这里使用TLS层加密
	// 服务端使用生成的server.crt和server.key文件启动TLS的端口监听
	rpc.Register(new(Cal))
	cert, _ := tls.LoadX509KeyPair("server.crt", "server.key")
	config := &tls.Config{
		Certificates: []tls.Certificates{cert},
	}
	listener, _ := tls.Listen("tcp", ":1234", config)
	log.Printf("Serving RPC server on port %d", 1234)

	for {
		conn, _ := listener.Accept()
		defer conn.Close()
		go rpc.ServerConn(conn)
	}
}
