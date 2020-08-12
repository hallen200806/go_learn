package main

import "fmt"

func main() {
	/*
	注意：
		else是跟在if的结束花括号后面的，不能另起一行
		没有这种表达： 100 <a < 200,只能使用if嵌套满足a>100以后在里面if a<200
	*/
	a := 100
	if a == 100 {
		fmt.Println("等于100")
	}else {
		fmt.Println("不等于100")
	}

	if b :=99; b == 100 {
		fmt.Println("等于100了")
	}else if b < 100{
		fmt.Println("小于100了")
	} else {
		fmt.Println("不等于100")
	}
}
