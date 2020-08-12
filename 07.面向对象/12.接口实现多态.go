package main

import "fmt"

// 有两个类，都有一个方法叫say_hi，一个类多了个say_hello
type student22 struct {
	name string
	age int
}

func (s *student22)say_hi()  {
	fmt.Printf("我是：%s,我今年:%d \n",s.name,s.age)
}

type teacher22 struct {
	name string
	age int
}

func (t *teacher22)say_hi()  {
	fmt.Printf("我是：%s,我今年:%d \n",t.name,t.age)
}

// 定义一个接口
type humaner interface {
	say_hi()
}

// 多态的实现，接口作为函数的参数
func say_hello(h humaner)  { // 普通函数，因为没有指定类型
	h.say_hi()
}
func main() {
	/*
	多态的定义：
		同一个接口，实现了不同的实例，执行不同的操作
	 */
	stu := student22{name:"小明",age:15}
	tea := teacher22{name:"老李",age:32}
	// 调用多态
	say_hello(&stu)
	say_hello(&tea)
	
}
