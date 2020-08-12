package main

import "fmt"

func exchange_value(num1,num2 *int)  {
	*num1,*num2 = *num2,*num1
	fmt.Println(num1,num2)
}
func main() {
	// 地址传参
	a := 10
	b := 20
	exchange_value(&a,&b)
	fmt.Println(a,b)	// 20 10变了，说明是地址传参

}
