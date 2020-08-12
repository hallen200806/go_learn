package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	/*
	文件的写入类型：
		1.WriteString：接收字符串类型
			后面加\n也不会换行的，windows文本文件中使用\r\n换行的
		2.Write:接收的是byte类型的切片

		3.WriteAt:在指定位置写入,参数是要写入的字符切片以及写入的位置(光标流位置)
			通过Seek获取位置: offset int64, whence int
			count2,_ := fp.Seek(0,io.SeekEnd)
			fp.WriteAt([]byte("hello WriteAt"),count2)
			也可以设置为想要的数字，int类型的
			fp.WriteAt([]byte("hello WriteAt"),0)
			注意：
				两次写入的位置相同，会覆盖之前写入的

	 */

	fp,err := os.Create("./b.txt")
	if err != nil{
		fmt.Println("创建文件失败")
	}
	// WriteString写入文件
	fp.WriteString("写入测试\r\n")   // 换行使用\r\n，只有\n不会换行
	fp.WriteString("test\r\n")

	//Write接收的是byte类型的字符切片
	a := []byte{'h','e','l','l','o'}
	count,err1 :=fp.Write(a)
	count1,err2 :=fp.Write([]byte("这是测试写入"))
	if err1 != nil{
		fmt.Println("写入失败")
		return
	}else {
		fmt.Println(count)
	}
	if err2 != nil{
		fmt.Println("写入失败")
		return
	}else {
		fmt.Println(count1)
	}

	// 在指定位置写入:WriteAt,要写入的字符切片以及写入的位置
	// 通过Seek获取位置: offset int64, whence int
	count2,_ := fp.Seek(0,io.SeekEnd)
	fp.WriteAt([]byte("hello WriteAt"),count2)


	// 延迟关闭
	defer fp.Close()

}
