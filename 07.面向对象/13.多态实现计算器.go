package main

import "fmt"

type operate struct {
	num1 int
	num2 int
}

type add_num struct {
	operate
}
type struc_num struct {
	operate
}
// 加法
func (a *add_num)operate_func() int {
	ret := a.num1 + a.num2
	//fmt.Println(ret)
	return ret
}
// 减法
func (s *struc_num)operate_func() int {
	ret := s.num1 - s.num2
	//fmt.Println(ret)
	return ret
}

// 定义一个接口
type operate_funcer interface {
	operate_func() int	// 如果方法中有返回值，这里必须指定返回值类型

}

// 多态
func operate_inter(o operate_funcer)  {
	ret := o.operate_func()
	fmt.Println("===========")
	fmt.Println(ret)
}

func main() {
	a := add_num{operate:operate{num1:10,num2:20}}
	s := struc_num{operate:operate{num1:10,num2:20}}
	//var o operate_funcer
	//o = &a
	//o.operate_func()
	//o = &s
	//o.operate_func()
	operate_inter(&a)
	operate_inter(&s)

}
