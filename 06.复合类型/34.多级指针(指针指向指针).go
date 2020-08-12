package main

import "fmt"

func main() {

	/*
	指针指向指针，为多级指针
		一级指针：*int
		二级指针：**int
		三级指针：***int
		...
	*/

	a := 10
	// 一级指针
	p := &a
	// 二级指针
	pp := &p
	fmt.Printf("%p\n",p)
	fmt.Printf("%p\n",pp)
	fmt.Println(*p)
	fmt.Println(*pp)
	fmt.Printf("%T\n",p)	// *int
	fmt.Printf("%T\n",pp)	// **int
}
