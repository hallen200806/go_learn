package main

import "fmt"

type person33 struct {
	id int
	name string
}

func (p person33)say()  {
	fmt.Println("说话的方法")
}

func (p person33)run()  {
	fmt.Println("跑的方法")
}

func main() {
	p := person33{1,"hallen"}
	// 方法表达式：通过类型
	f := (*person33).run
	f(&p)

	// 也可以不传指针
	f2 := (person33).say
	f2(p)
}
