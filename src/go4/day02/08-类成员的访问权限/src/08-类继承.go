package src

import "fmt"

//在go语言中，权限都是通过首字母大小来控制
//1. import ==》 如果包名不同，那么只有大写字母开头的才是public的
//2. 对于类里面的成员、方法===》只有大写开头的才能在其他包中使用

type Human struct {
	//成员属性:
	Name   string
	Age    int
	Gender string
}

// 在类外面绑定方法
func (this *Human) Eat() {
	fmt.Println("this is :", this.Name)
}

// 定义一个学生类，去嵌套一个Hum
type Student1 struct {
	Hum    Human //包含Human类型的变量, 此时是类的嵌套
	Score  float64
	School string
}

// 定义一个老师，去继承Human
type Teacher struct {
	Human          //直接写Huam类型，没有字段名字
	Subject string //学科
}
