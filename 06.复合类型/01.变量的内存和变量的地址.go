package main

import "fmt"

func main() {
	/*
	每个变量有两个含义：
		变量的内存：	教室
		变量的地址(也叫指针)：	教室外面的门牌号
	 */

	 var a int
	 fmt.Println(a)     // 变量的内存
	 fmt.Println(&a)	// 变量的地址，也叫指针


}
