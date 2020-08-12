package main

import "fmt"

func init()  {
	fmt.Println("init")
}
func main() {
	/*
	init:
		先执行导入的包里面的init函数，然后再执行自身的init函数
	 */
}
