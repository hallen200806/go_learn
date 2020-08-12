package main

import (
	"fmt"
	"runtime"
	"time"
)

func test_gosched()  {
	for i :=0; i<20;i++{
		fmt.Println("子协程")
	}
}
func main() {
	/*
	tuntime.Gosched()用于让出cpu时间片，让出当前goroutine的执行权限，调度安排其他等待的任务执行，并在某个时候从该位置恢复执行
		这就像接力比赛，A跑了一会碰到代码runtime.Gosched(),就把接力棒交给了B，A歇着了，B接着继续跑
	 */

	go test_gosched()
	for i :=0; i<10;i++{
		// 让出时间片，先让其他协程执行
		runtime.Gosched()   // 先从子协程开始，子协程执行完了，主协程记着继续
		fmt.Println("主协程")
		time.Sleep(time.Second)  // 这里得沉睡下，不然不使用runtime的函数子协程就不会执行
	}
	 
}
