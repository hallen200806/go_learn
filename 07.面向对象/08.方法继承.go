package main

import "fmt"

type animal struct {
	name string
	age int
}

type cat2 struct {
	animal // 继承animal
	color string
}

// animal的方法
func (a animal)animal_func()  {
	fmt.Println(a.name)
	fmt.Println("这是父类animal的方法")
}

func (c cat2)cat_func()  {
	fmt.Println("这是子类cat的方法")
}
func main() {
	// 结构体继承里另外的结构体，属性和方法都会继承

	a := animal{name:"动物",age:12}
	c := cat2{animal:animal{name:"猫",age:2},color:"蓝色"}
	a.animal_func()
	c.cat_func()
	// 因为结构体cat继承了animal，所以它也有animal的方法
	c.animal_func()

	
}
