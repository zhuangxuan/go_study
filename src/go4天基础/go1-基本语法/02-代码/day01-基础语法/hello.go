// 每个go程序，都必须有一个 package main
// 每个go程序，都是.go结尾的
// go程序中没有.h，没有.o, 只有.go
// 一个package，包名，相当于命名空间
// std::cout
package main

import "fmt"

//主函数，所有的函数必须使用func 开头
//一个函数的返回值，不会放在func前，而是放在参数后面
//函数左花括号必须与函数名同行，不能写到下一行

func main() {

	//name := "World"
	//ptr := &name

	//fmt.Println("hello", name)
	//fmt.Println("address of name:", ptr)

	name2Ptr := new(string)
	*name2Ptr = "World"
	fmt.Println("address of name2:", name2Ptr)
	fmt.Println("value of name2:", *name2Ptr)

	res := testPtr2()
	fmt.Println("res:", res, "value of name2:", *res)

	_ = testPtr2()
}

func testPtr2() *string {
	city := "上海"
	ptr := &city
	return ptr
}
