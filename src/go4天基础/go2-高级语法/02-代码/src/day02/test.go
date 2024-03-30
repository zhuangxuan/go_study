package main

import "fmt"

type father struct {
	name string
	Age  int
}

type Son struct {
	father
	School string
}

func main() {
	s := Son{}
	fmt.Println("s.School:", s.School)
	fmt.Println(s.name)
	fmt.Println(s.Age)
}
