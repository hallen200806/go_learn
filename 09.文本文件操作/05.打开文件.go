package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		打开文件：
			1.open：以只读方式打开，不能改变
				fp,err := os.Open("a.txt")	文件路径
			2.OpenFile:有三个参数，
				第一个是文件路径
				第二个参数表示打开模式
					O-RDONLY:只读模式
					O_WRONLY:只写模式
					O_EDWR:可读可写模式
					O_APPEND:追加模式

				第三个参数表示权限，取值范围：0~7
					0：表示没有任何权限
					1.执行权限，如果是可执行文件，是可以运行的
					2.写权限
					3.写权限与执行权限
					4.读权限
					5.读权限与可执行权限
					6.读权限和写权限
					7.读权限，写权限，可执行权限
	*/

	// *File, error
	//fp,err := os.Open("a.txt")  // 文件路径
	fp,err := os.OpenFile("a.txt",os.O_RDWR,6)

	if err != nil{
		fmt.Println("读取文件失败")
	}

	fp.WriteString("我的测试")   // 覆盖了原来的
	fp.WriteAt([]byte("我的测试啊"),20)

	// 关闭
	defer fp.Close()
}
