package main

import "fmt"

func main() {
	/*
	切片：
		数组[起始位置(low):结束位置(high):数组长度(容量max)]
	注意：左开右闭区间
		cap(数组)：切片数组的容量，总是>= len(数组)

	坑：
		对切片进行了重新赋值，会改变原始的数组
		e := a[1:4]
		e[2] = 999
		fmt.Println(e)
		fmt.Println(a)
			a = [1 2 3 4 56 7]
			e = [2 3 999]
			a = [1 2 3 999 56 7]

	 */

	a := []int{1,2,3,4,56,7}
	b := a[0:3]		// 0可以省略，a[:3]
	c := a[0:3:5]

	d := a[1:]	// 第二个到结束
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	//对切片进行了重新赋值，会改变原始的数组
	// 切片的索引是2，但是在原来的数组中的切片对应的是3
	e := a[1:4]
	e[2] = 999
	fmt.Println(e)
	fmt.Println(a)
	
}
