package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	/*
	读取文件的流程：
		1.打开文件
		2.对文件进行读取
		3.关闭文件

	读取方式：
		1.Read
			参数是字符前片，表示读取多少的长度
		2.创建缓存区，通过行读取
			r := bufio.NewReader(fp)
			// 行读取，截取标志 '\n'
			slice,_ := r.ReadBytes('\n')
			fmt.Println(string(slice)) // 第一行
	 */

	fp,err := os.Open("b.txt")

	// Read
	if err != nil {
		fmt.Println("打开文件失败")
	}
	a := make([]byte,1024*2)   // 2k大小

	for {
		n,err := fp.Read(a)
		if err == io.EOF{  // io.EOF是文件的结尾
			break
		}
		fmt.Println(string(a[:n]))
	}

	// 创建缓存区，通过行读取
	r := bufio.NewReader(fp)
	// 行读取，截取标志 '\n'
	slice,_ := r.ReadBytes('\n')
	fmt.Println(string(slice)) // 第一行
	// 循环读取所有的行
	for {
		buf,err := r.ReadBytes('\n')
		if err != nil{
			if err == io.EOF {
				fmt.Println("读完了")
				break
			}
			fmt.Println(err)

		}
		fmt.Println(string(buf))
	}



	defer fp.Close()
}
