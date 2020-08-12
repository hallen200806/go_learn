package main

import "fmt"
/*
	递归函数：调用函数自己本身
		递归相同的结构

*/


func test10(a int)  {
	if a == 1 {
		fmt.Println("匹配了")
		fmt.Println(a)
		// 结束递归
		return
	}
	fmt.Println(a)
	test10(a-1)
}
func main() {

	test10(5)
}
