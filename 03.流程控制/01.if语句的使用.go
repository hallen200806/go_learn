package main

import "fmt"

func main() {
	/*
	if 条件表达式 {

	}
	注意：条件表达式没有括号
	 */

	//第一种：推荐使用

	a := 100
	if a==100{
		fmt.Println("等于100")
	}


	// 第二种：初始化
	if b := 100;b == 100{
		fmt.Println("等于100了")
	}else {
		fmt.Println("不等于100")
	}

}
