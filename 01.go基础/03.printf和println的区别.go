package main

import "fmt"

func main()  {
	// printf是格式化输出，但是没有换行,后面可以加个\n换行符
	a,b := 10,12.3

	fmt.Println("a=",a)
	fmt.Printf("hello world%c",'\n') // 以\开头的字符都是字符，单引号，转义字符
	fmt.Printf("b = %f\n",b)
	/*
	%后面的参数说明：
		1. %d:整型
		2. %f:浮点类型
		3. %T:查看类型

	*/
}
