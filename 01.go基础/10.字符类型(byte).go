package main

import "fmt"
func main()  {
	// 字节类型也叫字符类型

	/*
	单引号和双引号的区别：
		单引号：字符，byte类型
		双引号：字符串，string类型

	 */

	// 结果是97，a的ASCII码值为97，具体查看"ASCII码表"，0是48,9是57
	var a byte = 'a'
	fmt.Println(a)
	fmt.Printf("%c\n",a) //结果为a
	fmt.Printf("%c\n",97) //结果也是a
}
