package main

import (
	"fmt"
	"time"
)

func main() {

	numsChan2 := make(chan int, 10)

	//写
	go func() {
		for i := 0; i < 50; i++ {
			numsChan2 <- i
			fmt.Println("写入数据:", i)
		}
		fmt.Println("数据全部写完毕，准备关闭管道!")
		close(numsChan2)
		//close(numsChan2) //重复关闭，崩溃
		//numsChan2 <- 10 // 关闭之后，再写数据，崩溃
	}()

	//遍历管道时，只返回一个值
	//for range是不知道管道是否已经写完了，所以会一直在这里等待
	//在写入端，将管道关闭，for range遍历关闭的管道时，会退出
	for v := range numsChan2 {
		fmt.Println("读取数据 :", v)
	}

	time.Sleep(3 * time.Second)
	i := <-numsChan2
	fmt.Println("已经关闭之后，继续读取:", i)

	fmt.Println("OVER!")
}
