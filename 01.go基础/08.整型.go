package main

import (
	"fmt"
	"unsafe"
)

func main()  {
	/*
	int和uint
		int：有符号整型,4个长度，32位
			比如：b := -10正常
			int8:1字节长度
			int16：2字节长度
			int32：4字节长度
			int64：8字节长度
		uint：无符号整型，8个长度，64位
			比如 a := -10就会报错
			uint8：1字节长度
			uint16:2字节长度
			uint32:4字节长度
			uint64:8字节长度
	*/
	a := -10
	fmt.Println(a)
	fmt.Printf("%T\n",a)
	// 4个字节长度
	var b int32 = 10
	fmt.Println(unsafe.Sizeof(b))
}
