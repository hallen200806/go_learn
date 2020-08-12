package main

import "fmt"

type student4 struct {
	id int
	name string
	age int
	score int
}
func main() {
	// 这里不给定长度
	arr4 := []student4{
		{id:1,name:"张三",age:18,score:90},
		{id:2,name:"李四",age:18,score:90},
		{id:3,name:"王五",age:18,score:90},
		{id:4,name:"王五1",age:18,score:90},
		{id:5,name:"王五2",age:18,score:90},

	}
	fmt.Println(arr4)
}
