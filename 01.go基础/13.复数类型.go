package main

import "fmt"

func main() {
	/*
	复数就是两个两个实数，两个浮点数，一个表示实部(real),一个表示虚部(imag)
		complex64：
		complex128：推导默认
	 */
	var a complex64
	a = 1.2 + 1.3i
	fmt.Println(a)

	b := 1.4 + 1.5i
	fmt.Println(b)
	fmt.Printf("%T\n",b)
	// 取出实部和虚部的值
	real_value := real(b)
	imag_value := imag(b)
	fmt.Println(real_value)
	fmt.Println(imag_value)

}
