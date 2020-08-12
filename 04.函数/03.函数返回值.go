package main

import "fmt"
/*
函数返回值：
	1.在花括号之前必须定义返回的类型，和实际返回的类型一致
	2.如果返回值指定了变量名，那么return后面可以不写返回的变量，
		如果返回值指定了变量名和类型，
		那么在函数体里面的返回变量就可以直接等于，而不用:=,
		以为返回的时候指定了变量的类型，参考test7
*/
func test6()  int{
	c := 2
	return c
}

func test7() (d int)  {
	// 注意这里是直接等于，而不是 := ，因为前面返回值定义了类型int，下面就不需要定义了
	d = 3
	return

}

// 返回多个值
func test8() (a,b,c int) {
	a,b,c = 5,6,7
	return
}
func main() {
	a1 := test6()
	fmt.Println(a1)

	a2 := test7()
	fmt.Println(a2)

	c1,c2,c3 := test8()
	fmt.Println(c1,c2,c3)
}
