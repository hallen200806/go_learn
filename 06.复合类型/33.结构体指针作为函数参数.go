package main

import "fmt"

type person1 struct {
	id int
	name string

}

func test_type_zhizhen(p *person1)  {
	p.name = "hl"
}
func main() {
	// 地址传递
	per := person1{id:11,name:"hallen"}
	fmt.Println(per)
	p := &per
	fmt.Printf("%p\n",p)
	fmt.Println(*p)
	test_type_zhizhen(p)
	fmt.Println(*p)
}
