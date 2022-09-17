package main

import "fmt"

type Person interface {
	getName() string
}

type Student struct {
	name string
	age  int
}

func (stu *Student) getName() string {
	return stu.name
}

func (stu *Student) getAge() int {
	return stu.age
}

type Worker struct {
	name   string
	gender string
}

func (w *Worker) getName() string {
	return w.name
}

func main() {
	// 这里将nil转换为*Student类型,再转换为Person接口,如果转换失败,
	// 说明Student并没有实现Person接口所有方法
	// 这里用来在编译期检测 接口实现的完整性
	var _ Person = (*Student)(nil)
	var _ Person = (*Worker)(nil)
	// go语言不需要显示的声明接口,只需要实现接口对应的方法即可
	var p Person = &Student{
		name: "Tom",
		age:  18,
	}

	fmt.Println(p.getName())

	// 这里将Person接口转换为实例类型
	stu := p.(*Student) // 接口转换实例
	fmt.Println(stu.getAge())

	// 这个map的value是一个空接口
	// 那么表明value可以是任何类型
	m := make(map[string]interface{})
	m["name"] = "Tom"
	m["age"] = 18
	m["scores"] = [3]int{98, 99, 85}
	fmt.Println(m)
}
