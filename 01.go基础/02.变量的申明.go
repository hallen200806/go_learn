package main

import "fmt"

/*
变量的申明:
	1.var 变量名 类型
		var name str
	2.

*/

func main()  {
	// 第一种 var 变量名 类型 = 值
	var age int = 10
	var name string = "hallen"
	fmt.Println(age)
	fmt.Println(name)

	// 第二种 ：var 变量名 类型，在下面再赋值
	var age1 int
	age1 = 18
	fmt.Println(age1)

	// 第三种：var 变量1,变量2 类型 = 值1,值2   或者在下面赋值
	//这里有几个变量就得有几个值
	var age2,age3 int = 19,20
	fmt.Println(age2,age3)

	// 第四种：类型自动推导： var a = 10 就会自动推导出类型为int
	var a = 10
	var b = "测试"
	// 打印类型用Printf
	fmt.Printf("%T",a)
	fmt.Println()
	fmt.Printf("%T",b)
	fmt.Println()

	// 省略var申明关键字：   c:=10
	c := 10
	fmt.Println(c)
	d := "hello world"
	fmt.Println(d)
	// 定义三个变量
	e,f,g := 11,"hello f",12.34
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

}