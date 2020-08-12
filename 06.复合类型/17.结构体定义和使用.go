package main

import "fmt"

// 定义结构体
type student struct { // 结构体名如果在其他文件使用则首字母得大写
	d int
	name string
	age int
	gender string
	addr string
}

func main() {
	/*
	定义：
		结构体定义在函数外面
		type 结构体名 struct{  // 结构体如果在其他文件使用，首字母得大写
			结构体成员列表（成员名 数据类型）
			定义的顺序和打印的顺序一致
			name string
			age int
		使用：
			第一种：var 变量名 结构体名 = 结构体名{...}
			第二种：var 变量名 结构体名
					变量名.字段名 = xxx
			第三种：自动推导：变量名 := 结构体名{...}
			第四种：通过new申请一个结构体
				p := new(结构体名)

		}
	 */

	// 第一种使用结构体
	var s1 student = student{id:1,age:18,name:"张三",gender:"男",addr:"市区"}
	fmt.Println(s1)

	// 第二种使用结构体
	var s2 student
	s2.id = 2
	s2.name = "张三2"
	s2.age = 18
	s2.addr = "市区"
	fmt.Println(s2)

	// 第三种使用结构体：自动推导
	s3 := student{id:3,name:"李四",age:20,gender:"男",addr:"郊区"}
	fmt.Println(s3)

	
}
