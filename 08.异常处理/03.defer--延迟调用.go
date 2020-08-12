package main

import "fmt"

func test03(i int)  {
	v := 100 / i
	fmt.Println(v)
}
func main() {
	/*
	defer：
		1.延迟调用
		2.只能使用在函数内
		3.后进先出的执行顺序
		4.函数或者某个延迟调用发生错误，这些错误依旧会被执行
		5.延迟调用还是会保持原来的赋值，
			而不会因为延迟执行就会因后面的值的改变而改变
	 */
	defer fmt.Println("hello")
	defer fmt.Println("老王")
	fmt.Println("你好老王")  // 会在defer之前执行
	defer test03(0)  // 最后执行
	defer fmt.Println("你好")



}
