package main

import "fmt"

func main() {
	/*
	range语法：迭代打印每个元素，默认返回两个值，一个是元素的位置，一个是元素本身，
		如果只有一个接收参数，默认返回的是下标，可以使用匿名变量
	 */
	str := "abcde"
	for i,data := range str{
		fmt.Println(i)
		fmt.Println(data)
	}

	for i := range str{
		fmt.Println(i)
	}
}
