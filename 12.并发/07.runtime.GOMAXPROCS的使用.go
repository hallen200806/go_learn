package main

import (
	"runtime"
	"fmt"
)

func main() {
	/*
	runtime.GOMAXPROCS()：
		用来设置可以并行计算的CPU合核数的最大值，并返回之前的值

	核数越大频繁交互越快

	 */

	 n := runtime.GOMAXPROCS(4)   // 指定1个核数运行,核数越大频繁交互越快
	 fmt.Println(n)   // 之前是8核的

	 for {
	 	go fmt.Println(1)

	 	fmt.Println(0)
	 }
}
