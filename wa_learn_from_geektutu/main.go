package main

import "syscall/js"

func main() {

	// 使用js.Global().Get("alert")获取全局的alert对象,
	// 通过Invoke方法调用,等价于在js中调用window.alert("Hello World")
	alert := js.Global().Get("alert")
	alert.Invoke("Hello World!")
}
