package main

import "fmt"

func main()  {
	/*
	区别：
		字符：单引号，一个字符，出了\n \t等转义字符
		字符串：双引号，不限字符，隐藏了\0结束符
	 */
	var a byte = 'a'
	var b string = "a"

	fmt.Println(a)
	fmt.Println(b)
}
