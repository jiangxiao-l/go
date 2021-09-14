package main

import "fmt"

const pi = 3.1415926
// const同时声明多个常量时，如果省略了值则表示和上面一行的值相同
const(
	n1 = 100
	n2
	n3
	n4
)
// 在const关键字出现时将被重置为0。const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)
const(
	a1 = iota // 0
	a2  // 1
	_   // 2
	a3  // 3
)

const (
	b1 = iota  // 0
	b2 = 100  // 100
	b3 = iota // 2
	b4       // 3
)

const (
	a, b = iota + 1, iota + 2 //1,2
	c, d                      //2,3
	e, f                      //3,4
)

func main() {

   //fmt.Printf("values n1 %v\n", n1)
   //fmt.Printf("values n2 %v\n", n2)
   //fmt.Printf("values n3 %v\n", n3)
   //fmt.Printf("values n4 %v\n", n4)

	//fmt.Printf("values a1 %v\n", a1)
	//fmt.Printf("values a2 %v\n", a2)
	//fmt.Printf("values a3 %v\n", a3)

	fmt.Printf("values b1 %v\n", b1)
	fmt.Printf("values b2 %v\n", b2)
	fmt.Printf("values b3 %v\n", b3)
	fmt.Printf("values b4 %v\n", b4)
}
