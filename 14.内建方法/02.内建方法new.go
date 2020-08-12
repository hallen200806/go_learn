package main

import (
	"fmt"
	"reflect"
)

func new_map()  {
	nmap := new(map[int]int)
	// 查看类型
	fmt.Println(reflect.TypeOf(nmap))  // *map[int]int 执行类型
	
}
func main() {
	/*
		内存置零
		返回传入类型的指针类型
	和make类似，区别：new是指针类型，make是引用类型
	 */

	 new_map()
}
