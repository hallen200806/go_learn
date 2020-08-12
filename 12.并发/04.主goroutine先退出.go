package main

import (
	"fmt"
	"time"
)

func NewTask()  {
	for{
		fmt.Println("这是子协程")
		time.Sleep(time.Second)
	}
}
func main() {
	/*
	主协程退出，其他子协程也跟着退出
		主协程先退出导致子协程没有来得及调用
	 */

	 go NewTask()
	 for i :=0;i<100;i++ {
		fmt.Println("主 goroutine")
		time.Sleep(time.Second)
		// 主协程退出后子协程也会退出
		if i == 2{
			break
		}
	 }
}
