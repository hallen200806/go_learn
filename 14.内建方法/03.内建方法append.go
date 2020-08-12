package main

import "fmt"

func test_append()  {
	mslice := make([]int,3)
	mslice[0] = 1
	mslice[1] = 2
	mslice[2] = 3
	// 有返回值，需要slice接收,可以一次性传多个值，一次性添加多个值
	// func append(slice []Type, elems ...Type) []Type
	mslice = append(mslice, 4,5)
	mslice = append(mslice, 6)
	fmt.Println(mslice)
	// 长度
	fmt.Println(len(mslice))
	// 容量
	fmt.Println(cap(mslice))


	
}
func main() {

	test_append()
}
