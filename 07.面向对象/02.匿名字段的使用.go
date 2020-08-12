package main

import "fmt"

// 定义一个结构体
type person2 struct {
	// 定义属性(成员变量)
	id int
	name string
	age int
}

// 匿名字段(继承person结构体)
// 学生继承person这个结构体
type student2 struct {
	person2	//匿名字段,拥有了上面的所有字段
	score int
	name string	// 父类中也有name属性
}

func main() {
	var s student2
	// 父类和子类中都有name，这里的name是子类的，优先赋值子类
	s.name = "hl"
	// 这样才是父类的name
	s.person2.name = "hallen"
	fmt.Println(s)
}
