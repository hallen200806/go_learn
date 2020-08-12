package main

import (
	"fmt"
	"time"
)

func printer(str string)  {
	for i:=0;i<len(str) ;i++  {
		fmt.Printf("%c",str[i])
		time.Sleep(time.Second)
	}
	fmt.Printf("\n")
}

func Person1()  {
	// 使用打印机
	printer("hello")
}

func Person2(){
	printer("world")
}
func main() {
	/*
	有一个公共的打印机，多个人使用，会导致资源竞争的问题
	 */

	 // 不会先把Person1执行完再执行Person2的，会争抢执行,
	 go Person1()
	 go Person2()


	 // 这里的目的是不中断主协程
	for  {
		time.Sleep(time.Second)
	}

}
