package main

import "fmt"

func main() {
	age := 18
	if age < 18 {
		fmt.Println("Kid")
	} else {
		fmt.Println("Adult")
	}

	// 简写
	if age := 18; age < 18 {
		fmt.Println("Kid")
	} else {
		fmt.Println("Adult")
	}
}
