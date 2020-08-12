package main

import "fmt"

func main() {
	/*
	类型转换，在要转换的值前面加上要转换的类型
		1:低位数转为高位数，比如float32转为float64
		2.整型转为浮点型
		3.bool类型不能转为int
	*/
	a := 123
	b := float64(a)
	fmt.Println(b)
	fmt.Printf("%T\n",b) // float64

	c,d,e := 12,14,15
	// 输出平均值，保留小数
	fmt.Println((c+d+e)/3) // 13,整数
	fmt.Println(float64(c+d+e)/3) //13.666666666666666
	
}
