package main

import "fmt"

func main() {
	// 要求：求数组中的最大值，最小值和所有元素的和

	// 问题：如果数组的元素中没有0，那就求不出来最小值,
	// 解决办法，赋值为数组中的某一个值，这样就是数组元素间比较了
	a := [5]int{1,2,3,4,5}
	fmt.Println(a)
	var max int	= a[4]
	//var min int = 10
	var min int = a[4]
	var sum int
	for i:=0;i<len(a) ;i++  {

		if a[i]>max {
			max = a[i]
		}
		if a[i]<min{
			min = a[i]
		}
		sum += i
	}
	fmt.Println(max)
	fmt.Println(min)
	fmt.Println(sum)
}
