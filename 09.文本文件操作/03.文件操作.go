package main

import (
	"fmt"
	"os"
)

func main() {
	/*
	步骤：
		1.导入os包，创建文件，读写文件的函数都在这个包
		2.指定创建的文件路径以及文件名，如果不指定路径，则默认为项目根目录下
		3.执行Create()函数，创建文件
		4.关闭文件，如果不关闭，造成内存浪费和会导致打开文件上限错误
	文件创建失败的可能原因：
		1.路径不存在
		2.文件的权限问题
		3.程序打开文件上限
	文件路径：
		1.D:\\software\\a.txt    对\进行转义
		2.D:/software/a.txt		使用/
	 */
	// func Create(name string) (*File, error) {...}
	fp,err := os.Create("a.txt") //如果不指定路径，则默认为项目根目录下
	if err != nil {
		fmt.Println("创建失败")
		return
	}
	// 可以使用defer延迟调用
	defer fp.Close()
	//fp.Close()
}
