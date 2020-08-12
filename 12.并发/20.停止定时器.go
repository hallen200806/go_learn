package main

import (
	"time"
	"fmt"
)

func main() {

	timer := time.NewTimer(2 * time.Second)

	go func() {
		<- timer.C
		fmt.Println("定时时间到，协程执行开始")

	}()


	// 停止定时器
	timer.Stop()    // 停止后不会再打印上面的：定时时间到，协程执行开始
	// 原因，两个协程是同时执行的，上面的定时器要3秒，这个是立马执行，所以会把定时器停掉，所以不会执行上面的打印语句

	// 这里的目的是不中断主协程
	for  {
		time.Sleep(time.Second)
	}


}
