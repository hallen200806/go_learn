package main

import "fmt"

// 定义一个结构体
type student2 struct {
	id int
	name string
	age int
	score int
}

func main() {
	// 结构体中使用数组:自动推导
	arr := [3]student2{
		{id:1,name:"张三",age:18,score:90},
		{id:2,name:"李四",age:18,score:90},
		{id:3,name:"王五",age:18,score:90},

	}
	fmt.Println(arr)

	//不使用自动推导
	var arr1 [3]student2 = [3]student2{
		{id:4,name:"张三1",age:18,score:90},
		{id:5,name:"李四1",age:18,score:90},
		{id:6,name:"王五1",age:18,score:90},
	}
	fmt.Println(arr1)
	// 取值
	fmt.Println(arr1[1])
	fmt.Println(arr1[1].name)
	// for循环取值
	for i:=0;i<len(arr1) ;i++  {
		fmt.Println(arr1[i])
	}

	// 修改值
	arr1[1].name = "李四222"
	fmt.Println(arr1)
}
