package main

import  "fmt"

// 批量的声明
var(
	name string
	age  int
)

func main() {
	//变量的声明 var 变量名  类型
	var name string
	name = "hello world"
	age = 11
	fmt.Println("name is ", name)
	fmt.Println("age is ", age)
	// 类型推导
	var a = "text"
	fmt.Println(a)
	// 短变量声明 只能在函数中使用
	b := 1
	fmt.Println(b)
}
