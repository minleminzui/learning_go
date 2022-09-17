package main

import (
	"fmt"
	"reflect"
)

func main() {
	arr := [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(arr); i++ {
		arr[i] += 100
	}

	fmt.Println(arr)

	// make 创建的是 切片
	slice1 := make([]float32, 0)
	slice2 := make([]float32, 3, 5)

	fmt.Println("len(slice1): ", len(slice1), "cap(slice2): ", cap(slice2))

	slice2 = append(slice2, 1, 2, 3, 4, 5)

	fmt.Println(len(slice2), cap(slice2))

	// 子切片
	sub1 := slice2[3:]
	sub2 := slice2[:3]
	sub3 := slice2[1:4]

	// append中 sub2... 是解构写法
	combined := append(sub1, sub2...)
	fmt.Println(reflect.TypeOf(combined).Kind(), combined)
	fmt.Println(reflect.TypeOf(sub3).Kind())

}
