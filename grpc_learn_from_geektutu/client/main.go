package main

import (
	"log"
	"net/rpc"
)

// 在客户端中需要使用Result类型，所以这里拷贝Result定义
type Result struct {
	Num, Ans int
}

type MyType int

func main() {
	// rpc.DialHTTP创建HTTP客户端client，并且创建了与localhost:1234的连接
	client, _ := rpc.DialHTTP("tcp", "localhost:1234")
	var result Result
	// 使用rpc.Call调用远程方法，第一个参数是方法名Cal.Square,后面两个参数与Cal.Square的定义的参数相对应
	// rpc.Call是同步调用，会阻塞当前进程，直到结果返回
	// if err := client.Call("Cal.Square", 12, &result); err != nil {
	// 	log.Fatal("Failed to call Cal.Square.", err)
	// }

	// log.Printf("%d^2 = %d", result.Num, result.Ans)

	// var myVar MyType = 100
	// log.Printf("before minus_one, myVar is %d", myVar)
	// if err := client.Call("MyType.Minus_one", 1025, &myVar); err != nil {
	// 	log.Fatal("Failed to call myType.minus_one.", err)
	// }

	// log.Printf("after minus_one, myVar is %d", myVar)

	// 此处是异步调用
	asyncCall := client.Go("Cal.Square", 12, &result, nil)
	// 第一次异步调用，打印的值是result的零值
	log.Printf("%d^2 = %d", result.Num, result.Ans)

	<-asyncCall.Done
	log.Printf("%d^2 = %d", result.Num, result.Ans)
}
