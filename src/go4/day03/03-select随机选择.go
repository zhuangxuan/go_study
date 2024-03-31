package main

import "fmt"

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)

	//将三个管道都关闭，从关闭的管道读取数据，不会崩溃，会返回零值
	//close(c1)
	close(c2)
	close(c3)

	var c1Count, c2Count, c3Count int

	for i := 0; i < 1000; i++ {
		select {
		case <-c1:
			c1Count++
		case <-c2:
			c2Count++
		case <-c3:
			c3Count++
		}
	}

	fmt.Println("c1Count:", c1Count)
	fmt.Println("c2Count:", c2Count)
	fmt.Println("c3Count:", c3Count)
}
