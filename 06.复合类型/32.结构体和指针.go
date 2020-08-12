package main

import "fmt"

type person struct {
	id int
	name string
	age int

}
func main() {
	per := person{id:1,name:"yhl",age:18}
	// per地址赋值给指针
	p := &per		//也可以这样：var p1 *person = &per

	fmt.Println(per)
	fmt.Printf("%p\n",p)
	// 查看类型
	fmt.Printf("%T\n",p)

	// 修改结构体成员的值
	(*p).age = 20
	fmt.Println(*p)
	p.id = 110
	fmt.Println(*p)
}
