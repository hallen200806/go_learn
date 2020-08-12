package main

import "fmt"

// 定义一个结构体
type person4 struct {
	// 定义属性(成员变量)
	id int
	name string
	age int
}
// 定义一个结构体
type person5 struct {
	// 定义属性(成员变量)
	addr string
	gender string
}
// 学生,继承了person4和person5
type student4 struct {
	person4	//先继承的先初始化
	person5
	score int
}

func main() {
	// 注意：继承多个结构体，初始化的时候要注意：先继承的先初始化，否则会报错
	var s student4
	s.id = 1
	s.age = 18
	s.name = "hl"
	s.addr = "市区"
	s.gender = "男"
	s.score = 97
	fmt.Println(s)	//{{1 hl 18} {市区 男} 97}


	// 自动推导初始化：先继承的先初始化
	s1 := student4{person4:person4{id:2,name:"hallen",age:18},person5:person5{addr:"市区",gender:"男"},score:99}
	fmt.Println(s1)
}
