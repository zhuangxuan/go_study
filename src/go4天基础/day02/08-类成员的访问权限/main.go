package main

import (
	"day02/08-类成员的访问权限/src"
	"fmt"
)

func main() {
	s1 := src.Student1{
		Hum: src.Human{
			Name:   "Lily",
			Age:    18,
			Gender: "女生",
		},
		School: "昌平一中",
	}

	fmt.Println("s1.name:", s1.Hum.Name)
	fmt.Println("s1.school:", s1.School)

	t1 := src.Teacher{}
	t1.Subject = "语文"
	t1.Name = "荣老师" //下面这几个字段都是继承自Human
	t1.Age = 35
	t1.Gender = "女生"

	fmt.Println("t1 :", t1)
	t1.Eat()

	//继承的时候，虽然我们没有定义字段名字，但是会自动创建一个默认的同名字段
	//这是为了在子类中依然可以操作父类，因为：子类父类可能出现同名的字段
	fmt.Println("t1.Human.name:", t1.Human.Name)
}
