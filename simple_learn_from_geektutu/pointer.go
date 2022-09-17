package main

import "fmt"

func add(num int) {
	num += 1
}

func realAdd(num *int) {
	*num += 1
}

func main() {
	str := "Golang"
	var p *string = &str
	*p = "Hello"
	fmt.Println(str)

	num := 100
	add(num)
	fmt.Println(num)

	realAdd(&num)
	fmt.Println(num)
}
