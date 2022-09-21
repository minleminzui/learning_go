package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human  // 匿名字段
	school string
	loan   float32
}

type Employee struct {
	Human   // 匿名字段
	company string
	money   float32
}

// Human 实现 SayHi 方法
func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// Human 实现 Sing 方法
func (h Human) Sing(lyrics string) {
	fmt.Println("La la la la...", lyrics)
}

// Employee 重载 Human 的 SayHi 方法
func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone)
}

// Interface Men 被 Human, Student 和 Employee 实现
// 因为这三个类型都实现了这两个方法
type Men interface {
	SayHi()
	Sing(lyrics string)
}

type Element interface{}
type List []Element

type Person struct {
	name string
	age  int
}

// 定义了 String 方法，实现了 fmt.Stringer
func (p Person) String() string {
	return "(name: " + p.name + " - age: " + strconv.Itoa(p.age) + " years)"
}

func main() {
	// mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	// paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	// sam := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	// tom := Employee{Human{"Tom", 37, "222-444-XXX"}, "Things Ltd.", 5000}

	// var i Men
	// i = mike
	// fmt.Println("This is Mike, a Student:")
	// i.SayHi()
	// i.Sing("Novermber rain")

	// i = tom
	// fmt.Println("This is tom, an Employee")
	// i.SayHi()
	// i.Sing("Born to be wild")

	// fmt.Println("Let's use a silce of Men and see what happens")
	// x := make([]Men, 3)
	// x[0], x[1], x[2] = paul, sam, mike

	// for _, value := range x {
	// 	value.SayHi()
	// }

	// 使用Comma-ok 断言来判断一个interface变量中实际存储的是什么类型的
	list := make(List, 3)
	list[0] = 1
	list[1] = "Hello"
	list[2] = Person{"Dennis", 70}

	for index, element := range list {
		if value, ok := element.(int); ok {
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		} else if value, ok := element.(string); ok {
			fmt.Printf("list[%d] is an string and its value is %s\n", index, value)
		} else if value, ok := element.(Person); ok {
			fmt.Printf("list[%d] is an Person and its value is %s\n", index, value)
		} else {
			fmt.Printf("list[%d] is a different type\n", index)
		}
	}

	for index, element := range list {
		switch value := element.(type) {
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		case string:
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		case Person:
			fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
		default:
			fmt.Printf("list[%d] is of a different type", index)
		}
	}

	// 反射,就是检查程序运行时的状态
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())

}
