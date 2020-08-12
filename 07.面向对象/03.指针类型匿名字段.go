package main

import "fmt"

// 定义一个结构体
type person3 struct {
	// 定义属性(成员变量)
	id int
	name string
	age int
}

// 匿名字段(继承person结构体)
// 学生继承person这个结构体
type student3 struct {
	*person3	//匿名字段,拥有了上面的所有字段
	score int
}

func main() {
	/*
	student3使用person3的指针
		使用方式1：
			s.person3 = new(person3)
			然后才能使用.的方式赋值
		使用方式2：
			s2.person3 = &person3{id:11,name:"hl",age:18}
		取的方式：
			都是通过对象.属性获取
				s.id
				s.name

	*/
	var s student3
	s.score = 90
	//s.name = "hl" // 父类里面的属性，必须在new以后才可以使用.的方式赋值
	//s.id = 11 // 不能通过.的方式去赋值 panic: runtime error: invalid memory address or nil pointer dereference
	s.person3 = new(person3)
	s.id = 11
	s.name = "hl"
	fmt.Println(s)	//{0xc0000044a0 90}
	fmt.Println(s.id,s.score,s.name)

	// 方式2
	var s2 student3
	s2.person3 = &person3{id:11,name:"hl",age:18}
	fmt.Println(s2.age)
}
