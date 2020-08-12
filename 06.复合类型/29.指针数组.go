package main

import "fmt"

func main() {
	// 和数组指针区分开，指针数组里面存的都是指针,是指针类型

	a := 10
	b := 20
	c := 30
	var arr [3]*int = [3]*int{&a,&b,&c}
	fmt.Println(arr)

	// 修改指针
	*arr[1] = 200
	fmt.Println(b)  // 从原来的20改为了200
	
}
