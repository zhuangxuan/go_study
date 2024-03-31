package main

import "fmt"

//Person类，绑定方法：Eat，Run，Laugh, 成员
//public，private
/*
class Person {
public :
    string name
	int age

public :
	Eat() {
		xxx
	}
}

*/

type MyInt1 int

func (mi *MyInt1) printMyInt() {
	fmt.Println("MyInt value is:", *mi)
}

type Person struct {
	//成员属性:
	name   string
	age    int
	gender string
	score  float64
}

// 在类外面绑定方法
func (this *Person) Eat() {
	//fmt.Println("Person is eating")
	//类的方法，可以使用自己的成员
	//fmt.Println(this.name + " is eating!")
	this.name = "Duke"
}

func (this Person) Eat2() {
	fmt.Println("Person is eating")
	//类的方法，可以使用自己的成员
	this.name = "Duke"
}

func main() {
	lily := Person{
		name:   "Lily",
		age:    30,
		gender: "女生",
		score:  10,
	}

	lily1 := lily

	fmt.Println("Eat，使用p *Person，修改name的值 ...")
	fmt.Println("修改前lily:", lily) //lily
	lily.Eat()
	fmt.Println("修改后lily:", lily) //Duke

	fmt.Println("Eat2，使用p Person，但是不是指针 ...")
	fmt.Println("修改前lily:", lily1) //lily
	lily1.Eat2()
	fmt.Println("修改后lily:", lily1) //lily

	var myint1 MyInt1 = 100
	myint1.printMyInt()
}
