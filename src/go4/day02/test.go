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
	s.father.name = "father1"
	fmt.Println("s.School:", s.School)
	fmt.Println(s.name)
	fmt.Println(s.Age)
}
