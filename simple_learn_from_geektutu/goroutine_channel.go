package main

import (
	"fmt"
	"time"
)

var ch = make(chan string, 10) // 创建大小为10的缓冲信道

func download(url string) {
	fmt.Println("start to download", url)
	time.Sleep(time.Second)
	// 这里将url发送给管道
	ch <- url
}

func main() {
	for i := 0; i < 3; i++ {
		go download("a.com/" + string(i+'0'))
	}

	for i := 0; i < 3; i++ {
		msg := <-ch // 等待信道返回消息
		fmt.Println("finish", msg)
	}
	fmt.Println("Done!")
}
