package main

import "fmt"

type Student struct {
	name string
	age  int
}

func (stu *Student) hello(person string) string {
	return fmt.Sprintf("hello %s, I am %s", person, stu.name)
}

func main() {
	// 创建一个Student结构体
	stu := &Student{
		name: "Tom",
	}
	// stu := Student{
	// 	name: "Tom",
	// }

	msg := stu.hello("Jack")
	fmt.Println(msg)

	// 还可以用new关键字实例化
	stu2 := new(Student)
	fmt.Println(stu2.hello("Alice"))
}
