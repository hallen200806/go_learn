package main

import (
	"time"
	"fmt"
)

func main() {

	// 延时2秒后再继续执行
	timer := time.NewTimer(2 * time.Second)
	<- timer.C
	fmt.Println("时间到")		// 延时2秒再打印


	// 第二种方式
	time.Sleep(2 * time.Second)
	fmt.Println("时间到")		// 延时2秒再打印

	// 第三种方式
	<- time.After(2 * time.Second)  // 定时2秒，阻塞两秒，两秒后往channel写内容
	fmt.Println("时间到")		// 延时2秒再打印

}
