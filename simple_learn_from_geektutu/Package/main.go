package main

import "fmt"

// 如果类型/接口/方法/函数/字段的首字母大写，则是 Public 的，对其他 package 可见，如果首字母小写，则是 Private 的，对其他 package 不可见。
// 一个文件夹,可以视作一个package
func main() {
	fmt.Println(add(3, 5))
}
