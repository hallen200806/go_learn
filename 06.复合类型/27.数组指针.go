package main

import "fmt"

func main() {
	/*
	数组指针的定义：
		var p *[长度]类型
	通过指针操作数组：
		(*p1)[0] = 111	// 第一种方式：这里*p1必须用括号括起来，表示一个整体
		p1[1] = 222		// 第二种方式
	 */

	arr := [3]int{1,2,3}
	fmt.Println(arr)
	var p *[3]int
	p = &arr
	fmt.Println(&p)	// 0xc000098020
	fmt.Println(p)	// &[1 2 3]

	// 自动推导赋值
	p1 := &arr
	fmt.Println(&p1)
	fmt.Println(*p1)

	// 通过指针操作数组
	(*p1)[0] = 111	// 第一种方式：这里*p1必须用括号括起来，表示一个整体
	p1[1] = 222		// 第二种方式
	fmt.Println(*p1) // 把p1的第一个值改为了111
}
