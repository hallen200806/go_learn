package main

import (
	"fmt"
	"strings"
)
func main() {
	/*
	1.contains：查找一个字符串是否在另一个字符串中,返回bool类型
		strings.contains(被查找的字符串，查找的字符串)
	2.Index:查找一个字符串在另一个字符串中第一次出现的位置,返回下标，整型，-1表示找不到
		strings.Index(被查找的字符串，查找的字符串)

	3.Join：连接多个，相当于python中的join，返回字符串
		strings.Join(字符串列表,连接符)

	4.Repeat：重复一个字符串n次,返回字符串
		strings.Repeat("=",5)

	5.Replace：替换
		strings.Replace(要替换的字符串,旧的,新的,替换次数)
			strings.Replace(str1,"h","ttt",2)
			替换次数为-1则全部替换
	6.Split：切割字符串，返回slice切片
		strings.Split(要切割的字符串,切割符)
			strings.Split(str1,",")
	7.Trim：去除字符串首尾的内容
		strings.Trim(目标字符串,去除的内容)
			strings.Trim(str5,"")
	8.field：去除字符串中的空格，转成slice切片，一般用于统计单词的个数
		strings.Fields(目标字符串)
			strings.Fields(str5)
	 */
	str1 := "hello world"
	str2 := "h"
	// 查找字符串是否在另一个字符串中
	ret := strings.Contains(str1,str2)
	fmt.Println(ret)
	// 查找一个字符串在另一个字符串中第一次出现的位置
	ret2 := strings.Index(str1,str2)
	fmt.Println(ret2) // 下标为0
	// 字符串切片列表
	slice := []string{"123","456","789"}
	// 使用join串起来
	str3 := strings.Join(slice,"")
	fmt.Println(str3)

	// 重复一个字符串n次
	str4 := strings.Repeat("=",5)
	fmt.Println(str4)

	ret3 := strings.Replace(str1,"h","ttt",2)
	fmt.Println(ret3)

	ret4 := strings.Split(str1,",")
	fmt.Println(ret4)

	str5 := " 58h khuo ada "
	ret5 := strings.Trim(str5,"")
	fmt.Println(ret5)

	//field
	ret6 := strings.Fields(str5)
	fmt.Println(ret6)

}
