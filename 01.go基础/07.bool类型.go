package main

import "fmt"

func main()  {
	// bool类型只有true或false(首字母是小写，和python区别开),没有其他的值
	// 默认为false
	var a bool
	b := true
	// 重新赋值为false
	b = false
	c := (1!=2)
	fmt.Println(c)
	fmt.Println(a)
	fmt.Println(b)
}
