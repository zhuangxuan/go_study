package main

import (
	"fmt"
	"time"
)

//return  ===> 返回当前函数
//exit ===> 退出当前进程
//GOEXIT ===> 提前退出当前go程

func main() {
	go func() {
		go func() {
			for {
				fmt.Println("子go程内部循环")
				time.Sleep(1 * time.Second)
			}
			fmt.Println("子go程结束!") //这句会打印吗？ 会1：  不打印2
			fmt.Println("go 2222222222 ")

		}()
		//time.Sleep(2 * time.Second)
		//time.Sleep(100 * time.Second)
		fmt.Println("go 111111111111111")
	}()

	fmt.Println("这是主go程!")
	time.Sleep(10000 * time.Second)
	fmt.Println("OVER!")
}
