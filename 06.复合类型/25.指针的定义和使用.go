package main

import "fmt"

func main() {
	/*
	指针：保存某个变量的地址，  *int 保存int的类型       **int  保存*int的类型
		定义：
			var 变量名 *类型
				var p *int  这里的int取决于你要获取那个变量地址的类型
		赋值：
			直接赋值：
				p = &i	这里使用的&变量名    指针变量指向谁，就把谁的地址赋值给指针变量
			new函数：可以自动推导
				var p1 *int
				p1 = new(int)
				*p1 = 57		*p1 操作的不是p1的内存，而是p1所指向的内存
				fmt.Println(*p1)	// 57
				fmt.Println(p1)		// 0xc00000a110

		空指针：
			var p *int		空指针 <nil>
		野指针：
			var p *int
			*p = 1111	没有指向地址，指向了未知的空间，不允许出现，会报错

		不要操作没有合法指向的内存
	*/

	var i int = 12
	fmt.Println(i)
	fmt.Printf("%p\n",&i)

	// 定义指针
	var p *int
	fmt.Println(p)	// 空指针 <nil>
	// 把变量i的地址赋值给指针变量p
	p = &i
	fmt.Println(p)

	var p1 *int
	p1 = new(int)
	*p1 = 57
	fmt.Println(*p1)	// 57
	fmt.Println(p1)		// 0xc00000a110
}

