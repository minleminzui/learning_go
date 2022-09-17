package main

import (
	"fmt"

	"example/calc"

	"rsc.io/quote"
)

// go.mod 文件用来记录当前模块的模块名字以及所有以来包版本
func main() {
	fmt.Println(quote.Hello())
	fmt.Println(calc.Add(1, 2))

}
