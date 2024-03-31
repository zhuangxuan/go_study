package main

import (
	"fmt"
	"time"
)

func main() {
	//sync.RWMutex{}
	//当涉及到多go程时，c语言使用互斥量，上锁来保持资源同步，避免资源竞争问题
	//go语言也支持这种方式，但是go语言更好的解决方案是使用管道、通道 channel
	//使用通道不需要我们去进行加解锁
	//A 往通道里面写数据  B从管道里面读数据，go自动帮我们做好了数据同步

	//创建管道：  创建一个装数字的管道 ==> channel
	//strChan := make(chan string) //装字符串的管道

	//make(map[int]string, 10)
	//装数字的管道，使用管道的时候一定要make， 同map一样，否则是nil
	//此时是无缓冲的管道
	//numChan := make(chan int)

	//有缓冲的管道
	numChan := make(chan int, 10)

	//创建两个go程，父亲写数据，儿子读数据
	go func() {
		for i := 0; i < 50; i++ {
			data := <-numChan
			fmt.Println("子go程1 读取数据  ===》 data:", data)
		}
	}()

	go func() {
		for i := 0; i < 20; i++ {
			//runtime.Gosched()
			//向管道中写入数据
			numChan <- i
			fmt.Println("子go程2 写入数据:", i)
			//time.Sleep(1 * time.Second)
		}

	}()

	for i := 20; i < 50; i++ {
		//向管道中写入数据
		numChan <- i
		fmt.Println("======> 这是主go程, 写入数据:", i)
		//time.Sleep(1 * time.Second)
	}

	time.Sleep(5 * time.Second)
}
