package main

import "fmt"

type student5 struct {
	id int
	name string
	age int
	score int
}

func test_slice(arr []student5) ([]student5) {
	fmt.Println(arr)
	arr[0].name = "张三222"
	return arr
}
func main() {
	// 地址的传递，形参的改变会改变实参
	arr5 := []student5{
		{id:1,name:"张三",age:18,score:90},
		{id:2,name:"李四",age:18,score:90},
		{id:3,name:"王五",age:18,score:90},
		{id:4,name:"王五1",age:18,score:90},
		{id:5,name:"王五2",age:18,score:90},

	}
	arr_ret := test_slice(arr5)
	fmt.Println(arr_ret)
	// 地址的传递，形参的改变会改变实参
	fmt.Println(arr5)
}
