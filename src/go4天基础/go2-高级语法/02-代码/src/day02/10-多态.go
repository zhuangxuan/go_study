package main

import "fmt"

//实现go多态，需要实现定义接口
//人类的武器发起攻击，不同等级子弹效果不同

// 定义一个接口, 注意类型是interface
type IAttack interface {
	//接口函数可以有多个，但是只能有函数原型，不可以有实现
	Attack()
	//Attack1()
}

// 低等级
type HumanLowLevel struct {
	name  string
	level int
}

func (a *HumanLowLevel) Attack() {
	fmt.Println("我是:", a.name, ",等级为:", a.level, ", 造成1000点伤害")
}

// 高等级
type HumanHighLevel struct {
	name  string
	level int
}

func (a *HumanHighLevel) Attack() {
	fmt.Println("我是:", a.name, ",等级为:", a.level, ", 造成5000点伤害")
}

// 定义一个多态的通用接口,传入不同的对象，调用同样的方法，实现不同的效果 ==》 多态
func DoAttack(a IAttack) {
	a.Attack()
}

func main() {
	//var player interface{}
	var player IAttack //定义一个包含Attack的接口变量

	lowLevel := HumanLowLevel{
		name:  "David",
		level: 1,
	}

	highLevel := HumanHighLevel{
		name:  "David",
		level: 10,
	}

	lowLevel.Attack()
	highLevel.Attack()

	//对player赋值为lowLevel，接口需要使用指针类型来赋值
	player = &lowLevel
	player.Attack()

	player = &highLevel
	player.Attack()

	fmt.Println("多态......")
	DoAttack(&lowLevel)
	DoAttack(&highLevel)
}
