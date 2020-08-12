package main

import (
	"fmt"
)

// 定义一个结构体
type person6 struct {
	// 定义属性(成员变量)
	id int
	name string
	age int
}

// 学生,继承了person6
type student5 struct {
	person6
	score int
}

// 学生会，继承学生
type xueshenghui struct {
	student5
	kaihui string // 学生会有开会的属性
}

func main() {
	/*
	使用自动推导：
		必须按顺序来：学生会--学生的属性--person6的属性--学生会的属性
		x2 := xueshenghui{student5:student5{score:98,person6:person6{id:2,name:"hallen",age:19}},kaihui:"开会中"}
	 */
	var x xueshenghui
	x.id = 1
	x.name = "hl"
	x.age = 18
	x.score = 89
	x.kaihui = "开会中"
	fmt.Println(x)		// {{{1 hl 18} 89} 开会中}

	// 自动推导初始化
	// 必须按顺序来：学生会--学生的属性--person6的属性--学生会的属性
	x2 := xueshenghui{student5:student5{score:98,person6:person6{id:2,name:"hallen",age:19}},kaihui:"开会中"}
	fmt.Println(x2)
}
