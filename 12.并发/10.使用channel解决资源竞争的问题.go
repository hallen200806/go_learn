package main

import (
	"fmt"
	"time"
)


// 全局变量，创建一个channel
var ch = make(chan int)

// 需求：Person11执行完，再执行Person22

func printer1(str string)  {
	for i:=0;i<len(str) ;i++  {
		fmt.Printf("%c",str[i])
		time.Sleep(time.Second)
	}
	fmt.Printf("\n")
}

func Person11()  {
	// 使用打印机
	printer1("hello")

	// 给管道发送数据，发送
	ch <- 666   // 因为前面定义的是int类型，这里发送的必须也是int类型
}

func Person22(){
	// 从管道读取数据,接收，如果通道没有数据，就会阻塞
	<-ch
	printer1("world")
}
func main() {
	/*
	channel解决资源争抢的问题：相当于一个执行完给下一个执行者一个信号，我执行完了，你继续吧

		Person11先执行：给管道发送数据，发送数据，这个是执行完以后定义的
			ch <- 666   // 因为前面定义的是int类型，这里发送的必须也是int类型

		再执行Person22：从管道读取数据,接收数据，这个是最前面定义的
			<-ch		从管道读取数据,接收，如果通道没有数据，就会阻塞
	 */

	// 会先把Person1执行完再执行Person2的
	go Person11()
	go Person22()


	// 这里的目的是不中断主协程
	for  {
		time.Sleep(time.Second)
	}


}