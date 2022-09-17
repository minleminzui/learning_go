package main

import "fmt"

type Gender int8

// go语言没有枚举,这里使用常量模拟枚举
const (
	MALE   Gender = 1
	FEMALE Gender = 2
)

func main() {
	gender := MALE

	// go语言switch,不需要break
	switch gender {
	case FEMALE:
		fmt.Println("female")
	case MALE:
		fmt.Println("male")
	default:
		fmt.Println("unkonwn")
	}

	// 反而 需要继续执行的时候,需要使用fallthrough
	switch gender {
	case FEMALE:
		fmt.Println("female")
		fallthrough
	case MALE:
		fmt.Println("male")
		fallthrough
	default:
		fmt.Println("unknown")
	}
}
