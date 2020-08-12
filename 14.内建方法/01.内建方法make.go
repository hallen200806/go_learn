package main

import (
	"fmt"
	"reflect"
)

func make_slice()  {
	// slice
	mslice := make([]int,3)
	mslice[0] = 1
	mslice[1] = 2
	mslice[2] = 3

	// 扩容使用append，不能通过下标，因为前面已经固定了容量为3
	mslice = append(mslice,4)
	fmt.Println(mslice)
}

func makemap()  {
	// map
	mmap := make(map[int]int)
	mmap[1] = 10
	mmap[2] = 20
	mmap[3] = 30
	fmt.Println(mmap)
	fmt.Println(reflect.TypeOf(mmap))   // map[int]int  引用类型
}

func makechan()  {
	mchan := make(chan int)
	close(mchan)
}
func main() {
	/*
		可以用来构建：slice、map、chan
		返回的是类型引用
 	*/
 	make_slice()
 	makemap()
 	makechan()



}
