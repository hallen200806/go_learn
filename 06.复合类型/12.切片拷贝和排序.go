package main

import "fmt"

func main() {
	/*
	拷贝：
		得有一个和要拷贝的一样的空间，才能拷贝原来的，
		只有一个值的，那就只拷贝一个
	对任何一个进行修改，都不会影响其他的
	 */
	a := []int{1,2,3,4,5}
	var b []int = []int{0}	//只有一个值的，那就只拷贝一个
	fmt.Println(b)
	copy(b,a)
	fmt.Println(b)

	// 开辟一个和a一样的空间，全是0
	c := make([]int,5)
	fmt.Println(c)
	// 拷贝a到c
	copy(c,a)
	//对任何一个进行修改，都不会影响其他的
	c[2] = 999
	a[3] = 888
	fmt.Println(c)	// 和a一样
	fmt.Println(a)


}
