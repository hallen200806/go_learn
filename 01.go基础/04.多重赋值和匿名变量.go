package main

import "fmt"

func test1() (a,b,c int) {
	// 匿名变量：_   丢弃的变量
	//a,b,c := 1,2,3
	return 1,2,3
}

func main()  {
	a,b := 10,20
	//a = 20
	//b = 10
	a,b = b, a
	fmt.Println(a)
	fmt.Println(b)
	//var d,e int

	// 这里我只需要第二个和第三个值，第一个值不需要，
	// 可以不用定义一个变量去接收，这就是匿名函数，用_代替
	_,d,e := test1()
	fmt.Println(d,e)
}
