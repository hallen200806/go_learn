package main

import "fmt"

func main() {
	/*
	判断是否是某个类型的变量，返回true或false
		语法：
			value,ok := element.(T)
				value就是变量的值,名字随便起
				ok是一个bool类型
				element是interface变量,自定义的变量
				T是断言的类型
	 */

	var i interface{}
	i = "10"
	value,ok := i.(int)
	fmt.Println(ok)
	if ok {
		fmt.Printf("整型数据%d",value)
	}else {
		fmt.Printf("不是整型%d",value)
	}


	
}
