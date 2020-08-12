package main

import (
	"fmt"
	//"time"
	"runtime"
)

func test_goexit()  {
	for i :=0;i<5;i++{
		fmt.Println("子协程")
	}

}

func test_defer()  {
	defer fmt.Println("这是defer")
	// 终止次函数
	runtime.Goexit()		// 终止当前goroutine的执行
	//return
}
func main() {
	/*
	runtime.Goexit()将会终止当前goroutine的执行，调度器确保所有已注册defer延迟调用被执行
	 */

	 go test_goexit()
	 go test_defer()
	 for i :=0;i<10;i++{

	 	fmt.Println("主协程")
	 	//time.Sleep(time.Second)  // 这里得沉睡下，不然不使用runtime的函数子协程就不会执行
	 }

}
