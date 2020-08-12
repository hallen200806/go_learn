package main

import (
	"os"
	"fmt"
)

func main() {
	/*
	os包处理命令行参数
	接收命令行参数，都是以字符串方式传递
	 */
	// 命令行执行：go run 10.获取命令行参数.go 1 2 3
	 // [C:\Users\lenovo\AppData\Local\Temp\go-build913927678\b001\exe\10.获取命令行参数.exe 1 2 3] 文件也是个参数
	 list := os.Args
	 n := len(list)
	 fmt.Println(list)
	 fmt.Println(n)

	 for _,data := range list{
	 	fmt.Println(data)
	 }
}
