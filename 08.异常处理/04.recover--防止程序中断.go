package main

import "fmt"

func testA(){
	fmt.Println("testA")
}

func testB(i int)  {
	// 防止程序奔溃
	defer func() {
		// 加if判断，否则不会正常的也是nil
		if err := recover();err != nil{
			fmt.Println(err)	// 打印错误信息
		}
		//recover() // 设置recover，必须在defer中使用
		//fmt.Println(recover())	// 打印错误信息
	}()

	var arr [3]int
	arr[i] = 99
	fmt.Println("testB")
}

func testC()  {
	fmt.Println("testC")
}

func main() {
	/*
	设置recover，必须在defer中使用
	记得加判断，否则正常的也会为nil
	 */
	testA()
	testB(3)  // 角标越界，会报错，需要在testB中进行处理
	//testB(2)  // 如果正常则需要加if判断
	testC()
}
