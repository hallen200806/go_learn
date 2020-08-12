package main

import "fmt"

type person2 struct {
	id int
	name string
}
func main() {
	/*
	结构体比较：
		==  或者!=


	 */

	 p1 := person2{1,"hl"}
	 p2 := person2{2,"hl"}
	 p3 := person2{1,"hl"}
	 fmt.Println(p1 == p2)
	 fmt.Println(p1 == p3)

	 // 结构体赋值
	 var p4 person2
	 p4 = p1
	 fmt.Println(p4)

}
