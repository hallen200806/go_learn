package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	/*
	案例：文件拷贝
	 */

	var srcFileName string
	var dstFileName string
	fmt.Println("请输入源文件名称")
	fmt.Scan(&srcFileName)
	fmt.Println("请输入目的文件名称")
	fmt.Scan(&dstFileName)
	if srcFileName == dstFileName {
		fmt.Println("源文件与目的文件不能相同")
		return
	}

	// 只读方式打开源文件
	fp,err := os.Open(srcFileName)
	if err != nil{
		fmt.Println("打开源文件失败",err)
		return
	}

	// 新建目的文件
	df,err2 := os.Create(dstFileName)
	if err2 != nil{
		fmt.Println("创建目的文件失败",err2)
		return
	}

	defer fp.Close()
	defer df.Close()

	buf := make([]byte,1024*4)  // 4k
	for {
		n,err3 := fp.Read(buf)
		if err3 != nil{
			fmt.Println("读取失败",err3)
			if err3 == io.EOF{
				fmt.Println("读完了")
				break
			}
		}
		df.Write(buf[:n])
	}

}
