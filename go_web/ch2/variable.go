package main

import (
	"errors"
	"fmt"
	"reflect"
)

// iota 用来声明enum,默认开始的值是0
const (
	a       = iota // a=0
	b       = "B"
	c       = iota             //c=2
	d, e, f = iota, iota, iota //d=3,e=3,f=3
	g       = iota             //g = 4
)

func main() {
	fmt.Println(reflect.TypeOf(a).Kind())
	fmt.Println(a, b, c, d, e, f, g)

	// 1. 字符串
	s := "hello"
	c := []byte(s)
	c[0] = 'c'
	s2 := string(c)
	fmt.Printf("%s\n", s2)

	m := "world"

	a := s + m
	fmt.Printf("%s\n", a)

	s3 := "c" + s[1:]
	fmt.Printf("%s\n", s3)

	// 对于``包裹地字符串,将原样输出
	s4 := `hello
			world`
	fmt.Printf("%s\n", s4)

	// 2. 错误类型
	err := errors.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Println(err)
	}

	// 2. array
	var arr [10]int
	arr[0] = 42
	arr[1] = 13
	fmt.Printf("The first element is %d\n", arr[0])
	fmt.Printf("The last element is %d\n", arr[9])

	arr1 := [...]int{3, 4, 6}
	fmt.Println(arr1)

	// 3. slice 动态数组, 本质是胖指针
	Array_a := [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	Slice_a := Array_a[2:5]
	fmt.Println(cap(Slice_a))

	// 4. map
	// 初始化一个字典
	rating := map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "C++": 2}
	// map 有两个返回值,第二个返回值,如果不存在 key,那么 ok 为 false, 否则为 true
	csharpRating, ok := rating["C#"]

	if ok {
		fmt.Println("C# is in the map and its rating is ", csharpRating)
	} else {
		fmt.Println("We have no rating assocaited with C# in the map")
	}

	delete(rating, "C")

	// map 是一种引用类型
	// rune 的实际类型是int32
	// byte 的实际类型是unit8
}
