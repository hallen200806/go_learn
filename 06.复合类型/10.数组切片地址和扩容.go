package main

import "fmt"

func main() {
	/*
	当使用append向切片中添加的使用，内存地址可能会发生改变
	 当使用append追加数据的时候，长度超过容量，自动扩容
	*/

	//a := []int{0}
	var a []int
	// 查看内存地址
	fmt.Printf("%p\n",a)	// 0x0
	a = append(a,12,3)
	fmt.Println(a)
	fmt.Printf("%p\n",a)		// 0xc00000a0f0

	// 容量
	fmt.Println(cap(a))  // 2
}
