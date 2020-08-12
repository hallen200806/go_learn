package main

import "fmt"

func main()  {
	// 多个变量的定义，这种只能是同一种类型
	var a,b int
	a=10
	b = 20
	var c,d int = 30,40
	e,f := 50,60
	fmt.Println(a,b,c,d,e,f)



	// 多个类型变量的定义
	var (
		g int = 10 // 这里没有逗号的
		h string = "hello"
	)
	// 这种需要赋值，如果只申明变量，而不赋值，可以使用上面的
	i,l := 10,"hello"

	fmt.Println(g,h,i,l)

	// 常量的定义：用const
	const j int = 15
	const name  = "hello"
	const (
		o = 10
		k = "hello"
	)
	fmt.Println(j)
	fmt.Println(name)
	fmt.Println(o)
	fmt.Println(k)
}
