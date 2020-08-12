package main

import "fmt"

func main()  {
	/*
	浮点类型：
		单精度类型：4字节长度
			float32:精确到小数点后7位
		双精度类型：8字节长度，自动推导的默认类型
			float64：精确到小数点后15位
	*/
	// 第七位四舍五入
	var a float32 = 10.211123124125241242
	//第十五位四舍五入
	var b float64 = 10.312312351231235132123
	fmt.Println(a)
	fmt.Println(b)

	// 自动推导的类型为float64
	c := 3.123124124
	fmt.Println(c)
}
