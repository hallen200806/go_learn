package main

import "fmt"

func test_copy()  {
	//
	mslice1 := make([]int, 3)
	mslice2 := make([]int, 4)
	mslice1[0] = 1
	mslice1[1] = 2
	mslice1[2] = 3

	mslice2[0] = 4
	mslice2[1] = 5
	mslice2[2] = 6
	mslice2[3] = 7
	// 把mslice2 拷贝给mslice1，超过的长度不拷贝
	copy(mslice1, mslice2)

	fmt.Println(mslice1)   // [4 5 6]
}
func main() {
	test_copy()
}
