package main

import "fmt"

// 不定长参数: 参数名 ...类型
// 注意，不定参数一定在最后，也就是位置参数的最后
func test3(b int,a ...int)  {

	fmt.Println(a)
	fmt.Printf("%T\n",a)
	// 循环出传的参数
	for i:=0;i<len(a) ;i++  {
		fmt.Println(a[i])
	}
}

func test44(a ...int)  {
	fmt.Println(a)
}
func test4(a ...int)  {
	// 不定长参数传个其他的不定长参数
	test44(a...)
	test44(a[:2]...)
	for i,data := range a{
		fmt.Println(i,data)
	}
}

func main() {
	test3(1,2,3,4)
	test4(1,2,3,4)
}
