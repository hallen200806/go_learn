package main

import "fmt"

func bubble_sort(a [10]int) ([10]int) {

	fmt.Println(a)
	for i:=0;i<len(a)-1 ;i++  {
		for j:=0;j<len(a)-1-i;j++{
			if a[j]>a[j+1] {
				a[j],a[j+1] = a[j+1],a[j]
			}
		}
	}
	return a

}
func main() {
	/*
	数组作为函数参数是值传递
	 */
	a := [10]int{0,9,2,6,4,5,3,7,8,1}
	arr := bubble_sort(a)
	fmt.Println(a)
	fmt.Println(arr)
}
