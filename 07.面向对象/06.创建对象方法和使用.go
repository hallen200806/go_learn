package main

import "fmt"

type cat struct {
	name string
	age int
}
type dog struct {
	name string
	age int
}

// 给cat结构体定义一个greet方法
func (c cat)greet()  {	// 对象名 c 随便起
	fmt.Println("喵喵叫")
}
func main() {
	/*
	定义对象方法：
		func (对象名 结构体名)方法名(){  // 对象名随便起
		}
	使用：
		使用结构体对象的变量去调用
			c1 := cat{name:"小花",age:2}
			c1.greet()
	注意：
		出了改结构体的对象，其他方式都不能使用
		greet()  不行
		d.greet() 不行
	 */
	c1 := cat{name:"小花",age:2}
	d := dog{name:"小兰",age:3}
	// 猫调用greet方法,必须使用cat的对象调用，其他方式都不能调用
	c1.greet()
	fmt.Println(d)
	// dog没有greet方法，不能调用
	//greet()   // 这种方式不行
	//d.greet() // 这种方式不行

}
