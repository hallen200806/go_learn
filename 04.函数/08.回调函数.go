package main

import "fmt"

type FuncType func(int, int) int

func Clas(a,b int,ftest FuncType) (result int) {
	result = ftest(a,b)
	return
}

func Add(a,b int) int  {
	return a + b
}
func main() {

	/*
	回调函数：函数有一个参数是函数类型，这个函数就是回调函数

	多态：多种形态，调用同一个接口，不同的表现，可以实现不同的实例
	 */

	 a := Clas(1,2,Add)
	 fmt.Println(a)
}
