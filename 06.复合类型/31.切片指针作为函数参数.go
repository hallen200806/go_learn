package main

import "fmt"

func test_slice_zhizhen(s *[]int)  {
	*s = append(*s,6,7,8,9)
}
func main() {
	// 地址传递
	s := []int{1,2,3,4,5}
	p := &s
	fmt.Printf("%p\n",p)
	fmt.Println(*p)
	test_slice_zhizhen(p)
	fmt.Println(*p)
}
