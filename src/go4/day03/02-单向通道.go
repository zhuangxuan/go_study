package main

import (
	"fmt"
	"time"
)

func main() {
	//- 单向读通道:
	//var numChanReadOnly <-chan int
	//- 单向写通道:
	//var numChanWriteOnly chan<- int

	//numRead <- "hello" //只读通道，不能写
	//<-numWrite //只写通道，不能读

	//单向通道的应用场景
	//1. 限制通道的方向，防止误操作
	//2. 作为函数参数，限制函数的操作
	//3. 作为函数的返回值，限制函数的操作

	//管道的应用场景
	//1. 用于多个go程之间的数据传递
	//2. 用于go程和主go程之间的数据传递
	//3. 用于多个go程之间的数据共享

	//生产者消费者模型
	//C: 数组+锁   thread1 : 写， thread2：读
	//Go: goroutine + channel

	//1. 在主函数中创建一个双向通道 numChan
	numChan1 := make(chan int, 5)

	//2. 将numChan 传递给producer， 负责生产
	go producer(numChan1) //双向通道可以赋值给同类型的单向通道, 单向不能转双向

	//2. 将numChan 传递给consumer 负责消费
	go consumer(numChan1)

	time.Sleep(2 * time.Second)
	fmt.Println("OVER!")
}

// producer 生产者   ===> 提供一个只写的通道
func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i
		//data := <-out  //写通道不允许有读取操作
		fmt.Println("====> 向管道中写入数据:", i)
	}

	close(out)
}

// consumer 消费者   ===> 提供一个只读的通道
func consumer(in <-chan int) {
	//in <- 10 //读通道不允许有写入操作
	for v := range in {
		fmt.Println("从管道读取数据：", v)
	}

	fmt.Println("consumer end 111111!")
}
