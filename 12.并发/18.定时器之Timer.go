package main

import (
	"time"
	"fmt"
)

func main() {
	/*
	Timer：只会响应一次，只会执行一次事件
		是一个定时器，代表未来的一个单一事件，你可以告诉Timer你要等待多长时间，
		它提供了一个channel，在将来的那个时间那个channel提供了一个时间值
	 */

	 // 创建一个定时器，2秒后
	 timer := time.NewTimer(2 * time.Second)

	 // 当前时间
	 fmt.Println(time.Now())

	 // 2s后往time.C写数据，有数据后就可以读取
	 t := <-timer.C
	 fmt.Println(t)


	 // 验证time.NewTimer(),时间到了，只会相应一次
	 timer_2 := time.NewTimer(3*time.Second)
	 for {
	 	<- timer_2.C		// 这里只会写入一次，然后会导致死锁：fatal error: all goroutines are asleep - deadlock!
	 	fmt.Println("时间到了")
	 }
}
