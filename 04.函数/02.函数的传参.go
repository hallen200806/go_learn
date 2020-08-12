package main

import "fmt"

// 普通参数
func test1(a int,b int){  // 如果蚕食类型相同，可以 a,b int这样定义
	fmt.Println(a,b)
}

// 不定长参数: 参数名 ...类型
// 注意，不定参数一定在最后，也就是位置参数的最后
func test2(b int,a ...int)  {

	fmt.Println(a)
	fmt.Printf("%T\n",a)
	// 循环出传的参数
	for i:=0;i<len(a) ;i++  {
		fmt.Println(a[i])
	}
}

// 调用test2传参
func test5(c int,b ...int)  {
	test2(c,b[:4]...)
}
func main() {
	test1(1,2)
	test2(1,2,3,4)
	test2(2,1,2,3,4,5)
}
