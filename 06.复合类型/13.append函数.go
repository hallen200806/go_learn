package main

import "fmt"

func main() {
	/*
	append是追加到末尾
	如果超过原来的容量，通常是以两倍的容量扩容
	 */
	s := make([]int,4)
	s1 := append(s, 2)
	// 追加到末尾
	fmt.Println(s1)
}
