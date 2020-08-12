package main

import "fmt"

func main() {

	s := []int{1,2,3,4,5}
	p := &s
	fmt.Println(p)
	fmt.Printf("%p\n",p)
	fmt.Println(*p)

	// 修改值
	(*p)[1] = 222
	fmt.Println(*p)
	// for循环取值
	for i:=0;i<len(s) ;i++  {
		fmt.Println((*p)[i])
	}
}
