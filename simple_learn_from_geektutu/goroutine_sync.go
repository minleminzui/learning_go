package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func download(url string) {
	fmt.Println("start to download", url)
	time.Sleep(time.Second) // 模拟耗时操作
	// wg.Add(1) 为wg添加一个计数,wg.Done(), 减去一个计数
	wg.Done()
}

func main() {
	// sync 方式实现并发多协程

	for i := 0; i < 3; i++ {
		wg.Add(1)
		// go download 启动新的协程并发执行download函数
		go download("a.com/" + string(i+'0'))
	}
	// 等待所有协程执行完毕,也就是每执行一次Add(1),
	// 就需要对应执行一次Done()
	// 类似生产者消费者问题
	wg.Wait()
	fmt.Println("Done!")
}
