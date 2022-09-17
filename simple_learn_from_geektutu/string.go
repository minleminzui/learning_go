package main

import (
	"fmt"
	"reflect"
)

func main() {
	str1 := "Golang"
	str2 := "Go语言"
	// reflect.TypeOf(str2[2]).Kind()用来知道某个变量的类型
	fmt.Println(reflect.TypeOf(str2[2]).Kind())
	fmt.Println(str1[2], string(str1[2]))
	fmt.Printf("%d %c\n", str2[2], str2[2])
	fmt.Println("len(str2):", len(str2))

	// 对于[]rune类型,字符串的每个字符,无论占多少个字节都用int32表示
	runArr := []rune(str2)

	fmt.Println(reflect.TypeOf(runArr[2]).Kind())
	fmt.Println(runArr[2], string(runArr[2]))
	fmt.Println("len(runeArr): ", len(runArr))
}
