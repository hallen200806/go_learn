package main

import "fmt"

type cat4 struct {
	name string
	age int
}
type dog2 struct {
	name string
	age int
}

func (c4 cat4)greet()  {
	fmt.Println("这是cat的方法")
}
func (d2 dog2)greet()  {
	fmt.Println("这是dog的方法")
}
func main() {
	// 闭包会用到方法值和方法表达式

	c := cat4{name:"小花",age:2}
	d := dog2{name:"小兰",age:3}
	d.greet()
	c.greet()
	// 方法值
	cat := c.greet
	cat()
	fmt.Println(cat)
	fmt.Printf("%T",cat) // 类型是func()

	// 方法表达式
	dog := dog2.greet // 这里是结构体.方法名
	dog(d)  // 传的是对象
}
