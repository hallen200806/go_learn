package main

import "fmt"

type Animal struct {
	Color string
} 

type Dog3 struct {
	Animal  // 匿名字段
	Id int       // 属性
	Name string
	Age int
}

func (a *Animal) Eat()  {
	fmt.Println(a.Color)
	
}
func (d *Dog3) Run()  {
	fmt.Println(d.Name)

}


func main() {
	/*
	多态：
	同一个行为具有多个不同表现形式或形态的能力
	接口的多种不同的实现方式即为多态
		1.继承
		2.重写
		3.引用还是指向基类
	 */
	// 第一种初始化
	dog := Dog3{Animal:Animal{Color:"red"},Id:1,Name:"tom", Age:3}
	dog.Run()
	dog.Eat()

	// 第二种初始化
	var dog2 Dog3
	dog2.Id = 2
	dog2.Name = "tom2"
	dog2.Age = 3
	dog2.Color = "yellow"
	dog2.Run()
	dog2.Eat()
}
