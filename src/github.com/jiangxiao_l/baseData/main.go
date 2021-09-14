package main

import "fmt"

func main() {
	// int  int32 和 int64 不是同一个数据类型
	//var a = int64(9)
	//fmt.Println(a)
	//fmt.Printf("%T\n", a)
	//fmt.Printf("values is %v", a)

	// 浮点数
	var b = float32(123.1231)
	if true{
		fmt.Println(1)
	} else {
		fmt.Println(b)
	}

	c := uint32(257)
	fmt.Println("c values is:", c)
}
