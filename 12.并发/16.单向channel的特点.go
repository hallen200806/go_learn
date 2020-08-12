package main

func main() {

	/*
	var ch1 chan int   		这是一个正常的channel，不是单向的
	var ch2 chan<- int		这是一个单向的channel，只能用于写数据
	var ch3 chan <-chan int  这是一个单向的channel，只能用于读数据


	chan<- 表示数据进入管道，要把数据写进管道，对于调用者就是输出
	<-chan 表示数据从管道出来，对于调用者就是获取管道的数据，就是输入

	注意：
		可以将channel隐式转换为单向队列，只收或只发，不能将单向channel转换为普通的channel
	 */

	 // 创建一个channel，双向的
	 ch := make(chan int)

	 // 双向的隐式转换为单向的写入管道,只能写，不能读
	 var write_ch chan<- int = ch
	 var read_ch <-chan int = ch

	 // 只能写
	 write_ch <- 666

	 //只能读
	<- read_ch

	// 单向的不能转换为双向的
	//var ch3 chan int = write_ch
}
