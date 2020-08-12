package main

import (
	"fmt"
	"time"
)

func main() {

	// 创建channel
	var ch = make(chan string)

	// defer延迟执行
	defer fmt.Println("主协程结束")

	go func() {
		defer fmt.Println("子协程调用结束")
		for i := 0; i<2 ;i++  {
			fmt.Println(i)
			time.Sleep(time.Second)
		}
		ch <- "我是子协程，我要执行结束了"
	}()

	// 主协程接着工作
	str := <-ch
	fmt.Println(str)
}
