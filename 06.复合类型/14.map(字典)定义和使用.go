package main

import "fmt"

func main() {
	/*
	map就是键值对构成的字典：无序的
		定义：
			var 变量名 map[int]string [int]这里是int，那么key就只能是int
				第一种：var mapa map[int]string
				第二种：mapc := make(map[string]int)
				第三种：mapb := make(map[int]string,3)
			初始化：
				mapd := map[string]string{"name":"张三","age":"18"}
		取值：
			map[key]
				mapd := map[string]string{"name":"张三","age":"18"}
				name_value :=mapd["name"]

				keyd,valued := mapd["name"]
				fmt.Println(keyd)
				// 如果有该key，则value为true，否则为false
				fmt.Println(valued)
			// 循环取出
			for k,v := range mapd{
				fmt.Println(k)
				fmt.Println(v)
			}
		字典中不能使用cap函数查看容量，只能使用len
			键不能重复，值可以重复

	 */
	// 定义1：var mapa map[int]string
	var mapa map[int]string
	fmt.Println(mapa)
	fmt.Println(len(mapa))
	// 定义2：mapb := make(map[int]string,3)
	mapb := make(map[int]string,3)
	fmt.Println(mapb)

	// 赋值：
	mapb[1] = "张三"
	mapb[2] = "李四"
	mapb[3] = "王五"
	mapb[4] = "dasd"
	fmt.Println(mapb)

	mapc := make(map[string]int)
	mapc["name"] = 18
	fmt.Println(mapc)

	mapd := map[string]string{"name":"张三","age":"18"}
	fmt.Println(mapd)
	name_value :=mapd["name"]
	fmt.Println(name_value)

	keyd,valued := mapd["name"]
	fmt.Println(keyd)
	fmt.Println(valued)	// 如果有值则为true，否则为false
}
