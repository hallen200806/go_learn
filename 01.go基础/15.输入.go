package main

import "fmt"

func main() {
	// 使用fmt.Scanf()

	var a int
	// 格式化输入
	fmt.Printf("请输入一个整数:")
	fmt.Scanf("%d",&a) // 将用户输入的值赋值给a变量

	fmt.Println(a)
	// 不格式输入的方式
	//fmt.Scan(&a)
}
