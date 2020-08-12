package main

import "fmt"

func main() {
	/*
	for 初始值;判断条件;条件变化{
		代码体
	}
	 */
	a :=10
	for i:=0;i<a;i++{
		fmt.Println(i)
	}

	// 计算1-10的和
	b :=10
	sum :=0
	for i:=1;i<=b;i++{
		sum+=i
	}
	fmt.Println(sum)
	// 遍历字符串
	c := "abc"
	for i:=0;i<len(c) ;i++  {
		//fmt.Println(c[i])
		fmt.Printf("%c\n",c[i])
	}
	fmt.Println("===========range")
	// range
	// 只要下标
	for i:=range c{
		fmt.Println(i)
	}
	//只要值
	for _,data:=range c{
		fmt.Printf("%c\n",data)
	}
	// 下标和值都要
	for i,data:=range c{
		//fmt.Println(i,data)
		fmt.Print(i)
		fmt.Printf("%c\n",data)
	}

	// for循环嵌套
	for i:=0;i<5 ;i++  {
		for j:=0;j<5 ;j++  {
			fmt.Println(i,j)
		}
	}
}
