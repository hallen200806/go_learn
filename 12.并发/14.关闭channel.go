package main

import "fmt"

func main() {

	/*

	关闭管道：关闭后无法再发送数据，但是可以继续接收数据，指到channel为空
		close(ch)

	接收的时候会有两个参数，一个是管道发送的数据，一个是ok的bool值，false则管道关闭，true管道没有关闭

	 */


	// 创建一个无缓冲区的channel
	ch := make(chan int)
	go func() {
		for i:=0;i<5;i++{
			ch <- i
		}

		// 关闭channel
		close(ch)
	}()

	for {
		num, ok := <-ch
		if ok == true{
			fmt.Println(num)

		}else { // 管道关闭了，退出循环
		fmt.Println("管道关闭了")
			break
		}
	}


}
