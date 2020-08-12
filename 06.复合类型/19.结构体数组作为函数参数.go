package main

import "fmt"

// 定义一个结构体
type student3 struct {
	id int
	name string
	age int
	score int
}

func test_type_arr(a [3]student3) ([3]student3) {
	fmt.Printf("===================\n")
	fmt.Println(a)
	a[0].name = "张三1111"
	return a
}
func main() {
	/*
	结构体数组是值传递
	 */

	arr3 := [3]student3{
		{id:1,name:"张三",age:18,score:90},
		{id:2,name:"李四",age:18,score:90},
		{id:3,name:"王五",age:18,score:90},

	}
	b := test_type_arr(arr3)
	fmt.Printf("----------------------\n")
	fmt.Println(b)
	// 值传递，形参改变不会改变实参的值
	fmt.Println(arr3)
}
