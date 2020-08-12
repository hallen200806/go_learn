package main

import "fmt"

func add(a,b int) int {
	return a + b
}

func subs(a,b int) int {
	return a - b
}

// 函数也是一种类型，通过type给一个函数类型起名
type funcType func(int, int)int  // 没有函数名，没有大括号{},funcType是一种函数类型
func main() {
	/*
	函数也是一种类型，通过type给一个函数类型起名
		没有函数名，没有大括号{},funcType是一种函数类型
		参数和要调用的函数参数一致
			type funcType func(int, int)int
	 */

	// 使用函数类型
	var funtest funcType

	// add调用
	funtest = add
	ret_add := funtest(1,2)
	fmt.Println(ret_add)


	// 也可以使用其他的函数
	funtest = subs
	ret_subs := funtest(2,3)
	fmt.Println(ret_subs)
	
}
