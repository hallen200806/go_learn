package main

import "fmt"

// 切片可以不指定长度，和数组的区别
func demo(s []int) (b []int) {
	fmt.Println(s)
	b =s[0:2]
	//fmt.Println(b)
	//切片和数组的区别：传递的是地址，如果形参里面改变了值，实参也会改变
	s[2] = 999
	// append不会改变实参的数据，只有修改原来的数据才会对实参修改
	fmt.Println(s)
	s = append(s,6,7,8,9)
	return s
}
func main() {
	/*
	切片本身就是个地址
	实参传递给形参的是个地址
	指向同一块内存地址，操作的是相同的数据
	形参的改变可以改变实参的值
	注意：
		切片和数组的区别：传递的是地址，如果形参里面改变了值，实参也会改变
		append不会改变实参的数据，只有修改原来的数据才会对实参修改
			因为append已经改变了内存地址

	 */
	a :=[]int{1,2,3,4,5}
	// 切片和数组的区别：传递的是地址，如果形参里面改变了值，实参也会改变
	c := demo(a)
	fmt.Println(a)
	fmt.Println(c)
}
