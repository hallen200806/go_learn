package main

import "fmt"

func main() {
	/*
	空接口可以接受任意类型的数据，没有方法
	 */
	var i interface{}
	i = 10
	fmt.Println(i)
	i = "Go"
	fmt.Println(i)

	var ii []interface{}
	ii = append(ii,1,2,3,4,"go")
	fmt.Println(ii)
	fmt.Printf("%T\n",ii)
	fmt.Printf("%p\n",ii)
}
