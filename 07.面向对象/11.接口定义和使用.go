package main

import "fmt"

// 加法
type add struct {
	num1 int
	num2 int
}

// 加法类的方法
func (a *add)add_func() {
	fmt.Println(a.num1)
}

// 接口:一般后面加er结尾
type adder interface {
	// 只有方法，没有实现,由别的类型(自定义类型)实现
	add_func()
}



func main() {
	/*
		接口是一种数据类型，可以接受满足对象的信息
		接口是虚的，方法是实的
		接口定义规则，方法实现规则
		接口定义的规则，方法中必须实现，也就是必须有该方法
		将对象赋值给接口，必须满足接口中的方法的什么格式

	定义：
		公共方法组合起来的，以封装特定功能的集合
	特点：
		抽象
		封装
		多态

	多态的使用：
		1.
			var h adder
			h = &a		将对象信息赋值给接口类型变量
			h.add_func()	调用方法
		2.
			adder(&a)
	 */


	// 只要实现了此接口方法的类型，那么这个类型的变量(接收者类型)就可以给h赋值
	var h adder

	a := add{num1:10,num2:20}
	//ret_a := a.add_func()

	// 将对象信息赋值给接口类型变量
	h = &a      // 这里取决于类型是指针
	// 调用方法
	h.add_func()
}
