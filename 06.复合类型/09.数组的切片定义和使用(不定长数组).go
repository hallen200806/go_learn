package main

import "fmt"

func main() {

	/*
	数组切片的定义：
		不固定数组的长度，后面可以随意添加
		1.var slice []int
		2.var slice []int = []int{1,2,3}
		3.自动推导
			slice := []int{1,2,3}
		4.make函数实现
			slice := make([]int,5,10)	5是指长度，默认为5个0，,10是指容量
			有了长度和容量，会默认设置为5个0，后面追加这几个0也会存在，
			可以都设置为0，后面追加的就是最新的
		根据下标修改值
		slice[0] = 10
	*/
	//第一种
	var slice []int
	// 这里必须有个变量接收，不然会报错
	slice = append(slice,1,2,3)
	fmt.Println(slice)
	// 第二种
	var slice2 []int = []int{1,2,3,4}
	slice2 = append(slice2,5,6,7)
	fmt.Println(slice2)
	// 第三种
	slice3 := []int{1,2,3}
	slice3 = append(slice3,7,8,9)
	fmt.Println(slice3)
	// 第四种
	slice4 := make([]int,0,0)   // 长度 容量
	slice4 = append(slice4,1,2,3,4)
	fmt.Println(slice4)

	// 根据下标修改值
	slice4[0] = 10
	fmt.Println(slice4)
}
