package main

import "fmt"

func main() {
	/*
	数组： 是指一系列同一类型数据的集合
		定义：var 数组名 [长度]类型
		数组的赋值：看下面的实例
		数组的初始化：var a [长度]int = [长度]int{数组的元素，逗号隔开}
	 */

	var a [10]int // 默认的都是0，[0 0 0 0 0 0 0 0 0 0]
	fmt.Println(a)
	fmt.Println(len(a))
	// 数组的赋值
	// 1.通过下标赋值
	a[0] = 1
	fmt.Println(a)
	// 2.通过for循环
	for i :=0;i<len(a) ;i++  {
		a[i] = i+1
	}
	fmt.Println(a)

	// 获取数组的元素
	for i :=0;i<len(a) ;i++  {
		fmt.Println(a[i])
	}
	// range获取数组的元素
	for _,j := range a{
		fmt.Println(j)
	}
}
