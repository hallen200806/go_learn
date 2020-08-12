package main

import "fmt"

func main()  {

	var a string = "hello world"
	fmt.Println(a)

	b := "hello b"
	fmt.Println(b)
	fmt.Printf("%T\n",b)
	// 取角标只是取得ASCII码值，可以通过printf打印
	fmt.Println(a[1])	// 结果是101，是对应的ASCII码值
	fmt.Printf("%c\n",a[1]) // 结果是e


	// 下面一个是字符a，一个是字符串a
	c := 'a'
	d := "a"  // 'a'+ '\0' 组成
	fmt.Println(c)
	fmt.Println(d)
	// len计算字符串长度,不包含'\0'的
	e := "abc"
	fmt.Println(len(e))
	// 一个汉字占3个字节，所以长度是12
	f := "知了课堂"
	fmt.Println(len(f))

	// 字符串拼接
	h := e + f

	fmt.Println(h)
}
