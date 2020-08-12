package main

import "fmt"

func swap(p *[3]int)  {
	(*p)[0] = 111
}
func main() {
	// 地址传递，形参改变实参也会改变
	a := [3]int{1,2,3}
	fmt.Println(a)
	swap(&a)
	fmt.Println(a)
}
