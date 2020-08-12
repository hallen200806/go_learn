package main

import "fmt"

func main() {
	/*
	数据的比较：
		只支持 == 或 !=
		比较是不是每个元素都一样
		两个数组比较，数组类型要一样，长度要一样

	同类型的数组可以赋值给另一个
	 */
	a := [6]int {1,2,3,4,5,6}
	b := [6]int {1,2,3,4,5,6}
	c := [6]int {1,2,3,5,6,7}
	//d := [3]int {1,2,3}
	fmt.Println(a == b)
	fmt.Println(a== c)
	//fmt.Println(a == d)   // 长度不一样，不能比较，会报错
	// 同类型的数组可以赋值给另一个
	a = c
	fmt.Println(a)

}
