package main

import (
	"errors"
	"fmt"
)

func hello(name string) error {
	// 可以通过errors.New返回自定义错误
	if len(name) == 0 {
		return errors.New("error: name is null")
	}
	fmt.Println("Hello,", name)
	return nil
}

func get(index int) (ret int) {
	// 这里的defer开了一个协程,来处理panic
	defer func() {
		// recover()会使程序从panic中恢复
		if r := recover(); r != nil {
			fmt.Println("Some error happened!", r)
			ret = -1
		}
	}()

	arr := [3]int{2, 3, 4}
	return arr[index]
}

func main() {
	// os.Open 函数有两个返回值,一个是*File, 一个是error

	// _, error := os.Open("filename.txt")

	// if error != nil {
	// 	fmt.Println(error)
	// }

	// if err := hello(""); err != nil {
	// 	fmt.Println(err)
	// }
	fmt.Println(get(5))
	fmt.Println("finished")
}
