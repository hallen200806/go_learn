package main

import "fmt"

func main() {
	/*
	匿名函数：没有函数名的函数
		变量名 = func(){}
	 */

	num := 10
	// 第一种：使用变量接收，在下面调用
	f :=func (){
		num ++
		fmt.Println(num)
	}
	// 函数调用
	f() //11

	// 第二种：不使用变量接收，直接调用
	func (){
		num ++
		fmt.Println(num)
	}()


	// 有参数和返回值的匿名函数
	d := func(a,b int)(c int){
		//fmt.Println(a+b)
		return a+b
	}(2,3)
	fmt.Println(d)

}
