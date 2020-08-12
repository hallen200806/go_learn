package main

import (
	"fmt"
	"time"
)

func new_task()  {
	for{
		fmt.Println("这是子协程")
		time.Sleep(time.Second)
	}
}

func main() {
	/*
	goroutine是go并发设计的核心，goroutine说到底就是个协程，但它比协程更小，十几个goroutine可能体现在底层就是五六个协程
		go语言内部实现了这些goroutine之间的内存共享
	执行goroutine只需极少的栈内存(大概4~5kb)，当然会根据相应的数据伸缩
	goroutine比thread更易用，更高效，更轻便
	在并发编程里面，通常我们想将一个过程切分成几块，然后让每个goroutine各自负责一块工作

	当一个程序启动时，其主函数即在一个单独的goroutine中运行，我们叫他main goroutine，新的goroutine继续创建
	 */



	// 需求：写一个死循环，执行这个死循环的时候还能调用另外一个函数

	go new_task()		// go关键字，新建一个协程,这个时候两个死循环会轮换着执行

	 for{
	 	fmt.Println("这是主协程")
	 	time.Sleep(time.Second) // 沉睡一秒
	 }
}
