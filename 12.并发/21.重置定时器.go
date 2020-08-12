package main

import (
	"time"
	"fmt"
)

func main() {


	timer := time.NewTimer(10 * time.Second)

	ok := timer.Reset(time.Second)		// 重置为1秒
	fmt.Println(ok)
	<- timer.C
	fmt.Println("定时时间到")



}
