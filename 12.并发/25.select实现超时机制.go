package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	quit := make(chan bool)




	go func() {
		for {
			select {
			case num := <-ch:
				fmt.Println(num)
			case <-time.After(3 * time.Second):		// 读取数据，如果阻塞3秒没有读取到数据
				fmt.Println("超时")
				quit <- true
				
			}
		}
	}()

	// 往channel写数据
	for i:=0; i<5; i++{
		ch <- i
		time.Sleep(time.Second)
	}


	<-quit  // 如果读取到true

	fmt.Println("程序结束")
}
