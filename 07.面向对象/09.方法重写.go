package main

import "fmt"

type animal1 struct {
	name string
	age int
}

type cat3 struct {
	animal1 // 继承animal
	color string
}

// animal的方法
func (a animal1)greet()  {
	//fmt.Println(a.name)
	fmt.Println("这是父类animal的方法")
}

func (c cat3)greet()  {
	fmt.Println("这是子类cat的方法")
}

func main() {
	//a := animal1{name:"动物",age:12}
	//a.greet()
	c := cat3{animal1:animal1{name:"猫",age:2},color:"蓝色"}
	c.greet()	// 父类和子类的方法重名，就重写了父类的方法
	c.animal1.greet()

	
}
