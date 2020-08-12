package main

import "fmt"

type Dog2 struct {
	Id int       // 属性
	Name string
	Age int
}

func (d *Dog2) Run()  {
	fmt.Println(d.Name)
	
}
func main() {

	/*
	两种作用域:
		1.包名
		2.函数名，首字母大写其他地方可以使用，小写只能本文件使用
	 */

	dog := Dog2{Id:2,Name:"tom",Age:3}
	dog.Run()
}
