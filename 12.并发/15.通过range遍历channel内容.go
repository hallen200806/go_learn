package main

import "fmt"

func main() {

	/*
	循环channel的两种方式：
		第一种：写一个死循环，然后判断channel是否关闭，如果关闭则退出
		第二种：循环chan对象

	注意：两个循环只能存在一个，循环第二次就没有值了，已经接收完了
	 */

	 ch := make(chan int)

	 // 发送channel
	 go func() {
	 	for i:=0; i<5; i++{
	 		ch <- i
		}
		//关闭channel
		close(ch)
	 }()

	 // 第一种循环channel
	 for {
	 	num,ok := <-ch
	 	if ok == true{
	 		fmt.Println(num)
		}else {
			break
		}
	 }

	 // 第二种循环

	 for data := range ch{
	 	fmt.Println(data)
	 }
}
