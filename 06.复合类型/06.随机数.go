package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*
	1.使用math/rand包
	2.随机数种子
	3.生成随机数
	 */
	// 随机数种子,如果种子不变则产生的随机数不变
	rand.Seed(time.Now().UnixNano())

	fmt.Println(rand.Int())

	fmt.Println(rand.Intn(10)) // 生成一个10以内的数字，不包括10

	// 生成六位数随机数
	var arr [6]int
	for i:=0;i<len(arr) ;i++  {
		c := rand.Intn(10)
		arr[i] = c
		//fmt.Println(c)
	}
	fmt.Println(arr)

	// 生成六位不相同的随机数
	var arr6 [6]int
	for i:=0;i<len(arr6);i++{
		v := rand.Intn(10)
		for j:=0;j<i ;j++  {
			if v == arr6[j]{
				v = rand.Intn(10)
				j = -1
			}
		}
		arr6[i] = v
	}

	fmt.Println(arr6)
}
