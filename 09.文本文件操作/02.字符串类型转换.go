package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
	go语言不允许隐式转换，所有的转换必须显示声明，而且转换只能发生在两种相互兼容的类型之间
		1.字符串转换为字符切片
			slice := []byte(str)
		2.字符切片转换为字符串
			string(slice_a)

	一.使用模块包：strconv
		1.将其他类型装换为字符串
			1.1.bool转换为字符串：FormatBool
				a1 := false
				str1 := strconv.FormatBool(a1)
			1.2.int转字符串：FormatInt
				str2 := strconv.FormatInt(10,10)
					第一个参数是要转换的整数
					第二个参数是要转换的进制，这里用的10进制
			1.3：float转字符串：
					strconv.FormatFloat(3.14,'f',4,32)
					要转换的小数
					打印方式，f是指以小数的方式
					保留几位小数
					进制
			1.4.Itoa
				str4 := strconv.Itoa(314) // 这里是整型
		2.字符串转为其他类型
			2.1.字符串转为bool类型
				ret5,err := strconv.ParseBool("true")
				fmt.Println(ret5)
				fmt.Println(err) // 转换成功，err为nil，否则不为nil
			2.2字符串转换为整型
				ret6,err := strconv.ParseInt("1234",16,64)
					目标字符串
					进制
					int长度
				fmt.Println(ret6,err) // 转换成功，err为nil，否则不为nil
				出错了ret6则为0
			2.3字符串转为浮点类型
				ret7,err := strconv.ParseFloat("3.14",64)
					目标字符串
					float长度
				fmt.Println(ret7,err)	// 转换成功，err为nil，否则不为nil
			2.4.Atoi
				ret8,err := strconv.Atoi("123")
				fmt.Println(ret8,err)	// 转换成功，err为nil，否则不为nil

	*/

	// 字符串转换为字符切片
	str := "hello world"
	slice := []byte(str)
	fmt.Println(slice)	// [104 101 108 108 111 32 119 111 114 108 100]

	// 字符切片转换为字符串
	slice_a := []byte{'h','e','l','l','o'}
	fmt.Println(string(slice_a))

	// bool转换为字符串
	a1 := false
	str1 := strconv.FormatBool(a1)
	fmt.Println(str1)
	fmt.Printf("%T\n",str1) // string

	// int转字符串
	str2 := strconv.FormatInt(10,10)
	fmt.Println(str2)
	fmt.Printf("%T\n",str2) // string

	// float转字符串
	str3 := strconv.FormatFloat(3.14,'f',4,32)
	fmt.Println(str3)

	str4 := strconv.Itoa(314) // 这里是整型
	fmt.Println(str4)


	// 字符串类型转为bool
	ret5,err := strconv.ParseBool("true")
	fmt.Println(ret5)
	fmt.Println(err) // 转换成功，err为nil，否则不为nil

	// 字符串转换为整型
	ret6,err := strconv.ParseInt("1234",16,64)
	fmt.Println(ret6,err)

	// 字符串转为浮点类型
	ret7,err := strconv.ParseFloat("3.14",64)
	fmt.Println(ret7,err)

	// atoi
	ret8,err := strconv.Atoi("123")
	fmt.Println(ret8,err)
}
