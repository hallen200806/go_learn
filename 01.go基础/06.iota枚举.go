package main

import "fmt"

func main()  {
	/*
	iota:
		是一个常量的自动生成器，自动加1
	 	给常量赋值使用的
		遇到了const就会重置为0
		同一个const中的所有iota都会自动加1
		同一个const中，如果一次定义多个变量(同一行)，都是0
	*/
	const (
		a = iota	// 0
		b = iota	// 1
		c = iota	// 2
	)
	// iota 遇到了const就会重置为0
	const d  = iota

	const (
		// 全是0
		e,f,g,h = iota,iota,iota,iota
	)
	fmt.Println(a,b,c)
	fmt.Println(d)
	fmt.Println(e,f,g,h)
}
