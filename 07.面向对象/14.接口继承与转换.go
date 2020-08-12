package main

import "fmt"

// 两个类都有同一个方法，teacher还有一个唱歌的方法，然后有两个接口，一个接口继承另外一个接口
type student23 struct {
	name string
	age int
}

func (s *student23)say_hi()  {
	fmt.Printf("我是：%s,我今年:%d \n",s.name,s.age)
}

type teacher23 struct {
	name string
	age int
}
func (t *teacher23)say_hi()  {
	fmt.Printf("我是：%s,我今年:%d \n",t.name,t.age)
}

func (t *teacher23)sing()  {
	fmt.Printf("我是：%s,我今年:%d,我会唱歌 \n",t.name,t.age)
}

// 第一个接口
type person7 interface {
	say_hi()
}
// 第二个接口
type person8 interface {
	person7 // 匿名字段，继承person7
	sing() // 还有一个唱歌的方法
}
func main() {
	/*
	一个接口中继承另一个接口，继承的称为超集，被继承的称为子集
	一个接口赋值给另一个接口
		注意：超集中包含所有子集的方法，所以可以将超级赋值给子集，反过来不行
		p_s = p
	 */
	s := student23{name:"小红",age:18}
	t := teacher23{name:"老王",age:33}
	var p person8
	//p = &s  //这里报错，因为student23里面没有sing方法，person8中有
	p = &t
	p.say_hi()
	p.sing()

	// 学生没有sing方法，只能调用没有sing的接口
	var p_s person7
	p_s = &s
	p_s.say_hi()

	// 一个接口赋值给另一个接口
	// 注意：超集中包含所有子集的方法，所以可以将超级赋值给子集，反过来不行
	p_s = p
}
