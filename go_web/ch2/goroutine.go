package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched() // 让出时间片
		fmt.Println(s)
	}
}

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	// go say("world")
	// say("hello")
	// c := make(chan int, 1) // 修改 2 为 1 就报错，修改 2 为 3 可以正常运行
	// c <- 1
	// c <- 2
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
