package main

import "fmt"

type hero2 struct {
	name string
	age int
	power int
}

func teststruct(a hero2) (hero2) {
	fmt.Println(a)
	a.name = "李四"
	return a
}
func main() {
	// 结构体作为参数是值传递，形参改变实参不会改变
	a := hero2{name:"张三",age:18,power:89}
	b := teststruct(a)
	fmt.Println(b)
	fmt.Println(a)
}
