package main

import "fmt"

type cat1 struct {
	name string
	age int
}

func (c cat1)greet1()  {	// 值传递
	name := c.name
	fmt.Println(name)
	// 这里修改不会修改原来的对象属性，可以通过指针或者返回值的方式
	c.age = 3
	fmt.Println(c.age)
}
// 指针的方式改变对象的属性值,地址传递
// 建议这样
func (c *cat1)greet2()  {
	c.name = "小兰"
	fmt.Println(c.name)
}
func main() {
	// 结构体对象方法中访问结构体属性
	c := cat1{name:"小花",age:2}
	c.greet1()
	fmt.Println(c)
	c.greet2()
	fmt.Println(c)
}
