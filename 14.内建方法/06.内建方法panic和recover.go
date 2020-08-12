package main

import "fmt"

func test_panic_recover()  {

	// 捕获异常，配合defer使用
	defer func() {
		msg := recover()
		fmt.Println(msg)
	}()
	// 抛出异常
	panic("this is a panic")

}
func main() {
	// panic:抛出异常
	// recover:捕获异常,配合defer使用, msg := recover() ，接收值接收异常信息
	test_panic_recover()

}
