package main

import "fmt"

// 定义一个结构体
type person struct {
	// 定义属性(成员变量)
	id int
	name string
	age int
}

// 匿名字段(继承person结构体)
// 学生继承person这个结构体
type student struct {
	person	//匿名字段,拥有了上面的所有字段
	score int
}

func main() {
	/*
	类就是结构体，结构体中可以定义成员变量(属性)，成员方法(函数)
	 Student{Person:Person{id:1,name:"小明",age:18},score:80}
		注意：Student里面是Person:Person
	指定初始化成员，如果没有初始化的部分，使用默认值
		子类中没有初始化父类中的属性，就是使用默认值

	*/
	s1 := student{person:person{id:1,name:"小明",age:18},score:80}
	fmt.Println(s1)
	// 没有初始化的属性使用默认值，这里的父类的字段值都为默认值
	s2 := student{score:90}
	fmt.Println(s2)

	// 修改属性值
	s1.age = 20
	s1.score = 99
	s2.id = 2
	s2.person.name = "hl"
	fmt.Println(s1)
	fmt.Println(s2)
	// 可以对整个父类的字段进行修改
	s1.person = person{id:11,name:"hl",age:17}
	fmt.Println(s1)

}
