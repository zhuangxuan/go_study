package main

import (
	//_ "day02/05-init/sub" //此时，只会调用sub里面的init函数，编译还不会出错
	_ "awesomeProject2/src/go4/day02/05-init/sub"
	"fmt"
	//"fmt"
)

func main() {
	//res := sub.Sub(10, 5)
	//fmt.Println("sub.Sub(20,10) =", res)
	fmt.Println("Hello world")
}
