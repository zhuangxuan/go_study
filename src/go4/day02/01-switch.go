package main

import (
	"fmt"
	"os"
)

//从命令输入参数，在switch中进行处理22

func main() {
	//C: argc , **argv
	//Go: os.Args ==> 直接可以获取命令输入，是一个字符串切片 []string
	cmds := os.Args

	//os.Args[0] ==> 程序名字
	//os.Args[1] ==> 第一个参数 ，以此类推
	for key, cmd := range cmds {
		fmt.Println("key:", key, ", cmd:", cmd, ", cmds len:", len(cmds))

	}

	if len(cmds) < 2 {
		fmt.Println("请正确输入参数！")
		return
	}

	switch cmds[1] {
	case "hello":
		fmt.Println("hello")
		//go 的switch, 默认加上break了，不需要手动处理
		//如果想向下穿透的话，那么需要加上关键字: fallthrough
		fallthrough
	case "world":
		fmt.Println("world")
		fallthrough
	case "world2":
		fmt.Println("world2")
	default:
		fmt.Println("default called!")
	}
}
